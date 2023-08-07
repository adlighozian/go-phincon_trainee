package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Result struct {
	Status  int
	Method  string
	Message string
}

type Example struct {
	Name string `json:"name"`
	Umur int    `json:"umur"`
}

func main() {
	app := fiber.New()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("🥇 First handler")
		logger.Info("failed to fetch URL",
			// Structured context as strongly typed Field values.
			zap.String("url", ""),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
		return c.Next()
	})

	// routing
	app.Get("/", func(c *fiber.Ctx) error {

		hasil := Result{
			Status:  http.StatusOK,
			Method:  c.Method(),
			Message: "selamat",
		}

		jsonMarshal, err := json.Marshal(hasil)
		if err != nil {
			log.Println("error")
		}

		fmt.Println("2")

		return c.Send(jsonMarshal)
	})

	app.Post("/", func(c *fiber.Ctx) error {

		var body Example
		err := c.BodyParser(&body)
		if err != nil {
			log.Println("error")
		}

		fmt.Println(body)

		return c.SendString("ini post")
	})

	// params
	app.Get("/param/:name", func(c *fiber.Ctx) error {

		hasil := Result{
			Status:  http.StatusOK,
			Method:  c.Method(),
			Message: c.Params("name"),
		}

		nama := c.Params("name")
		header := c.Get("Ambil")

		fmt.Println("nama saya", nama)
		fmt.Println("tanggal", header)

		return c.JSON(hasil)
	})

	// Static files
	app.Static("/folder", "./public/asd.txt")

	app.Static("/rumah", "./view/")

	// set header
	app.Get("/set-header/:key", func(c *fiber.Ctx) error {

		c.Set("kunci", c.Params("key"))
		headers := c.GetRespHeader("kunci")

		fmt.Println(headers)
		return c.SendString(headers)

	})

	//  run server
	log.Fatal(app.Listen(":3000"))
}
