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
// @Summary User login
// @Description Authenticates a user and returns a JWT token and user details
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body services.LoginRequest true "Login credentials"
// @Success 200 {object} services.AuthResponse
// @Failure 400 {object} utils.APIError
// @Failure 401 {object} utils.APIError
// @Router /api/v1/auth/login [post]
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
// @Summary User registration
// @Description Registers a new user and returns a JWT token and user details
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body services.SignUpRequest true "Sign-up details"
// @Success 201 {object} services.AuthResponse
// @Failure 400 {object} utils.APIError
// @Failure 409 {object} utils.APIError
// @Router /api/v1/auth/signup [post]
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
	// @Summary User registration (alternative endpoint)
	// @Description Registers a new user and returns a JWT token and user details
	// @Tags Authentication
	// @Accept json
	// @Produce json
	// @Param request body services.SignUpRequest true "Sign-up details"
	// @Success 201 {object} services.AuthResponse
	// @Failure 400 {object} utils.APIError
	// @Failure 409 {object} utils.APIError
	// @Router /api/v1/signup [post]
	router.POST("/signup", h.SignUp)
}
