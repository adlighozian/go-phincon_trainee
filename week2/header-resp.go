package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func responseHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "PhinCon")
	fmt.Fprint(w, "OK")
}
func main() {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	responseHeaderHandler(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-by"))
}
