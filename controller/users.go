package controller

import (
	"webApp/model"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	userModel := []model.TbUsers{}

	DBConn.Debug().Table("tb_users").Find(&userModel)

	// return c.JSON(userModel)
	return c.Render("index", fiber.Map{
		"users": userModel,
	})

}
