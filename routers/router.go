// routers/user_router.go
package routers

import (
	"GolandProyectos/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes configura las rutas relacionadas con los usuarios.
func SetupUserRoutes(app *fiber.App, userHandler handlers.UserHandler) {
	app.Get("/api/users", userHandler.GetUsers)
	app.Get("/api/users/:id", userHandler.GetUser)
	app.Post("/api/users", userHandler.CreateUser)
	app.Put("/api/users/:id", userHandler.UpdateUser)
	app.Delete("/api/users/:id", userHandler.DeleteUser)
}