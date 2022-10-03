package util

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/javiergomezve/backend-admin/models"
)

const SecretKey = "secret"

func GenerateJwt(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":   user.Id,
		"name": user.FirstName,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(SecretKey))
	if err != nil {
		return "", errors.New("token couldn't be generated")
	}

	return token, nil
}

func ParseJwt(c *fiber.Ctx) float64 {
	userEncode := c.Locals("user").(*jwt.Token)
	claims := userEncode.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	//name := claims["name"].(string)

	return id
}
