package main

import (
	"log"
	"webApp/controller"
	"webApp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {

	controller.CreateConnection()

	app := fiber.New(fiber.Config{
		Views: html.New("./template/views", ".html"),
	})

	// Load static files like CSS, Images & JavaScript
	app.Static("/", "./template/css")
	app.Static("/", "./template/js")
	app.Static("/", "./template/images")

	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Declare & initialize logger
	routes.AppRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
