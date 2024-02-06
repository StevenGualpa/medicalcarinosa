// routers/user_router.go
package routers

import (
	"GolandProyectos/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes configura las rutas relacionadas con los usuarios.
func SetupUserRoutes(app *fiber.App, userHandler handlers.UserHandler) {
	app.Get("/Users/GetAll", userHandler.GetUsers)
	app.Get("/api/users/:id", userHandler.GetUser)
	app.Post("/Users/Insert", userHandler.CreateUser)
	app.Put("/api/users/:id", userHandler.UpdateUser)
	app.Delete("/api/users/:id", userHandler.DeleteUser)
}
