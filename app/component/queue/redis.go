package aqueue

import (
	"github.com/alicfeng/gin_template/app/repository"
	"github.com/go-redis/redis"
	"github.com/kataras/golog"
)

type Queue struct {
	redis *redis.Client
}

// callable 定义订阅回调函数类型
type callable func(channel string, message string, v interface{})

// New 虚拟构造 /**
func New() *Queue {
	return &Queue{
		redis: repository.GetRedis(),
	}
}

// Publish  发布消息 /**
func (queue *Queue) Publish(channel string, message interface{}) (int64, error) {
	return queue.redis.Publish(channel, message).Result()
}

// Subscribe 订阅消息 /**
func (queue *Queue) Subscribe(channel string, callback callable, v interface{}) {
	golog.Infof("subscribing about %s", channel)
	subScribe := queue.redis.Subscribe(channel)
	defer subScribe.Close()

	golog.Infof("subscribed")
	for message := range subScribe.Channel() {
		golog.Infof("subscribe callback")
		callback(message.Channel, message.Payload, v)
	}
}
