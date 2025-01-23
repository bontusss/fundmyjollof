package main

import (
	"context"
	"fmj/config"
	doc "fmj/docs"
	"fmj/internal/auth"
	"fmj/internal/creator"
	"fmj/internal/email"
	"fmj/middleware"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// runServer runs a new HTTP server with the loaded environment variables.
func runServer(db *mongo.Database, cfg *config.Config) error {
	port := cfg.Port

	// Initialize services
	emailService := email.NewService(cfg)
	authRepo := auth.NewRepository(db, context.Context(context.Background()))
	authService := auth.NewService(authRepo, emailService)
	authHandler := auth.NewHandler(authService, cfg)
	userService := creator.NewService(authRepo, emailService)
	userHandler := creator.NewUserHandler(userService, cfg)

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

	// CORS middleware configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	doc.SwaggerInfo.BasePath = "/api/v1"
	apiV1 := router.Group("/api/v1")

	// Register routes
	authHandler.RegisterAuthRoutes(apiV1)
	userHandler.RegisterUserRoutes(apiV1)

	apiV1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	apiV1.GET("/:username", userHandler.GetCreator)

	apiV1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// protected ungrouped routes
	protected := router.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Welcome to Fund my jollof",
			})
		})
	}

	// Serve Nuxt static files in production
	//router.NoRoute(func(c *gin.Context) {
	//	c.File("./dist/index.html")
	//})
	//router.Static("/", "./dist")

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
