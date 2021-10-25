package config

import (
	"github.com/alicfeng/gin_template/app/helper"
	"github.com/kataras/golog"
	"log"
	"os"
)

type logger struct {
	Path   string
	Level  string
	Prefix string
}

var Logger logger

func init() {
	Logger.Level = helper.GetEnvDefault("SG_LOGGER_LEVEL", "debug")
	Logger.Prefix = helper.GetEnvDefault("SG_LOGGER_PREFIX", "samego.server ")
	Logger.Path = helper.GetEnvDefault("SG_LOGGER_PATH", "/usr/local/samego/log/runtime.log")

	// 设置日志级别 && 日志前缀 && 时间格式
	golog.SetLevel(Logger.Level)
	golog.SetPrefix(Logger.Prefix)
	golog.SetTimeFormat("2006-01-02 15:04:05")

	logFile, err := os.OpenFile(Logger.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("open log file error")
	}
	golog.SetOutput(logFile)
}
