package email

import (
	"crypto/tls"
	"fmj/config"
	"fmt"
	"log"
	"regexp"
	"time"

	"gopkg.in/mail.v2"
)

type Service interface {
	SendVerificationEmail(to string, code uint32) error
	SendWelcomeEmail(to, name string) error
	SendPasswordResetEmail(email string, token uint32) error
}

type service struct {
	config *config.Config
}

// SendPasswordResetEmail implements Service.
func (s *service) SendPasswordResetEmail(email string, token uint32) error {
	subject := "Reset Your Password"
	body := fmt.Sprintf("Hello,\n\nYou requested to reset your password. Copy the code to set a new password: %d\n\nThis code will expire in 24 hours.\n\nIf you did not request this, please ignore this email.", token)

	return s.sendEmail(email, subject, body)
}

func (s *service) SendVerificationEmail(to string, code uint32) error {
	subject := "Verify your email address"
	body := fmt.Sprintf("Hello creator,\n\nPlease verify your email by copying this code: %d", code)

	return s.sendEmail(to, subject, body)
}

func (s *service) SendWelcomeEmail(to, name string) error {
	subject := "Welcome to our platform!"
	body := fmt.Sprintf("Hello %s,\n\nWelcome to our platform. We're excited to have you!", name)

	return s.sendEmail(to, subject, body)
}

func (s *service) sendEmail(to, subject, body string) error {
	// Validate email address
	if to == "" {
		return fmt.Errorf("recipient email address cannot be empty")
	}

	// Optional: Add more thorough email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(to) {
		return fmt.Errorf("invalid email address format: %s", to)
	}

	// Validate other required fields
	if subject == "" {
		return fmt.Errorf("email subject cannot be empty")
	}
	if body == "" {
		return fmt.Errorf("email body cannot be empty")
	}
	if s.config.FromEmail == "" {
		return fmt.Errorf("sender email address cannot be empty")
	}

	m := mail.NewMessage()
	m.SetHeader("From", s.config.FromEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Debug logging with more detail
	log.Printf("Preparing to send email: from=%s, to=%s, subject=%s",
		s.config.FromEmail,
		to,
		subject,
	)

	d := mail.NewDialer(s.config.SMTPHost, s.config.SMTPPort, s.config.SMTPUsername, s.config.SMTPPassword)
	d.SSL = false
	d.TLSConfig = &tls.Config{
		ServerName:         s.config.SMTPHost,
		InsecureSkipVerify: false,
	}
	d.StartTLSPolicy = mail.MandatoryStartTLS
	d.Timeout = 10 * time.Second

	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err := d.DialAndSend(m)
		if err == nil {
			log.Printf("Email sent successfully to %s", to)
			return nil
		}

		if i < maxRetries-1 {
			log.Printf("Attempt %d failed: %v. Retrying...", i+1, err)
			time.Sleep(time.Second * 2)
			continue
		}

		return fmt.Errorf("failed to send email after %d attempts: %w", maxRetries, err)
	}

	return nil
}

func NewService(config *config.Config) Service {
	return &service{config: config}
}
