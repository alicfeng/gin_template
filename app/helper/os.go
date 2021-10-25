package helper

import "os"

// GetEnvDefault 获取环境变量(支持默认值) /**
func GetEnvDefault(key string, def string) (val string) {
	// 1.获取环境变量
	val = os.Getenv(key)

	// 2.环境变量值为空则返回默认值
	if "" == val {
		return def
	}

	// 3.返回环境变量值
	return val
}
