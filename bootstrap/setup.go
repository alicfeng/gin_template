package bootstrap

import (
	"github.com/alicfeng/gin_template/app/middleware"
	"github.com/alicfeng/gin_template/app/validator"
	"github.com/alicfeng/gin_template/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type Server struct {
	Engine *gin.Engine
	Cron   *cron.Cron
}

// New 创建服务应用 /**
func New() *Server {
	return &Server{
		Engine: gin.Default(),
		Cron:   cron.New(),
	}
}

// Run 运行服务应用 /**
func (server *Server) Run() {
	// 1.设定运行模式 debug release test
	gin.SetMode(config.Server.Mode)

	// 2.注册运行定时任务服务
	//server.registerCronTask().Start()

	// 3.使用捆绑中间件
	server.Engine.Use(middleware.CommunicateLogMiddleware())
	server.Engine.Use(middleware.RequestTraceIdMiddleware())
	server.Engine.Use(middleware.AuthenticTokenMiddleware())
	server.Engine.Use(cors.Default())

	// 4.引擎注册路由
	server.registerRoute()

	// 5.注册校验规则
	_ = validator.Register()

	// 6.运行引擎服务
	_ = server.Engine.Run(config.Server.Host)
}
