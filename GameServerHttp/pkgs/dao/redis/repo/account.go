package repo_redis

import (
	entity_pgsql "SlotGameServer/pkgs/dao/postgresql/entity"
	entity_redis "SlotGameServer/pkgs/dao/redis/entity"
	"SlotGameServer/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type userBaseRedisRepo struct {
	redisClient *redis.Client
}

type UserRedisRepo interface {
	// 保存用户
	SetUser(ctx *gin.Context, userBase *entity_pgsql.Account, expiration time.Duration) error
	// 获取用户
	GetUser(ctx *gin.Context, userId uint64) (*entity_pgsql.Account, error)
	// 获取Field
	GetUserField(ctx *gin.Context, userId uint64, field string) (string, error)
}

func NewUserRedisRepo(redisClient *redis.Client) UserRedisRepo {
	return &userBaseRedisRepo{
		redisClient: redisClient,
	}
}

// 保存用户
func (r *userBaseRedisRepo) SetUser(ctx *gin.Context, userBase *entity_pgsql.Account, expiration time.Duration) error {
	if r.redisClient == nil || userBase == nil || userBase.Id <= 0 {
		return utils.ErrParameter
	}

	key := fmt.Sprintf("%s:%d", entity_redis.RedisUser, userBase.Id)
	err := r.redisClient.HSet(ctx, key, userBase).Err()
	if err != nil {
		return err
	}

	err = r.redisClient.Expire(ctx, key, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// 获取用户
func (r *userBaseRedisRepo) GetUser(ctx *gin.Context, userId uint64) (*entity_pgsql.Account, error) {
	if r.redisClient == nil || userId <= 0 {
		return nil, utils.ErrParameter
	}

	tmpUser := &entity_pgsql.Account{}
	err := r.redisClient.HGetAll(ctx, fmt.Sprintf("%s:%d", entity_redis.RedisUser, userId)).Scan(tmpUser)
	if err != nil {
		if err == redis.Nil {
			return nil, utils.ErrRedisNotKey
		}
	}
	if tmpUser.Id <= 0 {
		return nil, utils.ErrRedisNotKey
	}
	return tmpUser, nil
}

// 获取Field
func (r *userBaseRedisRepo) GetUserField(ctx *gin.Context, userId uint64, field string) (string, error) {
	if r.redisClient == nil || userId <= 0 || len(field) <= 0 {
		return "", utils.ErrParameter
	}

	result, err := r.redisClient.HGet(ctx, fmt.Sprintf("%s:%d", entity_redis.RedisUser, userId), field).Result()
	if err != nil {
		if err == redis.Nil {
			return "", utils.ErrRedisNotKey
		}
		return "", err
	}

	return result, nil
}
