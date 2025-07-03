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

type clubRedisRepo struct {
	redisClient *redis.Client
}

type ClubRedisRepo interface {
	// 保存俱乐部
	SetClub(ctx *gin.Context, club *pgsql_entity.Club) error
	// 获取俱乐部
	GetClub(ctx *gin.Context, cid uint64) (*pgsql_entity.Club, error)
	// 获取Name
	GetClubName(ctx *gin.Context, cid uint64) (string, error)
	// 获取Bank
	GetClubBank(ctx *gin.Context, cid uint64) (float64, error)
	// 计算Bank
	IncrClubBank(ctx *gin.Context, cid uint64, amount float64) error
	// 获取Fund
	GetClubFund(ctx *gin.Context, cid uint64) (float64, error)
	// 计算Fund
	IncrClubFund(ctx *gin.Context, cid uint64, amount float64) error
	// 获取Lock
	GetClubLock(ctx *gin.Context, cid uint64) (float64, error)
	// 计算Lock
	IncrClubLock(ctx *gin.Context, cid uint64, amount float64) error
	// 获取俱乐部所有资金信息
	GetCash(ctx *gin.Context, cid uint64) (float64, float64, float64, error)
	// 计算cash
	IncrClubCash(ctx *gin.Context, cid uint64, bank, fund, lock float64) error
	// 获取Rate
	GetClubRate(ctx *gin.Context, cid uint64) (float64, error)
	// 获取MRTP
	GetClubMRTP(ctx *gin.Context, cid uint64) (float64, error)
}

func NewClubRedisRepo(redisClient *redis.Client) ClubRedisRepo {
	return &clubRedisRepo{
		redisClient: redisClient,
	}
}

// 保存俱乐部
func (r *clubRedisRepo) SetClub(ctx *gin.Context, club *pgsql_entity.Club) error {
	if r.redisClient == nil || club == nil || club.CId <= 0 {
		return utils.ErrParameter
	}

	err := r.redisClient.
		HSet(ctx, fmt.Sprintf("%s:%d", redis_entity.RedisClub, club.CId), club).
		Err()
	if err != nil {
		return err
	}
	return nil
}

// 获取俱乐部
func (r *clubRedisRepo) GetClub(ctx *gin.Context, cid uint64) (*pgsql_entity.Club, error) {
	if r.redisClient == nil || cid <= 0 {
		return nil, utils.ErrParameter
	}

	tmpClub := &pgsql_entity.Club{}
	err := r.redisClient.
		HGetAll(ctx, fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid)).
		Scan(tmpClub)
	if err != nil {
		return nil, err
	}

	if tmpClub.CId == 0 {
		return nil, utils.ErrDataNotFound
	}
	return tmpClub, nil
}

// 获取Name
func (r *clubRedisRepo) GetClubName(ctx *gin.Context, cid uint64) (string, error) {
	if r.redisClient == nil || cid <= 0 {
		return "", utils.ErrParameter
	}

	name, err := r.redisClient.
		HGet(
			ctx,
			fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid),
			redis_entity.RedisClubName,
		).Result()
	if err != nil {
		return "", err
	}
	return name, nil
}

// 获取Bank
func (r *clubRedisRepo) GetClubBank(ctx *gin.Context, cid uint64) (float64, error) {
	if r.redisClient == nil || cid <= 0 {
		return 0, utils.ErrParameter
	}

	bank, err := r.redisClient.
		HGet(
			ctx,
			fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid),
			redis_entity.RedisClubBank,
		).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(bank, 64)
}

// 计算Bank
func (r *clubRedisRepo) IncrClubBank(ctx *gin.Context, cid uint64, amount float64) error {
	if r.redisClient == nil || cid <= 0 || amount <= 0 {
		return utils.ErrParameter
	}

	_, err := r.redisClient.
		HIncrByFloat(
			ctx,
			fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid),
			redis_entity.RedisClubBank,
			amount,
		).Result()
	if err != nil {
		return err
	}

	return nil
}

// 获取Fund
func (r *clubRedisRepo) GetClubFund(ctx *gin.Context, cid uint64) (float64, error) {
	if r.redisClient == nil || cid <= 0 {
		return 0, utils.ErrParameter
	}

	fund, err := r.redisClient.
		HGet(
			ctx,
			fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid),
			redis_entity.RedisClubFund,
		).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(fund, 64)
}

// 计算Fund
func (r *clubRedisRepo) IncrClubFund(ctx *gin.Context, cid uint64, amount float64) error {
	if r.redisClient == nil || cid <= 0 || amount <= 0 {
		return utils.ErrParameter
	}

	_, err := r.redisClient.
		HIncrByFloat(
			ctx,
			fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid),
			redis_entity.RedisClubFund,
			amount,
		).Result()
	if err != nil {
		return err
	}

	return nil
}

// 获取Lock
func (r *clubRedisRepo) GetClubLock(ctx *gin.Context, cid uint64) (float64, error) {
	if r.redisClient == nil || cid <= 0 {
		return 0, utils.ErrParameter
	}

	lock, err := r.redisClient.
		HGet(
			ctx,
			fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid),
			redis_entity.RedisClubLock,
		).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(lock, 64)
}

// 计算Lock
func (r *clubRedisRepo) IncrClubLock(ctx *gin.Context, cid uint64, amount float64) error {
	if r.redisClient == nil || cid <= 0 || amount <= 0 {
		return utils.ErrParameter
	}

	_, err := r.redisClient.
		HIncrByFloat(
			ctx,
			fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid),
			redis_entity.RedisClubLock,
			amount,
		).Result()
	if err != nil {
		return err
	}

	return nil
}

// 获取俱乐部所有资金信息
func (r *clubRedisRepo) GetCash(ctx *gin.Context, cid uint64) (float64, float64, float64, error) {
	if r.redisClient == nil || cid <= 0 {
		return 0, 0, 0, utils.ErrParameter
	}

	values, err := r.redisClient.HMGet(
		ctx,
		fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid),
		redis_entity.RedisClubBank,
		redis_entity.RedisClubFund,
		redis_entity.RedisClubLock).Result()

	if err != nil {
		return 0, 0, 0, err
	}

	var bank, fund, lock float64
	var convErr error

	if len(values) == 3 {
		// 处理 bank
		if values[0] != nil {
			bank, convErr = strconv.ParseFloat(values[0].(string), 64)
			if convErr != nil {
				return 0, 0, 0, convErr
			}
		}

		// 处理 fund
		if values[1] != nil {
			fund, convErr = strconv.ParseFloat(values[1].(string), 64)
			if convErr != nil {
				return 0, 0, 0, convErr
			}
		}

		// 处理 lock
		if values[2] != nil {
			lock, convErr = strconv.ParseFloat(values[2].(string), 64)
			if convErr != nil {
				return 0, 0, 0, convErr
			}
		}
	} else {
		return 0, 0, 0, utils.ErrDataNotFound
	}

	return bank, fund, lock, nil
}

// 计算cash
func (r *clubRedisRepo) IncrClubCash(ctx *gin.Context, cid uint64, bank, fund, lock float64) error {
	if r.redisClient == nil || cid <= 0 || (bank <= 0 && fund <= 0 && lock <= 0) {
		return utils.ErrParameter
	}

	err := r.IncrClubBank(ctx, cid, bank)
	if err != nil {
		return err
	}

	err = r.IncrClubFund(ctx, cid, fund)
	if err != nil {
		return err
	}

	err = r.IncrClubLock(ctx, cid, lock)
	if err != nil {
		return err
	}
	return nil
}

// 获取Rate
func (r *clubRedisRepo) GetClubRate(ctx *gin.Context, cid uint64) (float64, error) {
	if r.redisClient == nil || cid <= 0 {
		return 0, utils.ErrParameter
	}

	rate, err := r.redisClient.
		HGet(
			ctx,
			fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid),
			redis_entity.RedisClubRate,
		).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(rate, 64)
}

// 获取MRTP
func (r *clubRedisRepo) GetClubMRTP(ctx *gin.Context, cid uint64) (float64, error) {
	if r.redisClient == nil || cid <= 0 {
		return 0, utils.ErrParameter
	}

	mrtp, err := r.redisClient.
		HGet(
			ctx,
			fmt.Sprintf("%s:%d", redis_entity.RedisClub, cid),
			redis_entity.RedisClubMRTP,
		).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(mrtp, 64)
}
