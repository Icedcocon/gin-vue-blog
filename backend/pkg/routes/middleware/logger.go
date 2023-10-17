package middleware

import (
	"backend/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 记录请求结束时间
		cost := time.Since(start)
		// 打印日志
		utils.GLogger.Info(c.Request.URL.Path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			// zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
