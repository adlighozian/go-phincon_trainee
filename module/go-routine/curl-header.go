package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://reqres.in/api/users"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}
