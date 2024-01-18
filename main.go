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

	// Read environment variables
	config.AutomaticEnv()

	config.SetDefault("APP_PORT", "3000")
	config.SetDefault("APP_ENV", "development")

	// Read the config file
	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath(".")
	config.AddConfigPath("/etc/secrets/")
	// config.AddConfigPath("/workspaces/api-ortografia")

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

	app.Listen(":" + config.GetString("APP_PORT"))
}
