package handlers

import (
	"GolandProyectos/models"
	"GolandProyectos/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type AgendaHandler interface {
	CreateAgenda(c *fiber.Ctx) error
	UpdateAgenda(c *fiber.Ctx) error
	DeleteAgenda(c *fiber.Ctx) error
	GetAllAgendas(c *fiber.Ctx) error
	GetAgendaById(c *fiber.Ctx) error
}

type agendaHandler struct {
	repo repository.AgendaRepository
}

func NewAgendaHandler(repo repository.AgendaRepository) AgendaHandler {
	return &agendaHandler{repo}
}

func (h *agendaHandler) CreateAgenda(c *fiber.Ctx) error {
	var agenda models.Agenda
	if err := c.BodyParser(&agenda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdAgenda, err := h.repo.Create(agenda)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdAgenda)
}

func (h *agendaHandler) UpdateAgenda(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var agenda models.Agenda
	if err := c.BodyParser(&agenda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	agenda.ID = uint(id)

	updatedAgenda, err := h.repo.Update(agenda)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedAgenda)
}

func (h *agendaHandler) DeleteAgenda(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *agendaHandler) GetAllAgendas(c *fiber.Ctx) error {
	agendas, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(agendas)
}

func (h *agendaHandler) GetAgendaById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	agenda, err := h.repo.GetById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(agenda)
}
