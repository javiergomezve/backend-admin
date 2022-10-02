package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiergomezve/backend-admin/controllers"
	"github.com/javiergomezve/backend-admin/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.RegisterUser)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated())

	app.Get("/api/user", controllers.User)
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
}
