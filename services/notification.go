package services

import (
	"os"

	"github.com/NaySoftware/go-fcm"
	"github.com/vaziria/pdcnotification/models"
)

var serverKey = os.Getenv("NOTIF_SERVERKEY")

func SendNotification(user models.User, message string, image string) error {
	ids := user.Tokens

	data := map[string]string{
		"message": message,
		"image":   image,
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(ids, data)

	_, err := c.Send()

	return err
}
