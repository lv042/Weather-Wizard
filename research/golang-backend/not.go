package main

import (
	"bufio"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

var email string
var enabled bool
var key string

func setupEmail(e string, en bool) {
	email = e
	enabled = en
	readKey()
	sendMail(email)
}

func readKey() {
	file, err := os.Open("sendgrid.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key = scanner.Text()
	}
}

func sendMail(email string) {
	from := mail.NewEmail("Luca", "luca.v.kannen@gmail.com")
	subject := "Backend Notifications"
	to := mail.NewEmail("Luca", email)
	plainTextContent := "You are now subscribed to notifications. The backend will send from now on notifications to " +
		"this email address. You will receive a daily summary of all requests and errors. You can unsubscribe in the Admin Panel."
	htmlContent := "<p>Hi there, you are now subscribed to notifications.</p>\n\n<p>The backend will send from now on notifications to this email address. You will receive a daily summary of all requests and errors.</p>\n\n<p>You can unsubscribe in the Admin Panel."
	emailMessage := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(key)
	response, err := client.Send(emailMessage)
	if err != nil {
		fiberApp.Log(fmt.Sprintf("Error: %s", err))
	} else {
		fiberApp.Log(fmt.Sprintf("Sent email to %s", email))
		fiberApp.Log(fmt.Sprintf("Status code: %s", response.StatusCode))
		fiberApp.Log(fmt.Sprintf("Response: %s", response))

	}
}
