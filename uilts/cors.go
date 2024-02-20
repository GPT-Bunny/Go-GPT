package cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware 创建一个Gin中间件以处理CORS请求
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里将 "http://localhost:5173" 替换为您的前端应用的实际域名
		// 如果您正在本地开发，那么这个值可能是正确的
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

		// 确保添加了 "GET" 方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		// 这里列出了您愿意接受的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 如果是预检请求，直接返回200
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// 处理实际请求
		c.Next()
	}
}
