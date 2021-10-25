package bootstrap

import (
	"github.com/alicfeng/gin_template/app/repository"
	"github.com/kataras/golog"
)

func init() {
	golog.Debug("ready to connect ai_xr_mysql server")
	err := repository.ConnectMySQL()

	if err != nil {
		panic("failure to connect database")
	}
	golog.Info("connect ai_xr_mysql successful")
}
