package routes

import (
	"webApp/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	apiEndpoint := app.Group("/api")

	apiEndpoint.Get("/login", controller.ViewLogin)
	apiEndpoint.Post("/verify", controller.VerifyAccount)
	//User
	userEndpoint := apiEndpoint.Group("/user")
	userEndpoint.Post("/test", controller.Test)
}
