package main

import (
	"context"
	"fmj/config"
	"fmj/internal/auth"
	"fmj/internal/email"
	"fmj/middleware"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

// runServer runs a new HTTP server with the loaded environment variables.
func runServer(db *mongo.Database, cfg *config.Config) error {
	port := cfg.Port

	// Initialize services
	emailService := email.NewService(cfg)
	authRepo := auth.NewRepository(db, context.Context(context.Background()))
	authService := auth.NewService(authRepo, emailService)
	authHandler := auth.NewHandler(authService, cfg)

	// Create a new gin server.
	router := gin.Default()

	// Setup sessions with authentication key and encryption key
	authKey := []byte(cfg.SessionSecret)
	encryptionKey := []byte(cfg.SessionSecret[:32]) // Use first 32 bytes for encryption
	store := cookie.NewStore(authKey, encryptionKey)

	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24 * 10, // 10 days
		Path:     "/",
		Secure:   cfg.Environment == "production",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	router.Use(sessions.Sessions("auth_session", store))

	// Register auth routes
	authHandler.RegisterRoutes(router)

	// protected ungrouped routes
	protected := router.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/", func (c *gin.Context)  {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Welcome to Fund my jollof",
			})
		})
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	slog.Info("Starting server...",
		"port", port,
		"environment", cfg.Environment,
		"secure_cookies", cfg.Environment == "production",
	)

	return server.ListenAndServe()
}
