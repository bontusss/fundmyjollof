package main

import (
	"context"
	"fmj/config"

	"log"
	"log/slog"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.NewConfig()

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatal("Configuration error: ", err)
	}

	// Log configuration details
	slog.Info("Configuration loaded successfully",
		"environment", cfg.Environment,
		"database_url", cfg.MongoURI,
		"database_name", cfg.DatabaseName,
		"session_secret_length", len(cfg.SessionSecret),
		"smtp_host", cfg.SMTPHost,
		"base_url", cfg.BaseURL,
		"google_callback_url", cfg.GoogleCallbackURL,
	)

	// Connect to MongoDB using the configuration URI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, ctx)

	db := client.Database(cfg.DatabaseName)

	// Run your server.
	if err := runServer(db, cfg); err != nil {
		slog.Error("Failed to start server!", "details", err.Error())
		os.Exit(1)
	}
}
