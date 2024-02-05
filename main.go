// main.go
package main

import (
	"GolandProyectos/handlers"
	"GolandProyectos/models"
	"GolandProyectos/repository"
	"GolandProyectos/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	config := viper.New()

	// Configuración de Viper y carga de variables de entorno
	config.AutomaticEnv()

	config.SetDefault("APP_PORT", "3000")
	config.SetDefault("APP_ENV", "development")

	// Lectura del archivo de configuración
	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath(".")
	config.AddConfigPath("/etc/secrets/")

	err := config.ReadInConfig()
	if err != nil {
		log.Println("Error reading config file, using default settings:", err)
	}

	// Conexión a la base de datos PostgreSQL
	dsn := config.GetString("DATABASE_URL") // Asume que DATABASE_URL es una variable de entorno
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Automigración para el modelo User
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error automigrating: %v", err)
	}

	// Instancia del repositorio y los manejadores
	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	// Configuración de rutas
	routers.SetupUserRoutes(app, userHandler)

	// Ruta de bienvenida
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Iniciar el servidor Fiber
	log.Fatal(app.Listen(":" + config.GetString("APP_PORT")))
}
