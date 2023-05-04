package main

import (
	"contact-go/config"
	"contact-go/db"
	"contact-go/handler"
	"contact-go/helper"
	"contact-go/repository"
	"contact-go/usecase"
	"fmt"
	"net/http"
)

func main() {

	dbms, _ := helper.GetEnv("dbms", "mysql")
	db, err := db.GetDB(dbms).GetConnectionMysql()
	if err != nil {
		panic(err)
	}

	contactHandlerHttp := repository.NewContactRepository(db, dbms)
	contactUseCase := usecase.NewContactUseCase(contactHandlerHttp)
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
	fmt.Println("Server running localhost:", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

// func getConnectionShow() *sql.DB {
// 	connString := "root:@tcp(localhost:3306)/golang-trainee"
// 	db, err := sql.Open("mysql", connString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.SetMaxIdleConns(10)
// 	db.SetMaxOpenConns(100)
// 	db.SetConnMaxIdleTime(5 * time.Minute)
// 	db.SetConnMaxLifetime(60 * time.Minute)
// 	return db
// }
