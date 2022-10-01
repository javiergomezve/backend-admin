package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/javiergomezve/backend-admin/database"
	"github.com/javiergomezve/backend-admin/models"
	"golang.org/x/crypto/bcrypt"
	"time"
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

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

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

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect email or password",
		})
	}

	claims := jwt.MapClaims{
		"id":   user.Id,
		"name": user.FirstName,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"token":   token,
	})
}

func User(c *fiber.Ctx) error {
	userEncode := c.Locals("user").(*jwt.Token)
	claims := userEncode.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	name := claims["name"].(string)
	fmt.Println(name)

	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}
