package main

import (
	"GolandProyectos/models"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := fiber.New()
	config := viper.New()

	// Lee las variables de entorno
	config.AutomaticEnv()

	config.SetDefault("APP_PORT", "3000")
	config.SetDefault("APP_ENV", "development")

	// Lee el archivo de configuración
	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath(".")
	config.AddConfigPath("/etc/secrets/")

	err := config.ReadInConfig()
	if err != nil {
		log.Println("Error al leer el archivo de configuración:", err)
	}

	// Conexión a la base de datos
	dsn := "host=ep-lingering-snowflake-a5j9m53w.us-east-2.aws.neon.tech user=stevengualpa password=VamLyM2btnd4 dbname=carinosabd port=5432 sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	// Automigración para el modelo User
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error en la automigración: %v", err)
	}

	// Ruta de bienvenida
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Hola, Mundo!")
	})

	// Ruta POST para crear usuarios
	app.Post("/api/users", func(c *fiber.Ctx) error {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "No se puede analizar el cuerpo de la solicitud",
			})
		}

		// Crea el usuario en la base de datos
		result := db.Create(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "No se puede crear el usuario",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	})

	// Iniciar el servidor
	port := config.GetString("APP_PORT")
	log.Printf("Servidor iniciado en el puerto %s", port)
	log.Fatal(app.Listen(":" + port))
}
