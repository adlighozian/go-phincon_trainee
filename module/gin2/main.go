package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rabbitmq/middlewares"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `form:"name"`
}

// 1. membuat server
func main() {
	r := gin.New()

	// r.Use(middlewares.MyMiddleware(), gin.Recovery())

	rs := r.Group("/group1", middlewares.MyMiddleware())

	// cara 1
	r.GET("/", newGet)
	r.GET("/getall", getAll)
	r.GET("/user/respon", responRequest)
	r.GET("/user/", queryParams)
	rs.GET("/user/:name", pathParams)
	r.GET("/set-header", setHeader)
	r.GET("/set-cookie", setCookie)
	r.GET("/redirect", wrongEndPoint)
	rs.POST("/upload", formPost)
	r.Static("/static", "./upload")
	r.GET("/download", download)
	r.GET("/getuser", GetUser)

	r.Run()
}

// ==============================================================================================

// 2. routing
func newGet(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world")
	c.Writer.Write([]byte("redirect"))
}

// ==============================================================================================

// 3. request response - write reader
func responRequest(c *gin.Context) {

	c.JSON(http.StatusOK, "hello world")
}

// c.writer ---> respons
// c.Reader ---> request

// ==============================================================================================

// 4. query parameter / path parameter
func queryParams(c *gin.Context) {
	// query parameter

	// cara 1
	name := c.Query("name")

	// cara 2
	var user User
	_ = c.BindQuery(&user)

	log.Println(user)

	c.JSON(http.StatusOK, name)
}

func pathParams(c *gin.Context) {
	// path parameter

	name := c.Param("name")

	c.JSON(http.StatusOK, name)
}

// ==============================================================================================

// 5. set dan get cookies
func setCookie(c *gin.Context) {
	c.SetCookie("token", "1221", 300, "", "", true, true) //set cookie dari respon

	c.Writer.Header().Get("token") //get cookie dari respon
	c.Cookie("token")              //get cookie dari request

}

// ==============================================================================================

// 6. set dan get Header
func setHeader(c *gin.Context) {
	c.Header("test", "dummy") //set header dari respon

	c.Writer.Header().Get("Content-type") //get dari respon
	c.GetHeader("Content-type")           //get dari request

}

// ==============================================================================================

// 7. redirect
func wrongEndPoint(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/")
}

// ==============================================================================================

// 8. form post
func formPost(c *gin.Context) {
	file := c.PostForm("key")
	file2, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, file)
	c.SaveUploadedFile(file2, "./upload/"+file2.Filename)

}

// ==============================================================================================

// 9. file serve / serve file

// func staticFile(c *gin.Context) {

// }

// ==============================================================================================

// 10. upload download

func download(c *gin.Context) {
	filename := "example.txt"
	fileLocation := "./example.txt"
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.File(fileLocation)
}

// ==============================================================================================

// 11. middleware

type user struct {
	name   string
	notelp string
}

func getAll(c *gin.Context) {
	// fmt.Println("getAll")
	// var temp map[string]interface{}
	var users []user = []user{
		{
			name:   "adli",
			notelp: "893874",
		},
		{
			name:   "ghozi",
			notelp: "893874",
		},
	}

	// jsonData, err := json.Marshal(users)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(users)
	c.JSON(http.StatusOK, users)
}

// hit another service
func GetUser(c *gin.Context) {
	res, err := http.NewRequest(http.MethodGet, "http://localhost:8080/getall", nil)
	if err != nil {
		panic(err)
	}

	var temp map[string]interface{}

	readAll, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(readAll, &temp)
	if err != nil {
		panic(err)
	}

	c.JSON(200, temp)

}
