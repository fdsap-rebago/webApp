package controller

import (
	"fmt"
	"webApp/model"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx, branch string) error {
	userModel := []model.TbUsers{}

	DBConn.Debug().Table(branch + ".tb_users").Find(&userModel)
	// return c.JSON(userModel)
	return c.Render("users", fiber.Map{
		"users": userModel,
		"title": branch,
	})
}

func ViewLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "User Login",
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
			DBConn.Debug().Table("sme.tb_users").Where("username=?", log.Username).Find(&userModel)
			if userModel.Username == "" {
				c.SendString("user not found")
			}
		}
	}

	fmt.Println("Branch:", userModel.Branch)
	GetUsers(c, userModel.Branch)

	return nil
}

func Test(c *fiber.Ctx) error {
	return nil
}
