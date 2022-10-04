package main

import (
	"log"
	"webApp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./template/views", ".html"),
	})

	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Declare & initialize logger

	routes.AppRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
