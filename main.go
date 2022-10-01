package main

import (
	"github.com/javiergomezve/backend-admin/database"
	"github.com/javiergomezve/backend-admin/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3030")
}
