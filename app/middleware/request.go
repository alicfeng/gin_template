package middleware

import (
	"github.com/alicfeng/gin_template/app/enum"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// RequestTraceIdMiddleware 路由中间件.加载请求链路追踪标识 /**
func RequestTraceIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if "" == ctx.Request.Header.Get(enum.RequestIdFieldName) {
			ctx.Request.Header.Set(enum.RequestIdFieldName, uuid.Must(uuid.NewV4()).String())
		}

		ctx.Next()
	}
}
