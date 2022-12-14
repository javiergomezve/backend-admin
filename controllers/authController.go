package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/javiergomezve/backend-admin/database"
	"github.com/javiergomezve/backend-admin/models"
	"github.com/javiergomezve/backend-admin/util"
)

func RegisterUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect email or password",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect email or password",
		})
	}

	token, err := util.GenerateJwt(user)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"token":   token,
	})
}

func User(c *fiber.Ctx) error {
	id := util.ParseJwt(c)

	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id := util.ParseJwt(c)
	fmt.Println(id)

	user := models.User{
		Id:        uint(id),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Debug().Model(&user).Updates(user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	id := util.ParseJwt(c)

	user := models.User{
		Id: uint(id),
	}
	user.SetPassword(data["password"])
	database.DB.Model(&user).Where("id", id).Updates(user)

	return c.JSON(user)
}
