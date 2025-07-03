package redis_repo

import (
	pgsql_entity "SlotGameServer/pkgs/dao/postgresql/entity"
	redis_entity "SlotGameServer/pkgs/dao/redis/entity"
	"SlotGameServer/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type propsRedisRepo struct {
	redisClient *redis.Client
}

type PropsRedisRepo interface {
	// 保存属性
	SetProps(ctx *gin.Context, club *pgsql_entity.Props) error
	// 获取钱包余额
	GetWallet(ctx *gin.Context, uId, cId uint64) (float64, error)
	// 获取访问权限
	GetAL(ctx *gin.Context, uId, cId uint64) (uint64, error)
	// 获取主要返还玩家比率
	GetMRTP(ctx *gin.Context, uId, cId uint64) (float64, error)
}

func NewPropsRedisRepo(redisClient *redis.Client) PropsRedisRepo {
	return &propsRedisRepo{
		redisClient: redisClient,
	}
}

// 保存属性
func (r *propsRedisRepo) SetProps(ctx *gin.Context, props *pgsql_entity.Props) error {
	if r.redisClient == nil || props == nil || props.CId <= 0 || props.UId <= 0 {
		return utils.ErrParameter
	}

	err := r.redisClient.
		HSet(ctx, fmt.Sprintf("%s:%d:%d", redis_entity.RedisProps, props.UId, props.CId), props).
		Err()
	if err != nil {
		return err
	}
	return nil
}

// 获取钱包余额
func (r *propsRedisRepo) GetWallet(ctx *gin.Context, uId, cId uint64) (float64, error) {
	if r.redisClient == nil || cId <= 0 || uId <= 0 {
		return 0, utils.ErrParameter
	}

	walletStr, err := r.redisClient.
		HGet(
			ctx,
			fmt.Sprintf("%s:%d:%d", redis_entity.RedisProps, uId, cId),
			redis_entity.RedisPropsWallet,
		).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(walletStr, 64)
}

// 获取访问权限
func (r *propsRedisRepo) GetAL(ctx *gin.Context, uId, cId uint64) (uint64, error) {
	if r.redisClient == nil || cId <= 0 || uId <= 0 {
		return 0, utils.ErrParameter
	}

	alStr, err := r.redisClient.
		HGet(
			ctx,
			fmt.Sprintf("%s:%d:%d", redis_entity.RedisProps, uId, cId),
			redis_entity.RedisPropsAL,
		).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(alStr, 10, 32)
}

// 获取主要返还玩家比率
func (r *propsRedisRepo) GetMRTP(ctx *gin.Context, uId, cId uint64) (float64, error) {
	if r.redisClient == nil || cId <= 0 || uId <= 0 {
		return 0, utils.ErrParameter
	}

	mrtpStr, err := r.redisClient.
		HGet(
			ctx,
			fmt.Sprintf("%s:%d:%d", redis_entity.RedisProps, uId, cId),
			redis_entity.RedisPropsMRTP,
		).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(mrtpStr, 64)
}
