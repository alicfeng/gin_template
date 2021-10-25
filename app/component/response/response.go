package component

import (
	"github.com/alicfeng/gin_template/app/enum"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Status int
	Body   body
}

type Request struct {
	Id string `json:"uuid"`
}

type body struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Request Request     `json:"request"`
}

// Response 虚拟构造函数 /**
func Response() *Result {
	return options()
}

// StatusCode 设置 web 响应状态码 /**
func (result *Result) StatusCode(status int) *Result {
	result.Status = status

	return result
}

// options 默认响应实体 */
func options() *Result {
	return &Result{
		Status: http.StatusOK,
		Body: body{
			Code:    enum.SUCCESS.Code,
			Data:    struct{}{},
			Message: enum.SUCCESS.Message,
			Request: Request{
				Id: "",
			},
		},
	}
}

// Enum 设置报文 package.{code|message} 信息 /**
func (result *Result) Enum(enum enum.CodeEnum) *Result {
	result.Body.Code = enum.Code
	result.Body.Message = enum.Message

	return result
}

// Data 设置报文 package.data 信息 /**
func (result *Result) Data(data interface{}) *Result {
	result.Body.Data = data

	return result
}

// RequestId 设置报文 package.request.id 信息 /**
func (result *Result) RequestId(requestId string) *Result {
	result.Body.Request.Id = requestId

	return result
}

// Send 基于 gin.context 吐出响应报文 /**
func (result *Result) Send(context *gin.Context) {

	context.JSON(result.Status, result.Body)
}
