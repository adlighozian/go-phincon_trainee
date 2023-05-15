# 12 thing Gin must do

## 1. How to make a server

To create a server using Gin, you'll first need to import the Gin package and then create a router instance. The router instance is responsible for handling incoming HTTP requests and routing them to the appropriate handlers. Here's a basic example:

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

In this example, we create a default Gin router using `gin.Default()`, define a route `/ping` that responds with a JSON object, and start the server using `r.Run()` which listens on port 8080 by default (github.com).

## 2. Routing - handlers

Routing in Gin is done by defining routes and their corresponding handlers. Handlers are functions that process incoming requests and generate responses. You can define routes using methods like `GET`, `POST`, `PUT`, `DELETE`, etc., on the Gin router instance. Here's an example:

```go
func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    r.POST("/submit", func(c *gin.Context) {
        name := c.PostForm("name")
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello " + name,
        })
    })

    r.Run()
}
```

In this example, we define a `GET` route `/ping` and a `POST` route `/submit`. The handlers for each route process the requests and respond with JSON objects (github.com).

## 3. Request - response / reader - writer

In Gin, you can read request data and write response data using the `*gin.Context` object provided to the handler. The `*gin.Context` object contains methods to read request data (e.g., `PostForm`, `QueryParam`, `DefaultPostForm`, `DefaultQuery`) and write response data (e.g., `JSON`, `XML`, `String`, `HTML`). Here's an example:

```go
func main() {
    r := gin.Default()

    r.POST("/submit", func(c *gin.Context) {
        name := c.PostForm("name")
        age := c.DefaultPostForm("age", "18")
        c.JSON(http.StatusOK, gin.H{
            "name": name,
            "age":  age,
        })
    })

    r.Run()
}
```

In this example, we read the `name` and `age` fields from a `POST` request and respond with a JSON object containing the submitted data (github.com).

## 4. Query and path parameters

Gin allows you to access query parameters and path parameters from a request. Query parameters are part of the URL after the `?` symbol, while path parameters are part of the URL path itself. Here's an example:

```go
func main() {
    r := gin.Default()

    r.GET("/user/:id", func(c *gin.Context) {
        id := c.Param("id")
        action := c.DefaultQuery("action", "view")
        c.JSON(http.StatusOK, gin.H{
            "id":     id,
            "action": action,
        })
    })

    r.Run()
}
```

In this example, we define a route `/user/:id` with a path parameter `id`. We also read a query parameter `action` with a default value of `"view"`. The handler responds with a JSON object containing the path parameter and query parameter values (github.com).

## 5. Set, get, delete headers

You can set, get, and delete headers in the `*gin.Context` object. Here's an example:

```go
func main() {
    r := gin.Default()

    r.GET("/headers", func(c *gin.Context) {
        // Get header
        userAgent := c.GetHeader("User-Agent")

        // Set header
        c.Header("Content-Type", "application/json")

        // Delete header
        c.Header("X-Deleted-Header", "")
        c.Header("X-Deleted-Header", "")

        c.JSON(http.StatusOK, gin.H{
            "User-Agent": userAgent,
        })
    })

    r.Run()
}
```

In this example, we get the `User-Agent` header from the request, set a `Content-Type` header in the response, and delete a header by setting its value to an empty string. The handler responds with a JSON object containing the `User-Agent` header value ([Source 0](https://github.com/gin-gonic/gin)).

## 6. Set, get, delete cookies

Gin allows you to set, get, and delete cookies using the `*gin.Context` object. Here's an example:

```go
func main() {
    r := gin.Default()

    r.GET("/cookie", func(c *gin.Context) {
        // Get cookie
        cookie, err := c.Cookie("my_cookie")

        if err != nil {
            // Set cookie
            c.SetCookie("my_cookie", "my_value", 3600, "/", "", false, true)
            c.JSON(http.StatusOK, gin.H{
                "message": "Cookie set",
            })
        } else {
            // Delete cookie
            c.SetCookie("my_cookie", "", -1, "/", "", false, true)
            c.JSON(http.StatusOK, gin.H{
                "message": "Cookie deleted",
                "value":   cookie,
            })
        }
    })

    r.Run()
}
```

In this example, we get a cookie named `my_cookie`. If the cookie is not present, we set it with a value and an expiration time. If the cookie is present, we delete it by setting its expiration time to a negative value ([Source 0](https://github.com/gin-gonic/gin)).

## 7. How to set a response code

You can set the response code for a request using the `*gin.Context` object. The `JSON`, `XML`, `String`, and `HTML` methods accept a response code as their first argument. Here's an example:

```go
func main() {
    r := gin.Default()

    r.GET("/not_found", func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "message": "Resource not found",
        })
    })

    r.Run()
}
```

In this example, we set the response code to `http.StatusNotFound` (404) when responding with a JSON object ([Source 0](https://github.com/gin-gonic/gin)).

## 8. Redirects

Gin allows you to perform redirects using the `*gin.Context` object. You can use the `Redirect` method to redirect a request to a new URL. Here's an example:

```go
func main() {
    r := gin.Default()

    r.GET("/old", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/new")
    })

    r.GET("/new", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome to the new URL!")
    })

    r.Run()
}
```

In this example, we redirect requests from `/old` to `/new` with a `301 Moved Permanently` status code ([Source 0](https://github.com/gin-gonic/gin)).

## 9. Form post

You can handle form submissions in Gin using the `*gin.Context` object. The `PostForm` and `DefaultPostForm` methods allow you to read form data from a `POST` request. Here's an example:

```go
func main() {
    r := gin.Default()

    r.POST("/submit", func(c *gin.Context) {
        name := c.PostForm("name")
        age := c.DefaultPostForm("age", "18")
        c.JSON(http.StatusOK, gin.H{
            "name": name,
            "age":  age,
        })
    })

    r.Run()
}
```

In this example, we read the `name` and `age` fields from a `POST` request and respond with a JSON object containing the submitted data ([Source 0](https://github.com/gin-gonic/gin)).

## 10. File serve and serve file

Gin allows you to serve static files and serve individual files using the `Static` and `StaticFile` methods. Here's an example:

```go
func main() {
    r := gin.Default()

    // Serve static files
    r.Static("/assets", "./assets")

    // Serve a single file
    r.StaticFile("/favicon.ico", "./favicon.ico")

    r.Run()
}
```

In this example, we serve static files from the `./assets` directory at the `/assets` URL and serve a single file `./favicon.ico` at the `/favicon.ico` URL ([Source 0](https://github.com/gin-gonic/gin)).

## 11. Uploads and downloads

Gin allows you to handle file uploads and downloads using the `*gin.Context` object. Here's an example of handling file uploads:

```go
func main() {
    r := gin.Default()

    r.POST("/upload", func(c *gin.Context) {
        file, _ := c.FormFile("file")
        c.SaveUploadedFile(file, "./uploads/"+file.Filename)

        c.JSON(http.StatusOK, gin.H{
            "message": "File uploaded successfully",
            "file":    file.Filename,
        })
    })

    r.Run()
}
```

In this example, we handle file uploads by reading the uploaded file from a `POST` request and saving it to the `./uploads` directory (gin-gonic.com).

For handling file downloads, you can use the `File` method of `*gin.Context`. Here's an example:

```go
const DOWNLOADS_PATH = "downloads/"

func main() {
    r := gin.Default()

    r.GET("/download/:filename", func(c *gin.Context) {
        fileName := c.Param("filename")
        targetPath := filepath.Join(DOWNLOADS_PATH, fileName)

        if !strings.HasPrefix(filepath.Clean(targetPath), DOWNLOADS_PATH) {
            c.String(403, "Looks like you are attacking me")
            return
        }

        c.Header("Content-Description", "File Transfer")
        c.Header("Content-Transfer-Encoding", "binary")
        c.Header("Content-Disposition", "attachment; filename="+fileName)
        c.Header("Content-Type", "application/octet-stream")
        c.File(targetPath)
    })

    r.Run()
}
```

In this example, we serve a file for download by reading the file from the `./downloads` directory and sending it as a response. The `Context.File` method is used to serve the file, and we set the necessary headers for the file transfer (stackoverflow.com).

I apologize for the confusion. Here's the correct example for using middleware with Gin:

## 12. Middleware

In Gin, middleware can be added globally, per route, or per group of routes. Middleware functions are executed in the order they are added, and they can be used to perform various operations such as logging, authentication, or input validation. Here's an example of using middleware with Gin:

```go
func main() {
    // Creates a router without any middleware by default
    r := gin.New()

    // Global middleware
    // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
    // By default gin.DefaultWriter = os.Stdout
    r.Use(gin.Logger())

    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    r.Use(gin.Recovery())

    // Per route middleware, you can add as many as you desire.
    r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

    // Authorization group
    authorized := r.Group("/")
    // per group middleware! in this case we use the custom created
    // AuthRequired() middleware just in the "authorized" group.
    authorized.Use(AuthRequired())
    {
        authorized.POST("/login", loginEndpoint)
        authorized.POST("/submit", submitEndpoint)
        authorized.POST("/read", readEndpoint)

        // nested group
        testing := authorized.Group("testing")
        testing.GET("/analytics", analyticsEndpoint)
    }

    // Listen and serve on 0.0.0.0:8080
    r.Run(":8080")
}
```

In this example, we create a new Gin router and add global middleware using `r.Use()`. We then add per-route middleware by passing the middleware functions as arguments to the `r.GET()` method. Finally, we create an authorized group of routes and add the custom `AuthRequired()` middleware to that group. Nested groups can also have their middleware ([Source 1](https://gin-gonic.com/docs/examples/using-middleware/)).

## 13. Custom TLS configuration

You can use a custom TLS configuration with Gin by creating a certificate, signing it, and configuring the `http.Server` with the generated certificate and key. Here's an example of how to use a custom certificate with Gin:

```go
import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/tls"
    "crypto/x509"
    "crypto/x509/pkix"
    "encoding/pem"
    "math/big"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Your routes here...

    cert := &x509.Certificate{
        SerialNumber: big.NewInt(1658),
        Subject: pkix.Name{
            Organization:  []string{"ORGANIZATION_NAME"},
            Country:       []string{"COUNTRY_CODE"},
            Province:      []string{"PROVINCE"},
            Locality:      []string{"CITY"},
            StreetAddress: []string{"ADDRESS"},
            PostalCode:    []string{"POSTAL_CODE"},
        },
        NotBefore:    time.Now(),
        NotAfter:     time.Now().AddDate(10, 0, 0),
        SubjectKeyId: []byte{1, 2, 3, 4, 6},
        ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
        KeyUsage:     x509.KeyUsageDigitalSignature,
    }
    priv, _ := rsa.GenerateKey(rand.Reader, 2048)
    pub := &priv.PublicKey

    // Sign the certificate
    certificate, _ := x509.CreateCertificate(rand.Reader, cert, cert, pub, priv)

    certBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificate})
    keyBytes := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

    // Generate a key pair from your pem-encoded cert and key ([]byte).
    x509Cert, _ := tls.X509KeyPair(certBytes, keyBytes)

    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{x509Cert},
    }
    server := http.Server{Addr: ":3000", Handler: router, TLSConfig: tlsConfig}

    server.ListenAndServeTLS("", "")
}
```

This example creates a self-signed certificate and uses it to configure the `http.Server` for TLS ([Source 0](https://stackoverflow.com/questions/67625752/how-to-use-a-certificate-from-a-certificate-store-and-run-tls-in-gin-framework-i)).
