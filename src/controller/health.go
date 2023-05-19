package controller

import (
	"GoTemplate/src/middleware"

	"github.com/gin-gonic/gin"
)

// HealthAPIController HealthAPIController
type HealthAPIController struct{}

// HealthyAPIRegister HealthyAPIRegister
func HealthyAPIRegister(router *gin.RouterGroup) {
	curd := HealthAPIController{}
	router.GET("", curd.Health)
}

// Health Health
func (s *HealthAPIController) Health(c *gin.Context) {
	middleware.ResponseSuccess(c, "ok")
}
