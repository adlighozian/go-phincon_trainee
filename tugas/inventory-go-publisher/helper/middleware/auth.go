package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header.Get("auth") == "adli" {
			c.Next()
		} else {
			c.Abort()
			c.Redirect(http.StatusMovedPermanently, "/auth")
		}
	}

}
