package services

import (
	"fmt"
	"go-fiber/bootstrap"
	"log"
	"strconv"

	"gopkg.in/mail.v2"
)

var app = bootstrap.App()
var globalEnv = app.Env

func SendEmail(body string, to string) error {
	// Set up the message
	message := mail.NewMessage()

	// Set the sender, recipient, subject, and body
	message.SetHeader("From", globalEnv.SMTP.Username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Leave Request")
	message.SetBody("text/html", body)

	port, _ := strconv.Atoi(globalEnv.SMTP.Port)
	// Set up the SMTP server details
	dialer := mail.NewDialer(globalEnv.SMTP.Host, port, globalEnv.SMTP.Username, globalEnv.SMTP.Password)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Errorf("Could not send email: %v", err)
	} else {
		log.Println("Email sent successfully!")
	}
	return nil
}
