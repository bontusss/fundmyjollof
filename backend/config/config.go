package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI           string
	Port               int
	DatabaseName       string
	SessionSecret      string
	SMTPHost           string
	SMTPPort           int
	SMTPUsername       string
	SMTPPassword       string
	FromEmail          string
	BaseURL            string
	IsProd             bool
	GoogleClientID     string
	GoogleClientSecret string
	GoogleCallbackURL  string
	Environment        string
}

// NewConfig todo: Create .env file for these
func NewConfig() *Config {
	// Load .env first, before reading any environment variables
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: .env file not found in development mode")
		}
	}

	// Now read the environment variables after .env is loaded
	cfg := &Config{
		Environment:        os.Getenv("GO_ENV"),
		MongoURI:           os.Getenv("MONGO_URI"),
		Port:               func() int { port, _ := strconv.Atoi(os.Getenv("BACKEND_PORT")); return port}(),
		DatabaseName:       os.Getenv("DATABASE_NAME"),
		SMTPPort:           func() int { port, _ := strconv.Atoi(os.Getenv("SMTP_PORT")); return port }(),
		SessionSecret:      os.Getenv("SESSION_SECRET"),
		SMTPHost:           os.Getenv("SMTP_HOST"),
		SMTPUsername:       os.Getenv("SMTP_USERNAME"),
		SMTPPassword:       os.Getenv("SMTP_PASSWORD"),
		FromEmail:          os.Getenv("FROM_EMAIL"),
		BaseURL:            os.Getenv("BASE_URL"),
		IsProd:             false,
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		GoogleCallbackURL:  os.Getenv("GOOGLE_CALLBACK_URL"),
	}

	return cfg
}

func (c *Config) Validate() error {
	required := map[string]string{
		"MongoDB URI":          c.MongoURI,
		"Database Name":        c.DatabaseName,
		"Session Secret":       c.SessionSecret,
		"Google Client ID":     c.GoogleClientID,
		"Google Client Secret": c.GoogleClientSecret,
		"Google Callback URL":  c.GoogleCallbackURL,
		"SMTP Host":            c.SMTPHost,
		"SMTP Username":        c.SMTPUsername,
		"SMTP Password":        c.SMTPPassword,
		"From Email":           c.FromEmail,
		"Base URL":             c.BaseURL,
	}

	var missingVars []string
	for name, value := range required {
		if value == "" {
			missingVars = append(missingVars, name)
		}
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("missing required configuration variables: %v", missingVars)
	}

	return nil
}
