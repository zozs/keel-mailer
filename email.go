package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type Email struct {
	Subject string
	Body    string
}

func emailListener(c chan Email) {
	for {
		select {
		case email := <-c:
			if err := sendMail(email); err != nil {
				log.Printf("Failed to send email, got error: %s", err)
			}
		}
	}
}

func sendMail(email Email) error {
	hostname := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	from := os.Getenv("SMTP_SENDER")
	to := os.Getenv("SMTP_RECIPIENT")

	// Prepare the e-mail contents.
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = email.Subject

	// Message is headers + body
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + email.Body

	// Connect to the remote SMTP server with TLS and auth.
	c, err := smtp.Dial(hostname + ":" + port)
	if err != nil {
		return err
	}

	tlsconfig := &tls.Config{
		ServerName: hostname,
	}
	if err := c.StartTLS(tlsconfig); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), hostname)
	if err := c.Auth(auth); err != nil {
		return err
	}

	// Set the sender and recipient first
	if err := c.Mail(from); err != nil {
		return err
	}
	if err := c.Rcpt(to); err != nil {
		return err
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		return err
	}
	fmt.Fprintf(wc, message)
	wc.Close()

	c.Quit()

	log.Printf("Successfully sent email to %s", to)
	return nil
}
