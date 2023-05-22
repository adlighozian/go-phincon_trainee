package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("key") == "adli" {
			c.Redirect(http.StatusMovedPermanently, "http://www.youtube.com")
		} else {
			c.Redirect(http.StatusMovedPermanently, "http://www.google.com")
		}
	}
}

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before run")
		defer func() {
			err := recover()
			if err != nil {
				log.Println(err)
			}
		}()
		if c.Request.Header.Get("key") == "adli" {
			c.Next()
			fmt.Println("after run")
		} else {
			fmt.Println("redirect")
			c.Abort()
			c.Redirect(http.StatusMovedPermanently, "/")
		}

	}
}
