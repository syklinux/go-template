package router

import (
	"GoTemplate/src/controller"
	"GoTemplate/src/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter InitRouters
func InitRouter(middleWares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middleWares...)

	health := router.Group("/api/health")
	health.Use(middleware.RecoveryMiddleware())

	{
		controller.HealthyAPIRegister(health)
	}

	return router
}
