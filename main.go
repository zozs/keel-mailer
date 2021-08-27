package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type keelWebhookMessage struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt"`
}

func sendMail(subject string, body string) {
	hostname := os.Getenv("KWM_SMTP_HOST")
	port := os.Getenv("KWM_SMTP_PORT")
	from := os.Getenv("KWM_SENDER")
	to := os.Getenv("KWM_RECIPIENT")

	// Prepare the e-mail contents.
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subject

	// Message is headers + body
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the remote SMTP server with TLS and auth.
	c, err := smtp.Dial(hostname + ":" + port)
	if err != nil {
		log.Fatal(err)
	}

	tlsconfig := &tls.Config{
		ServerName: hostname,
	}
	if err := c.StartTLS(tlsconfig); err != nil {
		log.Fatal(err)
	}

	auth := smtp.PlainAuth("", os.Getenv("KWM_SMTP_USER"), os.Getenv("KWM_SMTP_PASS"), hostname)
	if err := c.Auth(auth); err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail(from); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt(to); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(wc, message)
	wc.Close()

	c.Quit()

	log.Printf("Successfully sent email to %s", to)
}

func webhook(w http.ResponseWriter, r *http.Request) {
	var msg keelWebhookMessage

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Print("Failed to parse JSON webhook body:", err)
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	log.Printf("Got webhook name: %s, with message: %s", msg.Name, msg.Message)

	sendMail("Keel update information", fmt.Sprintf("Hi,\n\n**%s**\n\n%s\n\nBest regards,\nkeel-webhook-mailer", msg.Name, msg.Message))
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	port := "6000"
	if val, ok := os.LookupEnv("KWM_PORT"); ok {
		port = val
	}

	http.HandleFunc("/webhook", webhook)

	listen := ":" + port
	log.Printf("Listening on %s", listen)
	http.ListenAndServe(listen, nil)
}
