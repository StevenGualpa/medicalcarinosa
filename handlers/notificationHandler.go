package handlers

import (
	"GolandProyectos/repository"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
	"log"
)

type NotificationHandler struct {
	DeviceTokenRepo repository.DeviceTokenRepository
}

func NewNotificationHandler(repo repository.DeviceTokenRepository) *NotificationHandler {
	return &NotificationHandler{DeviceTokenRepo: repo}
}

func (h *NotificationHandler) SendNotifications(c *fiber.Ctx) error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("credenciales/notificauteq-19631-firebase-adminsdk-ztqd5-1104d79c98.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	tokens, err := h.DeviceTokenRepo.GetAllDeviceTokens()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al recuperar tokens de dispositivo"})
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	for _, token := range tokens {
		message := &messaging.Message{
			Notification: &messaging.Notification{
				Title: "Título de Notificación",
				Body:  "Cuerpo de Notificación",
			},
			Token: token.Token,
		}

		response, err := client.Send(ctx, message)
		if err != nil {
			log.Printf("error sending message: %v\n", err)
			continue
		}

		fmt.Printf("Successfully sent message: %s\n", response)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Notificaciones enviadas correctamente"})
}
