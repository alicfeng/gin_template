package service

import (
	"github.com/alicfeng/gin_template/app/component/response"
	"github.com/alicfeng/gin_template/app/enum"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Context *gin.Context
}

// Failure 处理失败响应 /**
func (service *Service) Failure(code enum.CodeEnum, datum ...interface{}) {
	Response(datum, service.Context).Enum(code).Send(service.Context)
}

// Success 处理成功响应 /**
func (service *Service) Success(datum ...interface{}) {
	Response(datum, service.Context).Send(service.Context)
}

// Error 处理异常响应 /**
func (service *Service) Error(statusCode int, code enum.CodeEnum, datum ...interface{}) {
	Response(datum, service.Context).StatusCode(statusCode).Enum(code).Send(service.Context)
}

// Response 实例化 component.Result 并初始化 data 值 */
func Response(data []interface{}, context *gin.Context) (response *component.Result) {
	// case.当没有数据时
	if nil == data {
		response = component.Response()
	}

	// case.当只有一个数据时
	if 1 == len(data) {
		response = component.Response().Data(data[0])
	}

	// case.当有很多数据时
	if len(data) > 1 {
		response = component.Response().Data(data)
	}

	return response.RequestId(context.Request.Header.Get(enum.RequestIdFieldName))
}
