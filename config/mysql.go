package config

import (
	"fmt"
	"github.com/alicfeng/gin_template/app/helper"
)

type mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Charset  string
	Database string
}

const (
	// DSN 格式
	dsnFormat = "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local"
)

var Mysql mysql

func init() {
	Mysql.Host = helper.GetEnvDefault("SG_MYSQL_HOST", "127.0.0.1")
	Mysql.Port = helper.GetEnvDefault("SG_MYSQL_PORT", "3306")
	Mysql.Username = helper.GetEnvDefault("SG_MYSQL_USERNAME", "root")
	Mysql.Password = helper.GetEnvDefault("SG_MYSQL_PASSWORD", "root")
	Mysql.Charset = helper.GetEnvDefault("SG_MYSQL_CHARSET", "utf8mb4")
	Mysql.Database = helper.GetEnvDefault("SG_MYSQL_DATABASE", "demo")
}

// GetDsn 获取数据库连接地址*/
func (db *mysql) GetDsn() string {
	return fmt.Sprintf(dsnFormat, Mysql.Username, Mysql.Password, Mysql.Host, Mysql.Port, Mysql.Database, Mysql.Charset)
}
