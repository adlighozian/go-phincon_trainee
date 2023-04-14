package main

import (
	"fmt"
	"middleware/middleware"
	"net/http"
)

func main() {
	fmt.Println("Server running...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", HandlerGet)

	logMiddleWare := new(middleware.LogMiddleWare)
	logMiddleWare.Handler = mux

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: logMiddleWare,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func HandlerGet(write http.ResponseWriter, request *http.Request) {
	write.WriteHeader(http.StatusOK)
	fmt.Fprintf(write, "Success ini index")
}
