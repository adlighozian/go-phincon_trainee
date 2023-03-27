package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// name := request.URL.Query().Get("name")
	// age := request.URL.Query().Get("age")

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		req := r.URL.Query()
		age := r.URL.Query().Get("age")
		test := req["name"][0]
		test2 := req["name"]
		// test3 := req["age"]
		if test != "" {
			fmt.Fprint(w, "Hello nama saya: ", test, test2, age)
		} else {
			fmt.Fprint(w, "Ini")
		}

	}

	mux.HandleFunc("/", handler)
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello home")
		fmt.Fprintln(w)
		fmt.Fprint(w, r.Method)
		fmt.Fprintln(w)
		fmt.Fprint(w, r.RequestURI)
		fmt.Fprintln(w)
		fmt.Fprint(w, r.Header)
		fmt.Fprintln(w)
		contentType := r.Header.Get("content-type")
		fmt.Fprintln(w)
		fmt.Fprintln(w, "test ini content type :", contentType)

	})
	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		// header := r.Header.Get("Accept")
		fmt.Fprint(w, "Hello home")
		fmt.Fprint(w, r.Header)
	})
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello home")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
