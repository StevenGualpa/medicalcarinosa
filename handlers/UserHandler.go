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
	Login(c *fiber.Ctx) error // Método de inicio de sesión añadido

}

// userHandler es la implementación concreta de UserHandler.
type userHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) UserHandler {
	return &userHandler{repo}
}

// Login maneja la solicitud POST para el inicio de sesión de un usuario.
func (h *userHandler) Login(c *fiber.Ctx) error {
	var loginInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	user, message, err := h.repo.Login(loginInfo.Email, loginInfo.Password)
	if err != nil {
		// Puedes decidir devolver el mismo mensaje de error para ambos casos
		// (usuario no encontrado y contraseña incorrecta) para evitar dar pistas a posibles atacantes
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": message})
	}

	// Devuelve los datos del usuario y un mensaje de éxito
	return c.JSON(fiber.Map{"message": message, "user": user})
}

// GetUsers maneja la solicitud GET para recuperar todos los usuarios y su cantidad.
func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	users, count, err := h.repo.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.JSON(fiber.Map{"users": users, "count": count})
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
