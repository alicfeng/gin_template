package bootstrap

import (
	"github.com/alicfeng/gin_template/app/repository"
	"github.com/kataras/golog"
)

func init() {
	golog.Debug("ready to connect ai_xr_redis server")
	err := repository.ConnectRedis()

	if nil != err {
		golog.Error("ai_xr_redis auth error %v", err)
		panic("ai_xr_redis auth error")
	}

	golog.Info("connect ai_xr_redis successful")
}
