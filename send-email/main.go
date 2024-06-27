package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	// Load credentials from environment variables
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	recipient := os.Getenv("RECIPIENT_EMAIL")

	// Check if the necessary environment variables are set
	if smtpUser == "" || smtpPass == "" || recipient == "" {
		log.Fatalln("SMTP_USER, SMTP_PASS, and RECIPIENT_EMAIL environment variables must be set")
	}

	// Create a new message
	m := gomail.NewMessage()

	// Set email sender
	m.SetHeader("From", smtpUser)

	// Set email recipient
	m.SetHeader("To", recipient)

	// Set email subject
	m.SetHeader("Subject", "Payroll Notification")

	// Set email body
	m.SetBody("text/plain", "This is a notification about your payroll.")

	// Set up the SMTP dialer
	d := gomail.NewDialer("smtp.gmail.com", 587, smtpUser, smtpPass)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send the email:", err)
		return
	}

	log.Println("Email sent successfully!")
}
