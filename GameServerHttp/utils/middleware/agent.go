package utils_middleware

import (
	"SlotGameServer/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AgentMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tmpAppAgent := ctx.GetHeader("client-agent")
		if len(tmpAppAgent) <= 0 {
			logrus.WithContext(ctx).Error("client-agent is missing")
			utils.HandleError(ctx, utils.ErrParameter)
			return
		}

		if tmpAppAgent != utils.MD_AGENT_APP && tmpAppAgent != utils.MD_AGENT_WEB {
			logrus.WithContext(ctx).Error("client-agent is error")
			utils.HandleError(ctx, utils.ErrParameter)
			return
		}

		ctx.Next()
	}
}
