package routes

import "github.com/gofiber/fiber/v2"

func AppRoutes(app *fiber.App) {
	apiEndpoint := app.Group("/api")

	//User
	userEndpoint := apiEndpoint.Group("/user")
	userEndpoint.Get("/get_user")
}
