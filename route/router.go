package route

import (
	"github.com/alicfeng/gin_template/app/middleware"
	"github.com/gin-gonic/gin"
)

func StdHandle(handler interface{}) func(*gin.Context) {
	return middleware.StdHandle(handler)
}
