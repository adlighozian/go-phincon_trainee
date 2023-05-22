package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header.Get("auth") != "1221" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("not authenticate"))
		} else {
			c.Next()
		}
	}

}
