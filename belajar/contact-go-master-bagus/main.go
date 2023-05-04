package main

import (
	"contact-go/config"
	"contact-go/config/db"
	"contact-go/handler"
	"contact-go/repository"
	"contact-go/usecase"
	"log"
	"net/http"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	contactUC := createContactUsecase(config)

	switch config.Mode {
	case "http":
		contactHTTPHandler := handler.NewContactHTTPHandler(contactUC)
		err := NewServer(config, contactHTTPHandler)
		if err != nil {
			log.Fatal(err)
		}
	default:
		contactCLIHandler := handler.NewContactHandler(contactUC)
		handler.Menu(contactCLIHandler)
	}
}

func createContactUsecase(config *config.Config) usecase.ContactUsecase {
	var contactRepo repository.ContactRepository
	switch config.Storage {
	case "sql":
		switch config.Database.Driver {
		case "mysql":
			db, err := db.NewMysqlDatabase(config)
			if err != nil {
				log.Fatal(err)
			}
			contactRepo = repository.NewContactMysqlRepository(db)
		default:
			log.Fatalln("database driver not existed")
		}
	case "json":
		contactRepo = repository.NewContactJsonRepository()
	default:
		contactRepo = repository.NewContactRepository()
	}
	return usecase.NewContactUsecase(contactRepo)
}

func NewServer(config *config.Config, handler handler.ContactHTTPHandler) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch r.Method {
		case "GET":
			handler.List(w, r)
		case "POST":
			handler.Add(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	mux.HandleFunc("/contacts/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch r.Method {
		case "GET":
			handler.Detail(w, r)
		case "PATCH":
			handler.Update(w, r)
		case "DELETE":
			handler.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	server := &http.Server{
		Addr:    "localhost:" + config.Port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	log.Printf("live on http://localhost:%s", config.Port)
	return nil
}
