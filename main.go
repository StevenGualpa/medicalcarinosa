package main

import (
	"GolandProyectos/models"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func setupRoutes(app *fiber.App, db *gorm.DB) {
	// Grupo de rutas para API
	api := app.Group("/api")

	// Rutas relacionadas con usuarios
	users := api.Group("/users")
	users.Post("/", func(c *fiber.Ctx) error {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse body",
			})
		}

		if result := db.Create(&user); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Cannot create user",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	})

	// Aquí puedes agregar más rutas a tu grupo 'users' o crear nuevos grupos.
}

func main() {
	app := fiber.New()
	config := viper.New()

	// Configuración y conexión a la base de datos omitidas para brevedad...

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// Configurar rutas
	setupRoutes(app, db)

	// Ruta de bienvenida
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Iniciar el servidor
	port := config.GetString("APP_PORT")
	log.Fatal(app.Listen(":" + port))
}
