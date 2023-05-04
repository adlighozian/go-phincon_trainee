package main

import (
	client "contact-go/db"
	"contact-go/config"
	"contact-go/handler"
	"contact-go/helper"
	"contact-go/repository"
	"contact-go/usecase"
	"fmt"
	"net/http"
	"os"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	switch config.Storage {
	case "json":
		contactRepo := repository.NewContactJsonRepository()
		contactHandler := handler.NewContactHttpJsonHandler(contactRepo)
		HTTPServer(config, contactHandler)
	case "mysql":
		db := client.GetDB(config.Storage).GetMysqlConnection()
		contactRepo := repository.NewContactHTTPRepository(db)
		useCase := usecase.NewUseCase(contactRepo)
		contacHandler := handler.NewContactHttpDbHandler(useCase)
		HTTPDBServer(config, contacHandler)
	default: // cmd
		contactRepo := repository.NewContactRepository()
		contactHandler := handler.NewContactHandler(contactRepo)
		Menu(contactHandler)
	}
}

func HTTPDBServer(config *config.Config, contactHandler handler.ContactHttpDbHandlerInterface) {
	mux := http.NewServeMux()
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			contactHandler.List(w, r)
		} else if r.Method == http.MethodPost {
			contactHandler.Add(w, r)
		} else if r.Method == http.MethodPatch {
			contactHandler.Update(w, r)
		} else if r.Method == http.MethodDelete {
			contactHandler.Delete(w, r)
		}
	})

	server := http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	fmt.Println("Server run on", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func HTTPServer(config *config.Config, contactHandler handler.ContactHttpJsonHandlerInterface) {
	mux := http.NewServeMux()
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			contactHandler.List(w, r)
		} else if r.Method == http.MethodPost {
			contactHandler.Add(w, r)
		} else if r.Method == http.MethodPatch {
			contactHandler.Update(w, r)
		} else if r.Method == http.MethodDelete {
			contactHandler.Delete(w, r)
		}
	})

	server := http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	fmt.Println("Server run on ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

// Menu function only for CLI used
func Menu(contactHandler handler.ContactHandlerInterface) {
	fmt.Println("\nSelect menu")
	fmt.Println("1. List contact")
	fmt.Println("2. Add contact")
	fmt.Println("3. Update contach")
	fmt.Println("4. Delete contach")
	fmt.Println("5. Exit")

	var choose int
	fmt.Print("Select menu : \n")
	fmt.Scanln(&choose)

	switch choose {
	case 1:
		contactHandler.List()
		helper.ClearScreeen()
		Menu(contactHandler)
	case 2:
		contactHandler.Add()
		contactHandler.List()
		helper.ClearScreeen()
		Menu(contactHandler)
	case 3:
		contactHandler.List()
		contactHandler.Update()
		fmt.Printf("------------Updated Datas------------")
		contactHandler.List()
		helper.ClearScreeen()
		Menu(contactHandler)
	case 4:
		contactHandler.List()
		contactHandler.Delete()
		fmt.Printf("------------Updated Datas------------")
		contactHandler.List()
		helper.ClearScreeen()
		Menu(contactHandler)
	case 5:
		os.Exit(1)
	}
}
