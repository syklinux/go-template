package middleware

import (
	"errors"

	"github.com/syklinux/golib/log"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//先做一下日志记录
				log.Errorf("_com_panic err", err)
				ResponseStatusError(c, 500, errors.New("内部错误"))
				// ResponseError(c, 500, errors.New(fmt.Sprint(err)))
				return
			}
		}()
		c.Next()
	}
}
