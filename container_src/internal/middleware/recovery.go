package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/pkg/logger"
	"server/pkg/response"
)

// Recovery is a middleware that recovers from panics and logs the error
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic recovered",
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
				)

				response.InternalError(c, "An unexpected error occurred")
				c.Abort()
			}
		}()

		c.Next()
	}
}
