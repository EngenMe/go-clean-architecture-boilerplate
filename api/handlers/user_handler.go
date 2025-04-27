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
// @Summary Create a new user
// @Description Creates a new user (public endpoint)
// @Tags Users
// @Accept json
// @Produce json
// @Param command body commands.CreateUserCommand true "User creation details"
// @Success 201 {object} entities.UserDTO
// @Failure 400 {object} utils.APIError
// @Failure 409 {object} utils.APIError
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var command commands.CreateUserCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewAPIError(http.StatusBadRequest, err.Error()),
		)
		return
	}

	user, err := h.userService.CreateUser(c.Request.Context(), command)
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, utils.NewAPIError(statusCode, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUserByID gets a user by ID
// @Summary Get user by ID
// @Description Retrieves a user by their ID (protected endpoint)
// @Tags Users
// @Produce json
// @Param id path uint true "User ID"
// @Security BearerAuth
// @Success 200 {object} entities.UserDTO
// @Failure 400 {object} utils.APIError
// @Failure 401 {object} utils.APIError
// @Failure 404 {object} utils.APIError
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewAPIError(http.StatusBadRequest, "Invalid user ID"),
		)
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, utils.NewAPIError(statusCode, err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByEmail gets a user by email
// @Summary Get user by email
// @Description Retrieves a user by their email address (public endpoint)
// @Tags Users
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} entities.UserDTO
// @Failure 400 {object} utils.APIError
// @Failure 404 {object} utils.APIError
// @Router /api/v1/users/email/{email} [get]
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(
			http.StatusBadRequest,
			utils.NewAPIError(http.StatusBadRequest, "Email is required"),
		)
		return
	}

	user, err := h.userService.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, utils.NewAPIError(statusCode, err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers gets all users
// @Summary Get all users
// @Description Retrieves a list of all users (protected endpoint)
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {array} entities.UserDTO
// @Failure 401 {object} utils.APIError
// @Failure 500 {object} utils.APIError
// @Router /api/v1/users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers(c.Request.Context())
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, utils.NewAPIError(statusCode, err.Error()))
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser updates a user
// @Summary Update user
// @Description Updates a user's information (protected endpoint)
// @Tags Users
// @Accept json
// @Produce json
// @Param id path uint true "User ID"
// @Param command body commands.UpdateUserCommand true "User update details"
// @Security BearerAuth
// @Success 200 {object} entities.UserDTO
// @Failure 400 {object} utils.APIError
// @Failure 401 {object} utils.APIError
// @Failure 404 {object} utils.APIError
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var command commands.UpdateUserCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewAPIError(http.StatusBadRequest, err.Error()),
		)
		return
	}

	// Verify ID in path matches ID in body
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || uint(id) != command.ID {
		c.JSON(
			http.StatusBadRequest,
			utils.NewAPIError(
				http.StatusBadRequest,
				"ID in path must match ID in body",
			),
		)
		return
	}

	user, err := h.userService.UpdateUser(c.Request.Context(), command)
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, utils.NewAPIError(statusCode, err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user
// @Summary Delete user
// @Description Deletes a user by ID (protected endpoint)
// @Tags Users
// @Param id path uint true "User ID"
// @Security BearerAuth
// @Success 204
// @Failure 400 {object} utils.APIError
// @Failure 401 {object} utils.APIError
// @Failure 404 {object} utils.APIError
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewAPIError(http.StatusBadRequest, "Invalid user ID"),
		)
		return
	}

	err = h.userService.DeleteUser(c.Request.Context(), uint(id))
	if err != nil {
		statusCode := utils.ErrorToStatusCode(err)
		c.JSON(statusCode, utils.NewAPIError(statusCode, err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// RegisterRoutes registers user-related routes
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
