package main

import (
	"contact-go/config"
	"contact-go/handler"
	"contact-go/repository"
	"contact-go/usecase"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func main() {

	contactUseCase := usecase.NewContactUseCase()
	contactHandlerHttp := repository.NewContactRepository()
	contactHandler := handler.NewContactHandlerHttp(contactHandlerHttp, contactUseCase)
	NewServer(contactHandler)

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

func getConnectionShow() *sql.DB {
	connString := "root:@tcp(localhost:3306)/golang-trainee"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
