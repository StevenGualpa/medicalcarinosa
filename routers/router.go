// routers/user_router.go
package routers

import (
	"GolandProyectos/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes configura las rutas relacionadas con los usuarios.
func SetupUserRoutes(app *fiber.App, userHandler handlers.UserHandler) {
	app.Get("/api/getAll", userHandler.GetUsers)
	app.Get("/api/users/:id", userHandler.GetUser)
	app.Post("/api/insert", userHandler.CreateUser)
	app.Put("/api/update/:id", userHandler.UpdateUser)
	app.Delete("/api/delete/:id", userHandler.DeleteUser)
	app.Post("/api/login", userHandler.Login)
}
