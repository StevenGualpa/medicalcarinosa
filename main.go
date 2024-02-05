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

	// Load the config
	err := config.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	dsn := "host=ep-lingering-snowflake-a5j9m53w.us-east-2.aws.neon.tech user=stevengualpa password=VamLyM2btnd4 dbname=carinosabd port=5432 sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.User{})
	db.Create(&models.User{
		FirstName: "Steven",
		LastName:  "Gualpa",
	})

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

	app.Listen(":" + config.GetString("APP_PORT"))
}
