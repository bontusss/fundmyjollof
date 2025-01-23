package auth

import (
	"fmj/config"
	"fmj/internal/models"
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

func (h *Handler) RegisterAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
		auth.GET("/logout", h.Logout)
		//auth.GET("/:provider", h.GoogleLogin)
		//auth.GET("/:provider/callback", h.GoogleCallback)
		auth.POST("/forgot-password", h.ForgotPassword)
		auth.POST("/reset-password", h.ResetPassword)
		auth.POST("/verify", h.VerifyEmail)
	}
}

// Login Example:
// POST /api/v1/auth/login
// Request Body:
//
//	{
//	  "email": "newuser@example.com",
//	  "password": "password",
//	}
//
// Response:
//
//	{
//	  "data": *models.User,
//	}
//
// @Summary Login User
// @Description Logs a user in and creates a session
// @Tags auth
// @Accept json
// @Produce json
// @Param RegisterInputs body models.RegisterInputs true "RegisterInputs"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var data *models.RegisterInputs

	if err := c.ShouldBind(&data); err != nil {
		slog.Error("Failed to bind login data", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add debug logging
	slog.Info("Attempting login",
		slog.String("email", data.Email),
		slog.String("pass_length", fmt.Sprintf("%d", len(data.Pass))))

	user, err := h.service.Login(data.Email, data.Pass)
	if err != nil {
		slog.Error("Login failed",
			slog.String("email", data.Email),
			slog.String("error", err.Error()))

		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
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

// Register Example:
// POST /api/v1/auth/register
// Request Body:
//
//	{
//	  "email": "newuser@example.com",
//	  "password": "password",
//	}
//
// Response:
//
//	{
//	  "Success": "Registration successful! Please check your email to verify your account.",
//	}
//
// @Summary Register User
// @Description Register a user
// @Tags auth
// @Accept json
// @Produce json
// @Param RegisterInputs body models.RegisterInputs true "RegisterInputs"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /api/v1/auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var user *models.RegisterInputs
	fmt.Println("registering user with email :")

	if err := c.ShouldBind(&user); err != nil {
		slog.Error("Failed to bind user data", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Provide the required fields to register.",
		})
		return
	}

	if err := h.service.Register(c, user.Email, user.Pass); err != nil {
		if err.Error() == "email already registered" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "You already have an account, log in instead",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "An error occurred, try again.",
		})
		slog.Error("Error registering user in database", slog.String("email", user.Email), slog.String("error", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Success": "Registration successful! Please check your email to verify your account.",
	})
}

// VerifyEmail Example:
// POST /api/v1/auth/verify
// Request Body:
//
//	{
//	  "code": "",
//	}
//
// Response:
//
//	{
//	  "Success": "Email verified successfully! You can now login.",
//	}
//
// @Summary Verify User Email
// @Description Verifies a users email
// @Tags auth
// @Accept json
// @Produce json
// @Param VerifyEmailInputs body models.VerifyEmailInputs true "VerifyEmailInputs"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/auth/verify [post]
func (h *Handler) VerifyEmail(c *gin.Context) {
	var data *models.VerifyEmailInputs
	err := c.ShouldBind(&data)
	if err != nil {
		slog.Error("Failed to bind user data", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Provide the required fields to verify your email.",
		})
		return
	}
	// Convert string code to uint32
	codeUint, err := strconv.ParseUint(data.Code, 10, 32)
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
