package main

import (
	"github.com/alicfeng/gin_template/app/repository"
	"github.com/alicfeng/gin_template/database/migration"
	"github.com/alicfeng/gin_template/database/seeder"
	"github.com/kataras/golog"
)

func main() {
	// 1.连接数据库
	err := repository.ConnectMySQL()
	if nil != err {
		golog.Error("connect to mysql failure", err)
		return
	}
	database := repository.GetMySQL()

	// 1.初始化数据库结构
	migration.Handle(database)

	// 2.迁移数据库数据
	seeder.Handle(database)
}
