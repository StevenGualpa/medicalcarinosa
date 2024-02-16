package handlers

import (
	"GolandProyectos/models"
	"GolandProyectos/repository"
	"context"
	"github.com/arduino/iot-client-go"
	"github.com/gofiber/fiber/v2"
	cc "golang.org/x/oauth2/clientcredentials"
	"log"
	"net/url"
	"strconv"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetUsersRoles(c *fiber.Ctx) error
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

func (h *userHandler) GetUsersRoles(c *fiber.Ctx) error {
	var filter struct {
		Role string `json:"role"`
	}

	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	users, count, err := h.repo.GetAllUsersWithRoleFilter(filter.Role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error roles"})
	}

	// Limpia la información redundante antes de enviar la respuesta
	for i := range users {
		if users[i].Paciente != nil && users[i].Paciente.User.ID != 0 {
			users[i].Paciente.User = models.User{}
		}
		if users[i].Cuidador != nil && users[i].Cuidador.User.ID != 0 {
			users[i].Cuidador.User = models.User{}
		}
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
// userHandler.go
func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	var input struct {
		models.User
		Relacion         string `json:"relacion"`
		NumeroEmergencia string `json:"numeroEmergencia"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	// Construir roleData basado en el rol
	var roleData interface{}
	switch input.Roles {
	case "cuidador":
		roleData = models.Cuidador{
			Relacion: input.Relacion,
		}
	case "paciente":
		roleData = models.Paciente{
			NumeroEmergencia: input.NumeroEmergencia,
		}
	case "admin":
		// No se requiere acción adicional para el rol 'admin'
		break // No hay datos adicionales que asignar
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role"})
	}

	// Crear usuario con el rol correspondiente
	createdUser, err := h.repo.CreateUserWithRole(input.User, roleData)
	if err != nil {
		// El error ya no será siempre "Internal Server Error"
		// Puede ser más específico dependiendo de la validación en CreateUserWithRole
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
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

// Definir las credenciales del cliente de Arduino IoT como constantes o variables
const (
	clientID     = "k9asQ6bG8GoiJyMcdmiPSaAWCvntvIVe"
	clientSecret = "BZcIzieDfhb8mKF3auezYLuFi6cRVsUhFCiEzAMDJEvA2jaDzGM38YfEo95BT3X1"
)

// ArduinoHandler struct para implementar el manejador
type ArduinoHandler struct{}

// NewArduinoHandler crea una nueva instancia de ArduinoHandler
func NewArduinoHandler() *ArduinoHandler {
	return &ArduinoHandler{}
}

// GetArduinoDevices maneja la solicitud GET para obtener dispositivos de Arduino IoT
func (h *ArduinoHandler) GetArduinoDevices(c *fiber.Ctx) error {
	// Configurar OAuth2
	additionalValues := url.Values{}
	additionalValues.Add("audience", "https://api2.arduino.cc/iot")
	config := cc.Config{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		TokenURL:       "https://api2.arduino.cc/iot/v1/clients/token",
		EndpointParams: additionalValues,
	}

	// Obtener el token de acceso
	tok, err := config.Token(context.Background())
	if err != nil {
		log.Printf("Error retrieving access token, %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve access token"})
	}

	// Crear contexto con el token de acceso
	ctx := context.WithValue(context.Background(), iot.ContextAccessToken, tok.AccessToken)

	// Crear instancia del cliente de la API de Arduino IoT
	client := iot.NewAPIClient(iot.NewConfiguration())

	// Obtener la lista de dispositivos
	devices, _, err := client.DevicesV2Api.DevicesV2List(ctx, nil)
	if err != nil {
		log.Printf("Error getting devices, %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get devices"})
	}

	// Devolver la lista de dispositivos como respuesta
	if len(devices) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "No device found"})
	} else {
		return c.Status(fiber.StatusOK).JSON(devices)
	}
}
