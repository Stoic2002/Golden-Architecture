package handler

import (
	"github.com/arulkarim/golden-architecture/internal/infrastructure/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers auth routes
func RegisterRoutes(router *gin.RouterGroup, handler *Handler, jwtManager *auth.JWTManager) {
	authGroup := router.Group("/auth")
	{
		// Public routes
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/login", handler.Login)

		// Protected routes
		authGroup.GET("/profile", auth.AuthMiddleware(jwtManager), handler.Profile)
	}
}
