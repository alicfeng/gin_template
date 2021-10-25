package repository

import (
	"github.com/alicfeng/gin_template/config"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _mysql *gorm.DB
var _redis *redis.Client

// ConnectMySQL 连接MySQL数据库*/
func ConnectMySQL() (err error) {
	_mysql, err = gorm.Open(mysql.Open(config.Mysql.GetDsn()), &gorm.Config{})

	return err
}

// GetMySQL 获取MySQL数据库句柄 */
func GetMySQL() *gorm.DB {
	return _mysql
}

func ConnectRedis() (err error) {
	_redis = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Auth,
		DB:       0,
	})

	return err
}

// GetRedis 获取Redis连接句柄 */
func GetRedis() *redis.Client {
	return _redis
}
