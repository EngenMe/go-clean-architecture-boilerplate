package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/mehdihadeli/go-mediatr"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/EngenMe/go-clean-architecture/api/routes"
	"github.com/EngenMe/go-clean-architecture/application/services"
	"github.com/EngenMe/go-clean-architecture/infrastructure/database"
	"github.com/EngenMe/go-clean-architecture/infrastructure/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	utils.LoadConfig()

	// Set Gin mode based on environment
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize MediatR
	mediatr.ClearRequestRegistrations()
	mediatr.ClearNotificationRegistrations()

	// Connect to database
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	userRepository := database.NewPostgresUserRepository(db)

	// Register services
	userService := services.RegisterUserService(userRepository)
	authService := services.RegisterAuthService(userRepository)

	// Configure Gin
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, authService, userService)

	// Get port from environment
	port := utils.GetEnv("PORT", "8080")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server is running on port %s", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(
			err,
			http.ErrServerClosed,
		) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Give server 5 seconds to shut down gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
