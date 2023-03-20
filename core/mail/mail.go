package mail

import (
	"crypto/tls"
	"errors"
	"os"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

type SMTPServer struct {
	Host     string
	Port     int
	Username string
	Password string
}

// SendMail sends an email message to one or more recipients using the
// SMTP server specified by the environment variables MAIL_HOST, MAIL_PORT,
// MAIL_USERNAME, MAIL_PASSWORD, and MAIL_ENCRYPTION. The email message
// is addressed to the recipients specified in the "to" parameter, and has
// the subject and body specified in the "subject" and "body" parameters,
// respectively.
func SendMail(to []string, subject, body string, bodyType string) error {
	// Load the SMTP server configuration from environment variables
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		return err
	}
	smtpUsername := os.Getenv("MAIL_USERNAME")
	smtpPassword := os.Getenv("MAIL_PASSWORD")
	smtpEncryption := os.Getenv("MAIL_ENCRYPTION")

	// Create a new dialer and set TLS config based on encryption
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	if smtpEncryption == "tls" {
		d.TLSConfig = &tls.Config{}
	} else if smtpEncryption == "ssl" {
		d.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         smtpHost,
		}
	}

	// Set up the email message
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_FROM_ADDRESS"))
	m.SetHeader("To", strings.Join(to, ","))
	m.SetHeader("Subject", subject)

	switch bodyType {
	case "html":
		m.SetBody("text/html", body)
	case "text":
		m.SetBody("text/plain", body)
	default:
		return errors.New("invalid bodyType; must be 'html' or 'text'")
	}

	// Send the email
	err = d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
