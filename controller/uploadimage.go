package controller

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"webApp/controller/util"
	"webApp/model"

	"github.com/gofiber/fiber/v2"
)

// CONTROLLER
func ConvertToBase64(b []byte) string {
	fmt.Println("b: ", b)
	return base64.StdEncoding.EncodeToString(b)
}

func FileUpload(c *fiber.Ctx) error {

	imgSrc, uploadErr := c.FormFile("imageSource")
	if uploadErr != nil {
		return c.SendString("Error in FormFile")
	}
	imgPath := fmt.Sprintf("./template/images/%s", imgSrc.Filename)
	c.SaveFile(imgSrc, fmt.Sprintf("./template/images/%s", imgSrc.Filename))
	bytes, err := ioutil.ReadFile(imgPath)

	fmt.Println("Bytes:", string(imgPath))
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += ConvertToBase64(bytes)
	fmt.Println("base64: ", base64Encoding)
	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)

	imgStruct := model.Uploaded_Images{
		ImgData: base64Encoding,
	}

	util.DBConn.Debug().Create(&imgStruct)

	return c.JSON(base64Encoding)

}

// VIEWS
func UploadImage(c *fiber.Ctx) error {
	return c.Render("uploadimage", fiber.Map{
		"page":  "Testing Upload Image File",
		"title": "UPLOAD IMAGE",
	})
}
