package auth

import (
	"fmj/config"
	"fmj/internal/utils"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
	config  *config.Config
}

func NewHandler(service Service, cfg *config.Config) *Handler {
	return &Handler{service: service, config: cfg}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
		auth.GET("/logout", h.Logout)
		//auth.GET("/:provider", h.GoogleLogin)
		//auth.GET("/:provider/callback", h.GoogleCallback)
		auth.POST("/forgot-password", h.ForgotPassword)
		auth.POST("/reset-password", h.ResetPassword)
		auth.GET("/verify", h.VerifyEmail)
	}
}

func (h *Handler) Login(c *gin.Context) {
	email := utils.SanitizeInput(c.PostForm("email"))
	password := utils.SanitizeInput(c.PostForm("password"))

	user, err := h.service.Login(email, password)
	if err != nil {
		slog.Error("Error logging a user in database",
			slog.String("email", email),
			slog.String("error", err.Error()))

		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Invalid email or password.",
		})
		return
	}

	// Login succeeded, set session
	session := sessions.Default(c)
	session.Set("user_id", user.ID.Hex())
	if err := session.Save(); err != nil {
		slog.Error("Failed to save session",
			slog.String("user_id", user.ID.Hex()),
			slog.String("error", err.Error()))

		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "An error occurred while starting your session.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *Handler) Register(c *gin.Context) {
	email := utils.SanitizeInput(c.PostForm("email"))
	password := utils.SanitizeInput(c.PostForm("password"))

	fmt.Println("registering user with email: ", email)

	if err := h.service.Register(c, email, password); err != nil {
		if err.Error() == "email already registered" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "You already have an account, log in instead",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "An error occurred, try again.",
		})
		slog.Error("Error registering user in database", slog.String("email", email), slog.String("email", email), slog.String("error", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Success": "Registration successful! Please check your email to verify your account.",
	})
}

func (h *Handler) VerifyEmail(c *gin.Context) {
	code := utils.SanitizeInput(c.PostForm("code"))
	// Convert string code to uint32
	codeUint, err := strconv.ParseUint(code, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid verification code.",
		})
		slog.Error("Error parsing verification code", slog.String("error", err.Error()))
		return
	}

	if err := h.service.VerifyEmail(c, uint32(codeUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "An error occurred, try again.",
		})
		slog.Error("Error verifying user email", slog.String("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "Email verified successfully! You can now login.",
	})
}

func (h *Handler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Success": "you are logged out",
	})
}

func (h *Handler) ForgotPassword(c *gin.Context) {
	email := utils.SanitizeInput(c.PostForm("email"))
	if err := h.service.ForgotPassword(c, email); err != nil {
		slog.Error("Error processing forgot password request", slog.String("email", email), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "An error occurred, try again.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "If an account exists with that email, you will receive password reset instructions.",
	})
}

func (h *Handler) ResetPassword(c *gin.Context) {
	token := utils.SanitizeInput(c.PostForm("token"))
	password := utils.SanitizeInput(c.PostForm("password"))

	if err := h.service.ResetPassword(c, token, password); err != nil {
		slog.Error("Error resetting password", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "An error occurred, try again",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "Password has been reset successfully. You can now login.",
	})
}
