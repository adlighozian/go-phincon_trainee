package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	//"net/httptest"
	"time"

	"sales-go/helpers/gin-rest"
	logger "sales-go/helpers/logging"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-location-lat, x-location-long, x-unique-id")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
        if r.Method == "OPTIONS" {
            w.Write([]byte("allowed"))
            return
        }
	})
}

func HeaderVerificationMiddleware(ginCtx *gin.Context) {
	hashKeyStr := ginCtx.GetHeader("key")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/verification", nil)
	if err != nil {
		ginCtx.AbortWithStatus(http.StatusInternalServerError)
		rest.ResponseError(ginCtx, http.StatusInternalServerError, err)
		return
	}
	req.Header.Add("key", hashKeyStr)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ginCtx.AbortWithStatus(http.StatusInternalServerError)
		rest.ResponseError(ginCtx, http.StatusInternalServerError, err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ginCtx.AbortWithStatus(http.StatusInternalServerError)
		rest.ResponseError(ginCtx, http.StatusInternalServerError, err)
		return
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		ginCtx.AbortWithStatus(http.StatusInternalServerError)
		rest.ResponseError(ginCtx, http.StatusInternalServerError, err)
		return
	}

	if response["message"] == "Unauthorized" {
		ginCtx.AbortWithStatus(http.StatusUnauthorized)
		rest.ResponseError(ginCtx, http.StatusBadRequest, fmt.Errorf("key is empty or not authorized"))
		return
	}
}

func LoggingMiddleware(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-location-lat, x-location-long, x-unique-id")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
        if r.Method == "OPTIONS" {
            w.Write([]byte("allowed"))
            return
        }

		logger.Infof(fmt.Sprintf("Started %s localhost:5000%s", r.Method, r.URL.Path), r)

		mux.ServeHTTP(w, r)

		// handle panic error middleware
		defer func() {
			fmt.Println("MIDDLEWARE PASS 1")
			err := recover()
			if err != nil {
				fmt.Println("MIDDLEWARE PASS 2")
				w.WriteHeader(http.StatusInternalServerError)
				logger.Errorf(err.(error), r)
			} else {
				logger.Infof(fmt.Sprintf("Completed %s localhost:5000%s in %v", r.Method, r.URL.Path, time.Since(time.Now())), r)
			}
		}()
	})
}

func Use(middleware ...http.Handler) []http.Handler {
	var handlers []http.Handler
	handlers = append(handlers, middleware...)
	return handlers
}
