package routes

import (
	"gofiber_boilerplate/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// user grouping
	// user := app.Group("/users")
	// user.Post("/users", controllers.CreateUser)

	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:userId", controllers.GetUserDetails)
	app.Get("/users", controllers.UserList)
	app.Put("/users/:userId", controllers.UpdateUser)
	app.Delete("/users/:userId", controllers.DeleteUser)
}
