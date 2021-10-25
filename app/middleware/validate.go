package middleware

import (
	response "github.com/alicfeng/gin_template/app/component/response"
	"github.com/alicfeng/gin_template/app/enum"
	"github.com/alicfeng/gin_template/app/helper"
	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
	"net/http"
	"reflect"
)

// StdHandle 统一标准化注册路由处理控制器 所有控制器处理函数务必统一调用绑定 /**
func StdHandle(handler interface{}) func(*gin.Context) {
	// 1.基于反射获取函数与参数信息
	funcValue := reflect.ValueOf(handler)
	funcType := reflect.TypeOf(handler)

	// 2.判断是否函数类型值是否为函数
	if funcType.Kind() != reflect.Func {
		golog.Error("the route handler function must be function : %s", funcValue.String())

		panic("the route handler function must be function " + funcValue.String())
	}

	// 3.调用代理中间层 隐式触发控制器
	return func(ctx *gin.Context) {
		// 3.1当参数只有一个 则无 request 注入
		if 1 == reflect.TypeOf(handler).NumIn() {
			ctx.Next()
		}
		// 3.2当参数为两个 则认为注入 request
		if 2 == reflect.TypeOf(handler).NumIn() {
			request, pass := validate(ctx, funcType.In(1).Elem())
			if false == pass {
				return
			}

			// 3.2 注入表单参数触发Handler
			funcValue.Call(helper.BuildReflectValueParams(ctx, request))
		}
	}
}

// validate 路由中间件.请求参数校验 /**
func validate(ctx *gin.Context, typeParam reflect.Type) (interface{}, bool) {
	// 1.基于反射机制实例化请求Request
	request := reflect.New(typeParam).Interface()

	// 2.参数校验
	var err error
	if err = ctx.ShouldBind(request); err == nil {
		return request, true
	}
	if err = ctx.ShouldBindUri(request); err == nil {
		return request, true
	}

	// 2.1不通过直接响应
	requestValue := reflect.ValueOf(request)
	translateParam := helper.BuildReflectValueParams(requestValue, err)
	response.Response().
		StatusCode(http.StatusUnprocessableEntity).
		RequestId(ctx.Request.Header.Get(enum.RequestIdFieldName)).
		Enum(enum.CodeEnum{
			Code:    enum.FailureCode,
			Message: requestValue.MethodByName("Translate").Call(translateParam)[0].String(),
		}).Send(ctx)

	return nil, false
}
