package services

import (
	"os"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(to, subject, content string) {
	message := mail.NewSingleEmail(
		mail.NewEmail(os.Getenv("SENDER_NAME"), os.Getenv("SENDER_EMAIL")),
		subject,
		mail.NewEmail("An IDK User", to),
		content,
		content,
	)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		config.Logger.Println(err)
	}
}
