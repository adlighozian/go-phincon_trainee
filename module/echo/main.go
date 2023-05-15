package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", getAllUser)
	e.GET("/users/:id", getUser)
	e.POST("/users", saveUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.GET("/set-cookies", setCookie)
	e.GET("/get-cookies", readCookie)

	e.GET("/set-header", handleGetUsers)
	e.GET("/get-header", getHeader)

	e.GET("/redirect", redirect)

	e.File("/serve-file", "file/index.html")

	e.Static("/file-serve", "file")

	e.GET("/file-download", downloadFile)
	e.POST("/file-upload", uploadFile)

	e.Logger.Fatal(e.Start(":1234"))
}

// e.POST("/file-upload", downloadFile)
func uploadFile(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}

func downloadFile(c echo.Context) error {

	// Set attachment header to download the file
	c.Response().Header().Set("Content-Disposition", "attachment; filename=text.txt")

	return c.File("text.txt")
	// return c.Inline("inline.txt", "inline.txt")
	// return c.Attachment("attachment.txt", "attachment.txt")
}

// e.GET("/users", getAllUser)
func getAllUser(c echo.Context) error {
	nama := c.QueryParam("nama")
	noTelp := c.QueryParam("no_telp")
	return c.String(http.StatusOK, "nama: "+nama+", no HP: "+noTelp)
}

func redirect(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/")
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "id: "+id)
}

// e.POST("/users", saveUser)
func saveUser(c echo.Context) error {
	// Retrieve the form values from the request
	fmt.Println("submit-form")
	name := c.FormValue("name")
	email := c.FormValue("email")

	// Do something with the form data
	// ...

	// Redirect the user to a new page
	return c.JSON(http.StatusOK, "name: "+name+", email: "+email)
}

func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "ini show")
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "ini show")
}

// set cookies
func setCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "adli"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "cookies :"+cookie.Value)
}

// set cookies
func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie, "+"cookies: "+cookie.Name)
}

// set header
func handleGetUsers(c echo.Context) error {
	users := []string{"Alice", "Bob", "Charlie"}

	// Set the content type header to JSON
	c.Response().Header().Set("dummy", "header")

	// Return the users as a JSON response
	return c.JSON(http.StatusOK, users)
}

func getHeader(c echo.Context) error {
	// Get the value of the "Authorization" header
	header := c.Request().Header.Get("Date")
	content := c.Request().Header.Get("Content-Type")

	fmt.Println(header)
	fmt.Println(content)
	// Do something with the header value
	// ...

	return c.JSON(http.StatusOK, "Header value: "+header)

	// for key, values := range c.Request().Header {
	// 	fmt.Println(key)
	// 	for _, value := range values {
	// 		fmt.Println(value)
	// 	}
	// }
	// return nil
}
