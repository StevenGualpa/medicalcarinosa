// HorarioMedicamentosHandler.go
package handlers

import (
	"GolandProyectos/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type HorarioMedicamentosHandler interface {
	Insert(c *fiber.Ctx) error
	Insert2(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Delete2(c *fiber.Ctx) error
}

type horarioMedicamentosHandler struct {
	repo repository.HorarioMedicamentosRepository
}

func NewHorarioMedicamentosHandler(repo repository.HorarioMedicamentosRepository) HorarioMedicamentosHandler {
	return &horarioMedicamentosHandler{repo: repo}
}

func (h *horarioMedicamentosHandler) Insert(c *fiber.Ctx) error {
	var request struct {
		PacienteID    uint `json:"pacienteID"`
		MedicamentoID uint `json:"medicamentoID"`
		DosisInicial  int  `json:"dosisInicial"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request"})
	}

	err := h.repo.Insert(request.PacienteID, request.MedicamentoID, request.DosisInicial)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error inserting schedule"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Schedule created successfully"})
}

func (h *horarioMedicamentosHandler) GetAll(c *fiber.Ctx) error {
	horariosMedicamentos, err := h.repo.GetAll2()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error retrieving schedules"})
	}

	return c.JSON(horariosMedicamentos)
}

func (h *horarioMedicamentosHandler) Insert2(c *fiber.Ctx) error {
	var request struct {
		PacienteID    uint `json:"pacienteID"`
		MedicamentoID uint `json:"medicamentoID"`
		DosisInicial  int  `json:"dosisInicial"`
		Frecuencia    int  `json:"frecuencia"` // Nuevo campo para capturar la frecuencia
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request"})
	}

	// Asegúrate de pasar la frecuencia al método Insert del repositorio.
	err := h.repo.InsertFinal(request.PacienteID, request.MedicamentoID, request.DosisInicial, request.Frecuencia)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error inserting schedule"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Schedule created successfully"})
}

func (h *horarioMedicamentosHandler) Delete(c *fiber.Ctx) error {
	// Obtener el ID desde los parámetros de la URL
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// Si el ID no es válido, devuelve un error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	// Intentar eliminar el horario de medicamento con el ID especificado
	err = h.repo.Delete(uint(id))
	if err != nil {
		// Si ocurre un error durante la eliminación, devuelve un error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al eliminar el horario de medicamento"})
	}

	// Si la eliminación es exitosa, devuelve un mensaje de éxito
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Horario de medicamento eliminado con éxito"})
}

func (h *horarioMedicamentosHandler) Delete2(c *fiber.Ctx) error {
	// Ejemplo de extracción de parámetros de la URL
	pacienteID, err := strconv.Atoi(c.Params("pacienteID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID de paciente inválido"})
	}

	medicamentoID, err := strconv.Atoi(c.Params("medicamentoID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID de medicamento inválido"})
	}

	err = h.repo.Delete2(uint(pacienteID), uint(medicamentoID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al eliminar el horario de medicamento"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Horario de medicamento eliminado con éxito"})
}
