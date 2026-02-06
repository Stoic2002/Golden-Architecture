package handler

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers todo routes
func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	todos := router.Group("/todos")
	{
		todos.POST("", handler.Create)
		todos.GET("", handler.GetAll)
		todos.GET("/:id", handler.GetByID)
		todos.PUT("/:id", handler.Update)
		todos.DELETE("/:id", handler.Delete)
	}
}
