// routers/user_router.go
package routers

import (
	"GolandProyectos/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes configura las rutas relacionadas con los usuarios.
func SetupUserRoutes(app *fiber.App, userHandler handlers.UserHandler) {
	app.Get("/api/getAll", userHandler.GetUsers)
	app.Post("/user/getAllRoles", userHandler.GetUsersRoles)
	app.Get("/api/users/:id", userHandler.GetUser)
	app.Post("/user/insertadmin", userHandler.CreateUser)
	app.Post("/user/insertcuidador", userHandler.CreateUser)
	app.Post("/user/insertpaciente", userHandler.CreateUser)
	app.Put("/api/update/:id", userHandler.UpdateUser)
	app.Delete("/api/delete/:id", userHandler.DeleteUser)
	app.Post("/api/login", userHandler.Login)

	// Aquí se agrega la nueva ruta para obtener dispositivos de Arduino
	app.Get("/api/arduino/devices", handlers.NewArduinoHandler().GetArduinoDevices)
}

// SetupPacienteCuidadorRoutes configura las rutas para la entidad PacienteCuidador.
func SetupPacienteCuidadorRoutes(app *fiber.App, pcHandler handlers.PacienteCuidadorHandler) {
	app.Post("/pacientecuidador/insert", pcHandler.CreatePacienteCuidador)
	app.Put("/pacientecuidador/update/:id", pcHandler.UpdatePacienteCuidador)
	app.Delete("/pacientecuidador/delete/:id", pcHandler.DeletePacienteCuidador)
	app.Get("/pacientecuidador/getAll", pcHandler.GetAllPacienteCuidador)
	app.Get("/pacientecuidador/cuidadores/:pacienteId", pcHandler.GetCuidadoresByPaciente)
	app.Get("/pacientecuidador/pacientes/:cuidadorId", pcHandler.GetPacientesByCuidador)
}

// SetupAgendaRoutes configura las rutas para la entidad Agenda.
func SetupAgendaRoutes(app *fiber.App, agendaHandler handlers.AgendaHandler) {
	app.Post("/agenda/insert", agendaHandler.CreateAgenda)
	app.Put("/agenda/update/:id", agendaHandler.UpdateAgenda)
	app.Delete("/agenda/delete/:id", agendaHandler.DeleteAgenda)
	app.Get("/agenda/getAll", agendaHandler.GetAllAgendas)
	app.Get("/agenda/:id", agendaHandler.GetAgendaById)
}

// SetupMedicamentoRoutes configura las rutas para la entidad Medicamentos.
func SetupMedicamentoRoutes(app *fiber.App, MedicineHandler handlers.MedicineHandler) {
	app.Post("/medicamento/insert", MedicineHandler.CreateMedicine)
	app.Put("/medicamento/update/:id", MedicineHandler.UpdateMedicine)
	app.Delete("/medicamento/delete/:id", MedicineHandler.DeleteMedicine)
	app.Get("/medicamento/getAll", MedicineHandler.GetAllMedicines)
	app.Get("/medicamento/:id", MedicineHandler.GetMedicineById)
}

// SetupHorarioMedicamentosRoutes configura las rutas para la entidad HorarioMedicamentos.
func SetupHorarioMedicamentosRoutes(app *fiber.App, horarioMedicamentosHandler handlers.HorarioMedicamentosHandler) {
	app.Post("/horariomedicamentos/insert", horarioMedicamentosHandler.Insert2)
	app.Get("/horariomedicamentos/getAll", horarioMedicamentosHandler.GetAll)
}
