package route

import (
	controller "github.com/alicfeng/gin_template/app/controller/demo"
	"github.com/gin-gonic/gin"
)

// RegisterDemo Demo/**
func RegisterDemo(engine *gin.Engine) {
	user := engine.Group("demo") /*Demo*/
	{
		auth := user.Group("hello") /*Hello*/
		{
			auth.POST("login", StdHandle(controller.Foo)) // Foo
		}
	}
}
