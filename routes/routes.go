package routes

import (
	"synergy/controllers"
	"synergy/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Admin routes
	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	{
		admin.POST("/job", controllers.CreateJob)
	}
}
