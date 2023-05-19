package middleware

import (
	"GoTemplate/src/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		w := c.Writer
		origin := c.Request.Header.Get("Origin")
		whiteList := utils.Conf.HTTPConf.Cors
		for _, domain := range whiteList {
			if strings.Contains(origin, domain) {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
				w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
				w.Header().Add("Access-Control-Allow-Headers", "Access-Token")
				w.Header().Add("Access-Control-Allow-Headers", "x-requested-with")
				w.Header().Add("Access-Control-Allow-Headers", "X-Alert-API-Version")
				w.Header().Add("Access-Control-Allow-Headers", "Authorization")
				w.Header().Add("Access-Control-Allow-Headers", "cache-control")
				w.Header().Add("Access-Control-Expose-Headers", "Content-Disposition")
			}
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		c.Next()
	}
}
