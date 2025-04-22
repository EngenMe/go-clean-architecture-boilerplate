package middlewares

import (
	"net/http"
	"strings"

	"github.com/EngenMe/go-clean-architecture/infrastructure/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware for JWT authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				utils.NewAPIError(
					http.StatusUnauthorized,
					"Authorization header is required",
				),
			)
			return
		}

		// Check if the header has the Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				utils.NewAPIError(
					http.StatusUnauthorized,
					"Authorization header must be Bearer token",
				),
			)
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				utils.NewAPIError(
					http.StatusUnauthorized,
					"Invalid or expired JWT token",
				),
			)
			return
		}

		// Set user ID and email in the context
		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)

		c.Next()
	}
}
