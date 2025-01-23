package creator

import (
	"fmj/config"
	"fmj/internal/models"
	"fmj/middleware"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
	config  *config.Config
}

func NewUserHandler(service Service, config *config.Config) *Handler {
	return &Handler{service: service, config: config}
}

func (h *Handler) RegisterUserRoutes(r *gin.RouterGroup) {
	user := r.Group("/creator")
	user.Use(middleware.AuthRequired())
	{
		user.POST("/setup-creator", h.SetupCreator)
	}
}

// SetupCreator Example:
// POST /api/v1/creator/setup-profile
// Request Body:
//
//	{
//	  "username": "",
//	  "name": "",
//	  "bio": "",
//	  "country": "",
//	  "payment_method": [],
//	}
//
// Response:
//
//	{
//	  "data": {},
//	}
//
// @Summary Setup User Profile
// @Description Sets up creators profile data. Payment method must be an array of "MTN", "Paystack", "FlutterWave" or "Stripe"
// @Tags user
// @Accept json
// @Produce json
// @Param SetupUserInputs body models.SetupUserInputs true "SetupUserInputs"
// @Success 200
// @Failure 400
// @Router /api/v1/creator/setup-profile [post]
func (h *Handler) SetupCreator(c *gin.Context) {
	var data *models.SetupUserInputs

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("Error occurred")
		return
	}

	user, err := h.service.SetupUserProfile(c, data.Email, data.Username, data.Name, data.Bio, data.Country, data.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetCreator Example:
// GET /api/v1/creator/:username
//
// Response:
//
//	{
//	  "data": {},
//	}
//
// @Summary Get Creator data
// @Description This endpoint returns all creator data. Used for creator profile.
// @Tags Creator
// @Produce json
// @Success 200
// @Failure 400
// @Router /api/v1/creator/:username [get]
func (h *Handler) GetCreator(c *gin.Context) {
	username := c.Param("username")

	// Get visitor's IP address
	visitorIP := c.ClientIP()

	// First update analytics
	if err := h.service.UpdateAnalytics(username, visitorIP); err != nil {
		slog.Error("Failed to update analytics", "error", err)
		// Continue execution - don't fail the request just because analytics failed
	}

	user, err := h.service.FindCreatorByUsername(username)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		slog.Error(err.Error())
		slog.String("username", username)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
