package redis_repo

import (
	redis_entity "SlotGameServer/pkgs/dao/redis/entity"
	"SlotGameServer/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type verifyCodeRedisRepo struct {
	redisClient *redis.Client
}

type VerifyCodeRedisRepo interface {
	// 保存验证码
	SetVerifyCode(ctx *gin.Context, key, code string, expiration time.Duration) error
	// 获取验证码
	GetVerifyCode(ctx *gin.Context, key string) (string, error)
	// 获取并删除验证码
	GetDelVerifyCode(ctx *gin.Context, key string) (string, error)
	// 删除验证码
	DelVerifyCode(ctx *gin.Context, key string) error
}

func NewVerifyCodeRedisRepo(redisClient *redis.Client) VerifyCodeRedisRepo {
	return &verifyCodeRedisRepo{
		redisClient: redisClient,
	}
}

// 保存验证码
func (r *verifyCodeRedisRepo) SetVerifyCode(ctx *gin.Context, key, code string, expiration time.Duration) error {
	if r.redisClient == nil || len(key) <= 0 || len(code) <= 0 {
		return utils.ErrParameter
	}

	err := r.redisClient.Set(ctx, redis_entity.RedisVerifyCode+key, code, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// 获取验证码
func (r *verifyCodeRedisRepo) GetVerifyCode(ctx *gin.Context, key string) (string, error) {
	if r.redisClient == nil || len(key) <= 0 {
		return "", utils.ErrParameter
	}

	value, err := r.redisClient.Get(ctx, redis_entity.RedisVerifyCode+key).Result()
	if err != nil {
		return "", utils.ErrDataNotFound
	}

	return value, nil
}

// 获取并删除验证码
func (r *verifyCodeRedisRepo) GetDelVerifyCode(ctx *gin.Context, key string) (string, error) {
	if r.redisClient == nil || len(key) <= 0 {
		return "", utils.ErrParameter
	}

	value, err := r.redisClient.GetDel(ctx, redis_entity.RedisVerifyCode+key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

// 删除验证码
func (r *verifyCodeRedisRepo) DelVerifyCode(ctx *gin.Context, key string) error {
	if r.redisClient == nil || len(key) <= 0 {
		return utils.ErrParameter
	}

	err := r.redisClient.Del(ctx, redis_entity.RedisVerifyCode+key).Err()
	if err != nil {
		return err
	}

	return nil
}
