package middleware

import (
	"fmt"
	"net/http"
	"time"

	logger "sales-go/helpers/logging"
)

func LoggingHandler(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof(fmt.Sprintf("Started %s localhost:5000%s", r.Method, r.URL.Path))

		fmt.Println("MIDDLEWARE PASS 1")
		mux.ServeHTTP(w, r)
		fmt.Println("MIDDLEWARE PASS 2")

		// handle panic error middleware
		defer func() {
			fmt.Println("MIDDLEWARE PASS 3")
			err := recover()
			if err != nil {
				fmt.Println("MIDDLEWARE PASS 4")
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
