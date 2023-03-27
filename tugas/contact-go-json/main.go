package main

import (
	"contact-go/config"
	"contact-go/handler"
	"fmt"
	"net/http"
)

func main() {
	// config := config.LoadConfig()
	// var contactRepo repository.ContactRepository

	// switch config.Storage {
	// case "json":
	// 	contactRepo = repository.NewContactRepository()
	// default:

	// }

	// ContactHandler := handler.NewcontactHandler(contactRepo)
	// handler.Menu(ContactHandler)

	ContactHandler := handler.NewContactHandlerHttp()
	NewServer(ContactHandler)

}

func NewServer(handle handler.ContactHandlerHttp) {
	config := config.LoadConfig()
	// server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handle.HandlerGet(w, r)
		} else if r.Method == "POST" {
			handle.HandlerPost(w, r)
		}
	})

	server := http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Server running")
}
