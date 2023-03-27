package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func RequestHeaderHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprintln(w, contentType)
	fmt.Fprintln(w, "contentType")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello home")
	})

	request := httptest.NewRequest("get", "http://http://localhost:8080/hello", nil)
	request.Header.Add("Content-Type", "application/json")

	// requestNew := httptest.NewRequest(http.MethodGet, "http://http://localhost:8080/hello", nil)
	
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)
	// RequestHeaderHandler(recorder, requestNew)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

}
