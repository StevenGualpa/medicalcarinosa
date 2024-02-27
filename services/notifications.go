package services

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func SendNotification(app *firebase.App, title, body string) error {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	// Aquí deberías obtener los tokens de los dispositivos de tu base de datos
	// Para este ejemplo, usaremos un token ficticio
	tokens := []string{"your_device_token_here"}

	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	response, err := client.SendMulticast(ctx, message)
	if err != nil {
		return err
	}

	// Procesa la respuesta, si es necesario
	_ = response

	return nil
}
