package middleware

import (
	"fmt"
	"log"
	"net/http"
)

type Middlewares struct {
	Handler     http.Handler
	Middlewares []func(http.Handler) http.Handler
}

func

func (m *Middlewares) use(middlewares ...func(http.Handler) http.Handler) {
	m.Middlewares = append(m.Middlewares, middlewares...)
}

func (m *Middlewares) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before")
	m.Handler.ServeHTTP(w, r)
	fmt.Println("After")

	var h http.Handler = &m.ServerMux
	for _, next := range m.Middlewares {
		h = next(h)
	}
	h.ServeHTTP(w, r)
}

func AtasBawah(h http.Handler) http.Handler {
	log.Println("atas")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Println("bawah")
	})
}

func BeforeAfter(h http.Handler) http.Handler {
	log.Println("Before")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		defer log.Println("After")
	})
}
