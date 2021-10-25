package demo

import (
	"github.com/alicfeng/gin_template/app/request"
	"github.com/alicfeng/gin_template/app/service"
	"github.com/gin-gonic/gin"
)

type HelloService struct {
	service.Service
}

func NewAlgorithmService(context *gin.Context) *HelloService {
	return &HelloService{Service: service.Service{Context: context}}
}

// HelloService /**
func (service *HelloService) Foo(request *request.FooRequest) {
	service.Success(request.Name)
}
