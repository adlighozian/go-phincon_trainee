package main

import (
	"contact-go/config"
	"contact-go/db"
	"contact-go/handler"
	"contact-go/repository"
	"contact-go/usecase"
	"fmt"
	"net/http"
)

func main() {
	config := config.LoadConfig()

	contactUC := createContactUseCase(config)
	contactHTTPHandler := handler.NewContactHandlerHttp(contactUC)
	NewServer(contactHTTPHandler)
}

func createContactUseCase(config *config.Config) usecase.ContactUseCase {
	var contactRepo repository.ContactRepository
	db := db.GetMysql(config)
	contactRepo = repository.NewContactRepository(db)

	return usecase.NewContactUseCase(contactRepo)
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
	fmt.Println("Server running localhost:", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
