package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			var temp []byte
			// var temp_ map[string]interface{}

			if r.Body != nil {

				temp, _ = io.ReadAll(r.Body)
			}
			if r.URL.Query().Get("id") != "" {
				temp = []byte(r.URL.Query().Get("id"))
			}
			w.Write(temp)
		} else {
			w.WriteHeader(http.StatusPermanentRedirect)
			http.Redirect(w, r, "/landing", http.StatusPermanentRedirect)
		}

	})
	mux.HandleFunc("/landing", func(w http.ResponseWriter, r *http.Request) {
		cookie := new(http.Cookie)
		cookie.Name = "test"
		w.Write([]byte("redirect"))
	})

	// file server
	directory := http.Dir("view")
	fs := http.FileServer(directory)
	mux.Handle("/serve/", http.StripPrefix("/serve/", fs))

	// directory := http.Dir("view")
	// fs := http.FileServer(directory)
	// mux.Handle("/serve/", fs)

	// serve file
	mux.HandleFunc("/servefile", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./view/index.html")
	})

	mux.HandleFunc("/uploadFile", func(w http.ResponseWriter, r *http.Request) {
		file, fileheader, err := r.FormFile("gambar")
		if err != nil {
			panic(err)
		}

		destfile, err := os.Create("./files/" + fileheader.Filename)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(destfile, file)
		if err != nil {
			panic(err)
		}

		w.Write([]byte(fileheader.Filename + "telah diupload"))

	})

	mux.HandleFunc(download())

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}
	log.Println("running on", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func download() (string, func(w http.ResponseWriter, r *http.Request)) {
	return "/downloadfile", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Query().Get("file")
		if filename == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request"))
		}
		w.Header().Add("content-disposition", "attachment;filename=\""+filename+"\"")
		http.ServeFile(w, r, "./files/"+filename)
	}
}
