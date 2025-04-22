package handlers

import (
	"net/http"

	"github.com/EngenMe/go-clean-architecture/application/services"
	"github.com/EngenMe/go-clean-architecture/infrastructure/utils"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication-related requests
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var request services.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewAPIError(http.StatusBadRequest, err.Error()),
		)
		return
	}

	response, err := h.authService.Login(c.Request.Context(), request)
	if err != nil {
		status := utils.ErrorToStatusCode(err)
		c.JSON(status, utils.NewAPIError(status, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

// SignUp handles user registration
func (h *AuthHandler) SignUp(c *gin.Context) {
	var request services.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewAPIError(http.StatusBadRequest, err.Error()),
		)
		return
	}

	response, err := h.authService.SignUp(c.Request.Context(), request)
	if err != nil {
		status := utils.ErrorToStatusCode(err)
		c.JSON(status, utils.NewAPIError(status, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// RegisterRoutes registers authentication routes
func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
		authGroup.POST("/signup", h.SignUp)
	}

	// Add standalone signup route at the top level for better discoverability
	router.POST("/signup", h.SignUp)
}
