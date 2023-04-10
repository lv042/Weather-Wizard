package main

import (
	"bufio"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
	"strings"
	"time"
)

var email string
var enabled bool
var key string
var stopDailyMailer chan struct{}

func setupEmail(e string, en bool) {
	email = e
	enabled = en
	readKey()
	stopDailyMailer = make(chan struct{})
	if enabled {
		go func() {
			sendSetupMail(email)
			setupDailyMailer(stopDailyMailer)
		}()
	} else {
		disableDailyMailer()
	}
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

func sendSetupMail(email string) {
	mail.NewEmail("Luca", "luca.v.kannen@gmail.com")
	subject := "Backend Notifications"
	mail.NewEmail("Luca", email)
	plainTextContent := "You are now subscribed to notifications. The backend will send from now on notifications to " +
		"this email address. You will receive a daily summary of all requests and errors. You can unsubscribe in the Admin Panel."
	htmlContent := "<p>Hi there, you are now subscribed to notifications.</p>\n\n<p>The backend will send from now on notifications to this email address. You will receive a daily summary of all requests and errors.</p>\n\n<p>You can unsubscribe in the Admin Panel."
	sendSummaryMail(email, subject, plainTextContent, htmlContent)
}

func sendSummaryMail(email, subject, plainTextContent, htmlContent string) {
	from := mail.NewEmail("Luca", "luca.v.kannen@gmail.com")
	to := mail.NewEmail("Luca", email)
	emailMessage := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(key)
	response, err := client.Send(emailMessage)
	if err != nil {
		fiberApp.Log(fmt.Sprintf("Error: %s", err))
	} else {
		fiberApp.Log(fmt.Sprintf("Sent email to %s", email))
		fiberApp.Log(fmt.Sprintf("Status code: %d", response.StatusCode)) // Update the format specifier here
		fiberApp.Log(fmt.Sprintf("Response: %+v", response))              // Update the format specifier here
	}
}

func setupDailyMailer(stop chan struct{}) {
	ticker := time.NewTicker(24 * time.Hour)
	for {
		select {
		case <-ticker.C:
			sendDailySummary()
		case <-stop:
			ticker.Stop()
			return
		}
	}
}

func disableDailyMailer() {
	if stopDailyMailer != nil {
		close(stopDailyMailer)
	}
}

func sendDailySummary() {
	metricsData := fiberApp.metrics.GetMetrics()

	var requestCounts []string
	var errorCounts []string

	for k, v := range metricsData["requestCount"].(map[string]int) {
		requestCounts = append(requestCounts, fmt.Sprintf("%s: %d", k, v))
	}

	for k, v := range metricsData["errorCount"].(map[string]int) {
		errorCounts = append(errorCounts, fmt.Sprintf("%s: %d", k, v))
	}

	requestSummary := strings.Join(requestCounts, "\n")
	errorSummary := strings.Join(errorCounts, "\n")

	subject := "Daily Summary"
	plainTextContent := fmt.Sprintf("Daily request summary:\n%s\n\nDaily error summary:\n%s", requestSummary, errorSummary)
	htmlContent := fmt.Sprintf("<p><strong>THIDaily request summary:</strong></p><pre>%s</pre><p><strong>Daily error summary:</strong></p><pre>%s</pre>", requestSummary, errorSummary)
	sendSummaryMail(email, subject, plainTextContent, htmlContent)
}
