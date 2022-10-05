package routes

import (
	"webApp/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	apiEndpoint := app.Group("/api")

	//User
	userEndpoint := apiEndpoint.Group("/user")
	userEndpoint.Get("/get_user", controller.GetUsers)
	userEndpoint.Post("/login", controller.ViewLogin)
	userEndpoint.Post("/verify", controller.VerifyAccount)
}
