package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Logger() gin.HandlerFunc	 {
	return func(c *gin.Context) {
		start := time.Now()

		// Create a logger instance for the request
		logger := log.With().
			Str("requestId", c.Request.Header.Get("X-Request-ID")).
			Logger()

		// Log the request information
		logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("ip", c.ClientIP()).
			Msg("Incoming request")

		// Proceed with the request handling
		c.Next()

		// Log the response information
		logger.Info().
			Int("statusCode", c.Writer.Status()).
			Dur("responseTime", time.Since(start)).
			Msg("Outgoing response")
	}
}
