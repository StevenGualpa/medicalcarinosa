// handlers/user_handler.go
package handlers

import (
	"GolandProyectos/models"
	"GolandProyectos/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// UserHandler define la interfaz para el manejador de usuarios.
type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

// userHandler es la implementación concreta de UserHandler.
type userHandler struct {
	repo repository.UserRepository
}

// NewUserHandler crea una nueva instancia de userHandler.
func NewUserHandler(repo repository.UserRepository) UserHandler {
	return &userHandler{repo}
}

// GetUsers maneja la solicitud GET para recuperar todos los usuarios.
func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.repo.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(200).JSON(users)
}

// GetUser maneja la solicitud GET para recuperar un usuario por su ID.
func (h *userHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid User ID"})
	}
	user, err := h.repo.GetUserByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User Not Found"})
	}
	return c.Status(200).JSON(user)
}

// CreateUser maneja la solicitud POST para crear un nuevo usuario.
func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}
	createdUser, err := h.repo.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(201).JSON(createdUser)
}

// UpdateUser maneja la solicitud PUT para actualizar un usuario existente.
func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid User ID"})
	}
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}
	user.ID = uint(id)
	updatedUser, err := h.repo.UpdateUser(*user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(200).JSON(updatedUser)
}

// DeleteUser maneja la solicitud DELETE para eliminar un usuario.
func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid User ID"})
	}
	err = h.repo.DeleteUser(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(200).SendString("User successfully deleted")
}
