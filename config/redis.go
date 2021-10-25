package config

import (
	"github.com/alicfeng/gin_template/app/helper"
)

type redis struct {
	Host string
	Port string
	Auth string
}

var Redis redis

func init() {
	Redis.Host = helper.GetEnvDefault("SG_REDIS_HOST", "127.0.0.1")
	Redis.Port = helper.GetEnvDefault("SG_REDIS_PORT", "6379")
	Redis.Auth = helper.GetEnvDefault("SG_REDIS_AUTH", "")
}
