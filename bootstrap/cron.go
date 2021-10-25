package bootstrap

import (
	"github.com/robfig/cron/v3"
)

// registerCronTask 注册定时任务 /**
func (server *Server) registerCronTask() *cron.Cron {

	return server.Cron
}
