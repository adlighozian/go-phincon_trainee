package middleware

import (
	"fmt"
	"net/http"
	"time"

	logger "inventory-go-http-database/middleware/logging"
)

func LoggingHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof(fmt.Sprintf("Started %s localhost:5000%s", r.Method, r.URL.Path))

		handler.ServeHTTP(w, r)

		// handle error
		defer func() {
			err := recover()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Errorf(err.(error))
			} else {
				logger.Infof(fmt.Sprintf("Completed %s localhost:5000%s in %v", r.Method, r.URL.Path, time.Since(time.Now())))
			}
		}()
	})
}

func Use(middleware ...http.Handler) []http.Handler {
	var handlers []http.Handler
	handlers = append(handlers, middleware...)
	return handlers
}
