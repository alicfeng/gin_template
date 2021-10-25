package controller

import (
	"github.com/alicfeng/gin_template/app/request"
	service "github.com/alicfeng/gin_template/app/service/demo"
	"github.com/gin-gonic/gin"
)

// Login 登录 /**
func Foo(context *gin.Context, request *request.FooRequest) {
	svc := service.NewAlgorithmService(context)

	svc.Foo(request)
}
