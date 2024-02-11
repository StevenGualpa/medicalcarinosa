package handlers

import (
	"GolandProyectos/models"
	"GolandProyectos/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type userHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) UserHandler {
	return &userHandler{repo: repo}
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	users, count, err := h.repo.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.JSON(fiber.Map{"users": users, "count": count})
}

func (h *userHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid User ID"})
	}
	user, err := h.repo.GetUserByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User Not Found"})
	}
	return c.JSON(user)
}

// En tu handler CreateUser
func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	var input struct {
		models.User
		Relacion         string `json:"relacion"`
		NumeroEmergencia string `json:"numeroEmergencia"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	var roleData interface{}
	switch input.Roles {
	case "cuidador":
		roleData = models.Cuidador{
			Relacion: input.Relacion,
			// Asegúrate de que los campos adicionales necesarios estén aquí
		}
	case "paciente":
		roleData = models.Paciente{
			NumeroEmergencia: input.NumeroEmergencia,
			// Asegúrate de que los campos adicionales necesarios estén aquí
		}
	}

	createdUser, err := h.repo.CreateUserWithRole(input.User, roleData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(201).JSON(createdUser)
}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
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

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid User ID"})
	}
	err = h.repo.DeleteUser(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(200).SendString("User successfully deleted")
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	var loginInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginInfo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	user, message, err := h.repo.Login(loginInfo.Email, loginInfo.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": message})
	}

	return c.JSON(fiber.Map{"message": message, "user": user})
}
