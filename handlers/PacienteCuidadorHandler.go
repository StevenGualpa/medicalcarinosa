package handlers

import (
	"GolandProyectos/models"
	"GolandProyectos/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type PacienteCuidadorHandler interface {
	CreatePacienteCuidador(c *fiber.Ctx) error
	UpdatePacienteCuidador(c *fiber.Ctx) error
	DeletePacienteCuidador(c *fiber.Ctx) error
	GetAllPacienteCuidador(c *fiber.Ctx) error
	GetCuidadoresByPaciente(c *fiber.Ctx) error
	GetPacientesByCuidador(c *fiber.Ctx) error
}

type pacienteCuidadorHandler struct {
	repo repository.PacienteCuidadorRepository
}

func NewPacienteCuidadorHandler(repo repository.PacienteCuidadorRepository) PacienteCuidadorHandler {
	return &pacienteCuidadorHandler{repo: repo}
}

func (h *pacienteCuidadorHandler) CreatePacienteCuidador(c *fiber.Ctx) error {
	pc := new(models.PacienteCuidador)
	if err := c.BodyParser(pc); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	createdPc, err := h.repo.Create(*pc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdPc)
}

func (h *pacienteCuidadorHandler) UpdatePacienteCuidador(c *fiber.Ctx) error {
	pc := new(models.PacienteCuidador)
	if err := c.BodyParser(pc); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	updatedPc, err := h.repo.Update(*pc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedPc)
}

func (h *pacienteCuidadorHandler) DeletePacienteCuidador(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = h.repo.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *pacienteCuidadorHandler) GetAllPacienteCuidador(c *fiber.Ctx) error {
	pcs, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(pcs)
}

func (h *pacienteCuidadorHandler) GetCuidadoresByPaciente(c *fiber.Ctx) error {
	pacienteID, err := strconv.Atoi(c.Params("pacienteId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid paciente ID"})
	}

	cuidadores, err := h.repo.GetByPaciente(uint(pacienteID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(cuidadores)
}

func (h *pacienteCuidadorHandler) GetPacientesByCuidador(c *fiber.Ctx) error {
	cuidadorID, err := strconv.Atoi(c.Params("cuidadorId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid cuidador ID"})
	}

	pacientes, err := h.repo.GetByCuidador(uint(cuidadorID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(pacientes)
}
