package controller

import "github.com/gofiber/fiber/v2"

// CONTROLLER

// VIEWS
func ViewUpload(c *fiber.Ctx) error {
	return c.Render("uploadfile", fiber.Map{
		"page":  "Upload File",
		"title": "UPLOAD FILE",
	})
}
