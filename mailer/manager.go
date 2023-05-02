package mailer

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/mailgun/mailgun-go/v4"
)

var publicApiKey = "pubkey-8ab2abe7c32c6ea5a7b3b786975b201b"

func SendMail(to string, from string, message string) {
	fmt.Printf("Email sent to: %s\n", to)

	// Sender data.
	username := "postmaster@sandbox702a40ad803a41ac82532e4fcb1f4a13.mailgun.org"
	password := "70dfc650927b2d65673be3beccea961a-102c75d8-9b83b5a0"

	// Receiver email address.
	toList := []string{
		to,
	}

	// smtp server configuration.
	smtpHost := "smtp.mailgun.org:587"

	// Authentication.
	auth := smtp.PlainAuth("", username, password, "smtp.mailgun.org")

	// Sending email.
	err := smtp.SendMail(smtpHost, auth, from, toList, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		"Mailgun Sandbox <postmaster@sandbox702a40ad803a41ac82532e4fcb1f4a13.mailgun.org>",
		"Hello Vimvek",
		"Testing some Mailgun awesomeness!",
		"vimvekjomshi@gmail.com",
	)
	_, id, err := mg.Send(context.Background(), m)
	return id, err
}
