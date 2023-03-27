package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

func formPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")
	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}
func main() {
	requestBody := strings.NewReader("first_name=Adli&last_name=Ghozian")
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/", requestBody)

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()
	formPost(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
