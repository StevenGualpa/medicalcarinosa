// medicamentoHandler.go
package handlers

import (
	"GolandProyectos/models"
	"GolandProyectos/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type MedicamentoHandler interface {
	CreateMedicamento(c *fiber.Ctx) error
	UpdateMedicamento(c *fiber.Ctx) error
	DeleteMedicamento(c *fiber.Ctx) error
	GetAllMedicamentos(c *fiber.Ctx) error
	GetMedicamentoById(c *fiber.Ctx) error
}

type medicamentoHandler struct {
	repo repository.MedicamentoRepository
}

func NewMedicamentoHandler(repo repository.MedicamentoRepository) MedicamentoHandler {
	return &medicamentoHandler{repo: repo}
}

func (h *medicamentoHandler) CreateMedicamento(c *fiber.Ctx) error {
	var medicamento models.Medicamento
	if err := c.BodyParser(&medicamento); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdMedicamento, err := h.repo.Create(medicamento)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdMedicamento)
}

func (h *medicamentoHandler) UpdateMedicamento(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var medicamento models.Medicamento
	if err := c.BodyParser(&medicamento); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	medicamento.ID = uint(id)

	updatedMedicamento, err := h.repo.Update(medicamento)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedMedicamento)
}

func (h *medicamentoHandler) DeleteMedicamento(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *medicamentoHandler) GetAllMedicamentos(c *fiber.Ctx) error {
	medicamentos, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(medicamentos)
}

func (h *medicamentoHandler) GetMedicamentoById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	medicamento, err := h.repo.GetById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(medicamento)
}
