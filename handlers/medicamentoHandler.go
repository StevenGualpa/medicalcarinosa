// medicamentoHandler.go
package handlers

import (
	"GolandProyectos/models"
	"GolandProyectos/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type MedicineHandler interface {
	CreateMedicine(c *fiber.Ctx) error
	UpdateMedicine(c *fiber.Ctx) error
	DeleteMedicine(c *fiber.Ctx) error
	GetAllMedicines(c *fiber.Ctx) error
	GetMedicineById(c *fiber.Ctx) error
}

type medicineHandler struct {
	repo repository.MedicineRepository
}

func NewMedicineHandler(repo repository.MedicineRepository) MedicineHandler {
	return &medicineHandler{repo: repo}
}

func (h *medicineHandler) CreateMedicine(c *fiber.Ctx) error {
	var medicine models.Medicine
	if err := c.BodyParser(&medicine); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdMedicine, err := h.repo.Create(medicine)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdMedicine)
}

func (h *medicineHandler) UpdateMedicine(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var medicine models.Medicine
	if err := c.BodyParser(&medicine); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	medicine.ID = uint(id)

	updatedMedicine, err := h.repo.Update(medicine)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedMedicine)
}

func (h *medicineHandler) DeleteMedicine(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *medicineHandler) GetAllMedicines(c *fiber.Ctx) error {
	medicines, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(medicines)
}

func (h *medicineHandler) GetMedicineById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	medicine, err := h.repo.GetById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(medicine)
}
