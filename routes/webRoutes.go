package routes

import (
	"webApp/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	app.Get("/", controller.ViewLogin)
	apiEndpoint := app.Group("/api")

	//User
	userEndpoint := apiEndpoint.Group("/user")
	userEndpoint.Get("/register", controller.ViewRegistration)
	userEndpoint.Post("/save", controller.RegisterAccount)

}
