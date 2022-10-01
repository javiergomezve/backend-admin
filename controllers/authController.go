package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiergomezve/backend-admin/models"
)

func RegisterUser(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "Javier",
		LastName:  "Gomez",
	}

	return c.JSON(user)
}
