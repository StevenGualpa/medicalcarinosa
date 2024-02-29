// HorarioMedicamentosHandler.go
package handlers

import (
	"GolandProyectos/repository"
	"github.com/gofiber/fiber/v2"
)

type HorarioMedicamentosHandler interface {
	Insert(c *fiber.Ctx) error
	Insert2(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
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
	horariosMedicamentos, err := h.repo.GetAll()
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
