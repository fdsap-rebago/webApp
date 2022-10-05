package controller

import (
	"fmt"
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

func ViewLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

type LoginingIn struct {
	Username string `json:"username"`
}

func VerifyAccount(c *fiber.Ctx) error {
	userModel := model.TbUsers{}
	log := LoginingIn{}
	log.Username = c.FormValue("username")
	DBConn.Debug().Table("tb_users").Where("username=?", log.Username).Find(&userModel)
	if userModel.Username == "" {
		DBConn.Debug().Table("rbi.tb_users").Where("username=?", log.Username).Find(&userModel)
		if userModel.Username == "" {
			return c.SendString("user not found")
		}
	}
	fmt.Println("usermodel:", userModel)
	return c.JSON(userModel)
}

/*
"branch":"rbi"

*/
