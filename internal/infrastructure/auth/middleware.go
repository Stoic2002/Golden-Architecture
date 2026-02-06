package auth

import (
	"net/http"
	"strings"

	"github.com/arulkarim/golden-architecture/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	// AuthorizationHeader is the header key for authorization
	AuthorizationHeader = "Authorization"
	// BearerPrefix is the prefix for bearer token
	BearerPrefix = "Bearer "
	// ContextUserID is the context key for user ID
	ContextUserID = "userID"
	// ContextUserEmail is the context key for user email
	ContextUserEmail = "userEmail"
)

// AuthMiddleware creates a JWT authentication middleware
func AuthMiddleware(jwtManager *JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization header required", "missing authorization header")
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			response.Error(c, http.StatusUnauthorized, "Invalid authorization format", "authorization header must start with Bearer")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
		claims, err := jwtManager.ValidateToken(tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid token", err.Error())
			c.Abort()
			return
		}

		// Set user info in context
		c.Set(ContextUserID, claims.UserID)
		c.Set(ContextUserEmail, claims.Email)

		c.Next()
	}
}

// GetUserIDFromContext extracts user ID from gin context
func GetUserIDFromContext(c *gin.Context) (uint, bool) {
	userID, exists := c.Get(ContextUserID)
	if !exists {
		return 0, false
	}
	id, ok := userID.(uint)
	return id, ok
}

// GetUserEmailFromContext extracts user email from gin context
func GetUserEmailFromContext(c *gin.Context) (string, bool) {
	email, exists := c.Get(ContextUserEmail)
	if !exists {
		return "", false
	}
	e, ok := email.(string)
	return e, ok
}
