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

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			handle.HandlerUpdate(w, r)
		} else if r.Method == http.MethodPost {
			handle.HandlerPost(w, r)
		} else if r.Method == http.MethodGet {
			handle.HandlerGet(w, r)
		} else if r.Method == http.MethodDelete {
			handle.HandlerDelete(w, r)
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
