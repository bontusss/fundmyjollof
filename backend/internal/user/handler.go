package user

import (
	"fmj/config"
	"fmj/internal/models"
	"fmj/middleware"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type Handler struct {
	service Service
	config  *config.Config
}

func NewUserHandler(service Service, config *config.Config) *Handler {
	return &Handler{service: service, config: config}
}

func (h *Handler) RegisterUserRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	user.Use(middleware.AuthRequired())
	{
		user.POST("/setup-profile", h.SetupProfile)
	}
}

// SetupProfile Example:
// POST /api/v1/auth/register
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
// @Router /api/v1/user/setup-profile [post]
func (h *Handler) SetupProfile(c *gin.Context) {
	var data *models.SetupUserInputs

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("Error occurred")
		return
	}

	user, err := h.service.SetupUserProfile(c, data.Username, data.Name, data.Bio, data.Country, data.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
