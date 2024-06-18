package routes

import (
	"media-app-go/controllers"

	"github.com/gin-gonic/gin"
)

// Add media apis
func Routes(router *gin.Engine) {
	// Create groups for apis of path '/api'
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/help", controllers.GetHelp)
	}
}
