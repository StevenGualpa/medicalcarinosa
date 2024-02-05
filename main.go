package main

import (
	"GolandProyectos/handlers"
	"GolandProyectos/models"
	"GolandProyectos/repository"
	"GolandProyectos/routers" // Asegúrate de ajustar esta importación según tu estructura de paquetes
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := fiber.New()
	config := viper.New()

	// Configuración de Viper...
	config.AutomaticEnv()
	config.SetDefault("APP_PORT", "3000")
	config.SetDefault("APP_ENV", "development")
	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath(".")
	config.AddConfigPath("/etc/secrets/")
	if err := config.ReadInConfig(); err != nil {
		log.Println("Error al leer el archivo de configuración:", err)
	}

	// Conexión a la base de datos...
	dsn := config.GetString("DATABASE_URL") // Asume que tienes esta variable configurada
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	// Automigración para el modelo User...
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error en la automigración: %v", err)
	}

	// Crear instancia del repositorio y handler
	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	// Configurar rutas de usuarios
	routers.SetupUserRoutes(app, userHandler)

	// Iniciar el servidor...
	port := config.GetString("APP_PORT")
	log.Printf("Servidor iniciado en el puerto %s", port)
	log.Fatal(app.Listen(":" + port))
}
