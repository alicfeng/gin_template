package config

import (
	"github.com/alicfeng/gin_template/app/helper"
)

type server struct {
	Mode string // 运行模式
	Host string // 内部地址
	URL  string // 通信地址
}

var Server server

func init() {
	Server.Mode = helper.GetEnvDefault("SG_SERVER_MODE", "debug")
	Server.Host = helper.GetEnvDefault("SG_SERVER_HOST", "0.0.0.0:80")
	Server.URL = helper.GetEnvDefault("SG_SERVER_URL", "http://ai-rx-api.ai-app.svc.cluster.local:80")
}
