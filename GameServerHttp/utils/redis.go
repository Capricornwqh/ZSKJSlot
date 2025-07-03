package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

const Nil = redis.Nil

// 初始化Redis
func SetupRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         Conf.Redis.Addr,
		DB:           Conf.Redis.DB,
		Password:     Conf.Redis.Password,
		ReadTimeout:  time.Duration(Conf.Redis.Timeout * int(time.Second)),
		WriteTimeout: time.Duration(Conf.Redis.Timeout * int(time.Second)),
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Redis Server Open: db: %d, address: %s", Conf.Redis.DB, Conf.Redis.Addr)
}

func CloseRedis() error {
	if RedisClient != nil {
		return RedisClient.Close()
	}
	return nil
}
