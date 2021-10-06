package services

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/smtp"
	"os"
	"text/template"

	"github.com/anti-lgbt/learning-be/config"
)

type Email struct {
	FromAddress string
	FromName    string
	ToAddress   string
	Subject     string
	Reader      io.Reader
}

func SendEmail(to, subject, content string) {
	smtp_host := os.Getenv("SMTP_HOST")
	smtp_port := os.Getenv("SMTP_PORT")
	smtp_user := os.Getenv("SMTP_USER")
	smtp_password := os.Getenv("SMTP_PASSWORD")
	sender_name := os.Getenv("SENDER_NAME")
	sender_email := os.Getenv("SENDER_EMAIL")

	email := Email{
		FromAddress: sender_email,
		FromName:    sender_name,
		ToAddress:   to,
		Subject:     subject,
		Reader:      bytes.NewReader([]byte(content)),
	}

	tpl, err := template.ParseFiles("config/email.tpl")
	if err != nil {
		config.Logger.Errorf("Error: %v", err)
		return
	}

	buff := bytes.Buffer{}
	if err := tpl.Execute(&buff, email); err != nil {
		config.Logger.Errorf("Error: %v", err)
		return
	}

	text, err := ioutil.ReadAll(email.Reader)
	if err != nil {
		config.Logger.Errorf("Error: %v", err)
		return
	}

	msg := append(buff.Bytes(), "\r\n"...)
	msg = append(msg, text...)

	recipients := []string{email.ToAddress}
	auth := smtp.PlainAuth("", smtp_user, smtp_password, smtp_host)
	if err := smtp.SendMail(smtp_host+":"+smtp_port, auth, email.FromAddress, recipients, msg); err != nil {
		config.Logger.Errorf("Error: %v", err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
