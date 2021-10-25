package bootstrap

import "github.com/alicfeng/gin_template/route"

// registerRoute 注册路由 /**
func (server *Server) registerRoute() {
	// Demo
	route.RegisterDemo(server.Engine)
}
