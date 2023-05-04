package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func LoggingHandler(handler http.Handler) http.Handler {
	start := time.Now()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// start set CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}
		// end set CORS

		// start zero logger
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}
		multi := zerolog.MultiLevelWriter(consoleWriter)
		logger := zerolog.New(multi).With().Timestamp().Logger()
		// end zero logger

		// log before service
		handler.ServeHTTP(w, r)
		// log after service

		// start zero logger
		latency := time.Since(start)
		logger.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Str("user_agent", r.UserAgent()).
			Str("referer", r.Referer()).
			Str("proto", r.Proto).
			Str("remote_ip", r.RemoteAddr).
			Dur("latency", latency).
			Msg("")
		// end zero logger
	})
}

func Use(middleware ...http.Handler) []http.Handler {
	var handlers []http.Handler
	handlers = append(handlers, middleware...)
	return handlers
}
