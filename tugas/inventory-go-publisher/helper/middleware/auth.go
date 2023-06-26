package middleware

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header.Get("auth") == "1221" {
			c.Next()
		} else {
			c.Abort()
			c.Redirect(http.StatusMovedPermanently, "/auth")
		}
	}

}

func HeaderVerificationMiddleware(ginCtx *gin.Context) {
	hashKeyStr := ginCtx.GetHeader("key")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:5000/auth", nil)
	if err != nil {
		log.Println("error 1")
		ginCtx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	req.Header.Add("key", hashKeyStr)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error 2")
		ginCtx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ginCtx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		ginCtx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if response["message"] == "Gagal" {
		ginCtx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
