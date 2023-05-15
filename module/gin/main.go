package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//routing
	r.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "home",
		})
	})

	// query parameter, path parameter
	r.GET("/client", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"Message": "hello " + name,
		})
	})

	// set cookies
	r.GET("/set-cookies", func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:     "my-cookie",
			Value:    "hello world",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success set cookie",
		})
	})

	// get cookies
	r.GET("/get-cookies", func(c *gin.Context) {
		cookie, _ := c.Cookie("my-cookie")
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"cookies": cookie,
		})
	})

	// set header
	r.GET("/set-header", func(c *gin.Context) {
		header := "ini header"
		c.Writer.Header().Set("X-Custom-Header", header)
		c.Writer.Header().Add("Cache-Control", "no-cache")
		c.Writer.Header().Add("Cache-Control", "max-age=3600")

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"header": header,
		})
	})

	// get header
	r.GET("/get-header", func(c *gin.Context) {

		header := c.GetHeader("Content-Type")
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"header": header,
		})
	})

	// redirect
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/home")
		// c.Redirect(http.StatusOK, "/home")
	})

	// form post
	r.POST("/submit", func(c *gin.Context) {
		name := c.PostForm("name")
		telp := c.PostForm("telp")

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"name":   name,
			"telp":   telp,
		})
	})

	// Serve static files
	r.Static("/folder", "./folder/index.html")

	// Serve a single file
	r.StaticFile("/folders", "./folder")

	r.POST("/upload", func(c *gin.Context) {
		// Get the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{
				"error": "File not found",
			})
			return
		}

		// Save the file to disk
		err = c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to save file",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("File %s uploaded", file.Filename),
		})
	})

	//menjalankan server
	r.Run(":5000")
}
