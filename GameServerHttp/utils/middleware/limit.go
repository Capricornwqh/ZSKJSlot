package utils_middleware

import (
	entity_redis "SlotGameServer/pkgs/dao/redis/entity"
	repo_redis "SlotGameServer/pkgs/dao/redis/repo"
	"SlotGameServer/utils"
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 重复提交
func RepeatedMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tmpUserId := ctx.GetInt64("userId")
		if tmpUserId <= 0 {
			logrus.WithContext(ctx).Error("userId is missing")
			utils.HandleError(ctx, utils.ErrParameter)
			return
		}

		md5Id := ""
		hMD5 := md5.New()
		_, err := io.WriteString(hMD5, fmt.Sprintf("%s:%d%s", entity_redis.RedisMethodMD5, tmpUserId, ctx.FullPath()))
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			utils.HandleError(ctx, utils.ErrParse)
			return
		}
		md5Id = fmt.Sprintf("%x", hMD5.Sum(nil))
		if utils.RedisClient != nil {
			value, err := repo_redis.GetMethodMD5(ctx, md5Id)
			if err != nil && err != utils.Nil {
				logrus.WithContext(ctx).Error(err)
				utils.HandleError(ctx, utils.ErrUserNotFound)
				return
			}
			if value == "1" {
				logrus.WithContext(ctx).Error(utils.ErrRepeated)
				utils.HandleError(ctx, utils.ErrRepeated)
				return
			}
			err = repo_redis.SetMethodMD5(ctx, md5Id, "1", entity_redis.TimeMethodMD5*time.Second)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				utils.HandleError(ctx, utils.ErrOperation)
				return
			}
		}

		ctx.Next()

		if md5Id != "" {
			repo_redis.DelMethodMD5(ctx, md5Id)
		}
	}
}
