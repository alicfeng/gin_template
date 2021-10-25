package middleware

import (
	"github.com/gin-gonic/gin"
)

// CommunicateLogMiddleware 记录通讯请求与响应报文日志 TODO.SDK*/
func CommunicateLogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 记录请求日志
		ctx.Next()
		// 记录响应日志
	}
}
