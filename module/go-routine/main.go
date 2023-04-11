package main

import (
	"fmt"
	"net/http"
)

func HandlerGet(write http.ResponseWriter, request *http.Request) {
	write.WriteHeader(http.StatusOK)
	fmt.Println("Success", http.StatusOK)
	fmt.Fprintf(write, "Success ini index")
}

func HandlerFormPost(write http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")
	age := request.PostForm.Get("age")

	write.WriteHeader(http.StatusCreated)
	fmt.Println("Success", http.StatusCreated)
	fmt.Printf("Hello %s %s and I am %s years old\n", firstName, lastName, age)
	fmt.Fprintf(write, "Hello %s %s and I am %s years old\n", firstName, lastName, age)
}

func SetCookie(w http.ResponseWriter, r *http.Request) {
	// cookie
	cookie := new(http.Cookie)
	cookie.Name = "X-Powered-By"
	cookie.Value = r.URL.Query().Get("cookie")
	cookie.Path = "/"
	http.SetCookie(w, cookie)

	cookie, err := r.Cookie("X-Powered-By")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

func View(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contoh")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-Powered-By")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

func main() {
	directory := http.Dir("./assets")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.HandleFunc("/", HandlerGet)
	mux.HandleFunc("/setcookie", SetCookie)
	mux.HandleFunc("/post", HandlerFormPost)
	mux.HandleFunc("/belajar", View)
	mux.HandleFunc("/cookie", GetCookie)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Server running")
}
