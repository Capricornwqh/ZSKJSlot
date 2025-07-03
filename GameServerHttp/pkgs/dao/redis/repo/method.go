package redis_repo

import (
	"SlotGameServer/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// 保存md5Id
func SetMethodMD5(ctx *gin.Context, key, value string, expiration time.Duration) error {
	if utils.RedisClient == nil || len(key) <= 0 {
		return utils.ErrParameter
	}

	err := utils.RedisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// 获取md5Id
func GetMethodMD5(ctx *gin.Context, key string) (string, error) {
	if utils.RedisClient == nil || len(key) <= 0 {
		return "", utils.ErrParameter
	}

	value, err := utils.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

// 删除md5Id
func DelMethodMD5(ctx *gin.Context, key string) error {
	if utils.RedisClient == nil || len(key) <= 0 {
		return utils.ErrParameter
	}

	err := utils.RedisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
