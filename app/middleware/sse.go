package middleware

import "github.com/gin-gonic/gin"

// SSEEventStreamMiddleware 设置内容头部为SSE
func SSEEventStreamMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/event-stream")
		ctx.Next()
	}
}
