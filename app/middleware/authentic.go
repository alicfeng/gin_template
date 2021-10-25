package middleware

import (
	authorization "github.com/alicfeng/gin_template/app/component/jwt"
	response "github.com/alicfeng/gin_template/app/component/response"
	"github.com/alicfeng/gin_template/app/enum"
	"github.com/alicfeng/gin_template/app/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 不鉴权的白名单URI
var uriWhiteList = []string{
	"/health/ping",
	"/health/liveness",
	"/health/readness",
}

// AuthenticTokenMiddleware 路由中间件.加载请求链路追踪标识 /**
func AuthenticTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1.白名单过滤
		if _, exist := helper.SliceContain(uriWhiteList, ctx.Request.RequestURI); exist {
			ctx.Next()

			return
		}

		// 2.判断登录态是否为空 && 登录态是否有效
		token := ctx.Request.Header.Get(enum.RequestHeaderAuthorizationFieldName)
		auth, err := authorization.ParseToken(token)
		if "" == token || err != nil {
			response.Response().
				StatusCode(http.StatusUnauthorized).
				RequestId(ctx.Request.Header.Get(enum.RequestIdFieldName)).
				Enum(enum.CodeEnum{
					Code:    enum.FailureCode,
					Message: "请登录后再尝试请求",
				}).Send(ctx)

			ctx.Abort()
			return
		}

		// 3.判断
		ctx.Set("user_id", auth.UserId)

		ctx.Next()
	}
}
