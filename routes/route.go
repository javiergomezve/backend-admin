package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiergomezve/backend-admin/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
