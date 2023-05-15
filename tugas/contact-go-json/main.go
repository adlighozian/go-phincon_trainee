package main

import (
	"contact-go/config"
	"contact-go/db"
	"contact-go/handler"
	"contact-go/repository"
	"contact-go/usecase"
	"net/http"
)

func main() {
	config := config.LoadConfig()

	switch config.Gorm {
	case "true":
		db := db.GetMysqlGorm(config)
		contactRepo := repository.NewContactRepositoryGorm(db)
		usecase := usecase.NewContactUseCase(contactRepo)
		contactHTTPHandler := handler.NewContactHandlerHttp(usecase)
		NewServer(contactHTTPHandler)
	default:
		db := db.GetMysql(config)
		contactRepo := repository.NewContactRepository(db)
		usecase := usecase.NewContactUseCase(contactRepo)
		contactHTTPHandler := handler.NewContactHandlerHttp(usecase)
		NewServer(contactHTTPHandler)
	}

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
}
