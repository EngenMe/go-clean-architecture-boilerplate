package handlers

import (
	"net/http"
	"strconv"

	"github.com/EngenMe/go-clean-architecture/application/commands"
	"github.com/EngenMe/go-clean-architecture/application/services"
	"github.com/EngenMe/go-clean-architecture/infrastructure/utils"
	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related requests
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser handles user creation
func (h *UserHandler) CreateUser(c *gin.Context) {
	var command commands.CreateUserCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(c.Request.Context(), command)
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUserByID gets a user by ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByEmail gets a user by email
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	user, err := h.userService.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers gets all users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers(c.Request.Context())
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var command commands.UpdateUserCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify ID in a path matches ID in body
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || uint(id) != command.ID {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID in path must match ID in body"},
		)
		return
	}

	user, err := h.userService.UpdateUser(c.Request.Context(), command)
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.userService.DeleteUser(c.Request.Context(), uint(id))
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// RegisterRoutes RegisterUserRoutes registers user-related routes
func (h *UserHandler) RegisterRoutes(
	router *gin.RouterGroup,
	authMiddleware gin.HandlerFunc,
) {
	users := router.Group("/users")
	{
		// Public routes
		users.POST("", h.CreateUser)
		users.GET("/email/:email", h.GetUserByEmail)

		// Protected routes
		authenticated := users.Group("")
		authenticated.Use(authMiddleware)
		{
			authenticated.GET("", h.GetAllUsers)
			authenticated.GET("/:id", h.GetUserByID)
			authenticated.PUT("/:id", h.UpdateUser)
			authenticated.DELETE("/:id", h.DeleteUser)
		}
	}
}
