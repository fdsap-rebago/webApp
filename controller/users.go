package controller

import (
	"net/http"
	"strconv"
	"webApp/model"

	"github.com/gofiber/fiber/v2"
)

// CONTROLLER
func RegisterAccount(c *fiber.Ctx) error {
	// Get value from input tags
	ic := c.FormValue("instiCode")
	fn := c.FormValue("firstname")
	ln := c.FormValue("lastname")
	un := c.FormValue("username")
	pw := c.FormValue("password")

	cic, _ := strconv.Atoi(ic)

	userModel := model.TbUsers{
		Firstname: fn,
		Lastname:  ln,
		Username:  un,
		Password:  pw,
		InstiCode: cic,
	}

	if creatErr := DBConn.Debug().Table("tbl_users").Create(&userModel).Error; creatErr != nil {
		return creatErr
	}

	return c.Render("login", fiber.Map{
		"status": http.StatusOK,
	})

}

// VIEWS
func ViewLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"title": "USER LOGIN",
	})
}

func ViewRegistration(c *fiber.Ctx) error {
	instiModel := []model.M_Institution{}

	DBConn.Debug().Table("m_institution").Find(&instiModel)

	return c.Render("registration", fiber.Map{
		"title":        "USER REGISTRATION",
		"institutions": instiModel,
	})
}
