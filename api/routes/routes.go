package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/EngenMe/go-clean-architecture/api/handlers"
	"github.com/EngenMe/go-clean-architecture/api/middlewares"
	"github.com/EngenMe/go-clean-architecture/application/services"
)

// SetupRoutes configures all API routes
func SetupRoutes(
	router *gin.Engine,
	authService *services.AuthService,
	userService *services.UserService,
) {
	// Register global middlewares
	router.Use(middlewares.LoggingMiddleware())

	// Create an API group
	api := router.Group("/api/v1")

	// Register auth routes (no auth middleware - these are public endpoints)
	authHandler := handlers.NewAuthHandler(authService)
	authHandler.RegisterRoutes(api)

	// Register user routes with auth middleware
	userHandler := handlers.NewUserHandler(userService)
	userHandler.RegisterRoutes(api, middlewares.AuthMiddleware())

	// Health check route
	router.GET(
		"/health", func(c *gin.Context) {
			c.JSON(
				200, gin.H{
					"status": "up",
				},
			)
		},
	)
}
