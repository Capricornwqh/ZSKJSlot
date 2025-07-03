package controller

import (
	"SlotGameServer/pkgs/model"
	"SlotGameServer/pkgs/service"
	"SlotGameServer/utils"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	gameService *service.GameService
}

func NewGameController(service *service.GameService) *GameController {
	return &GameController{
		gameService: service,
	}
}

// 创建新游戏
func (c *GameController) GameNew(ctx *gin.Context) {
	req := &model.GameArg{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

}

// 加入游戏
func (c *GameController) GameJoin(ctx *gin.Context) {
}

// 获取游戏信息
func (c *GameController) GameInfo(ctx *gin.Context) {
}

// 获取RTP信息
func (c *GameController) GameRtpGet(ctx *gin.Context) {
}

// 获取slot投注
func (c *GameController) SlotBetGet(ctx *gin.Context) {
}

// 设置slot投注
func (c *GameController) SlotBetSet(ctx *gin.Context) {
}

// 获取slot选择
func (c *GameController) SlotSelGet(ctx *gin.Context) {
}

// 设置slot选择
func (c *GameController) SlotSelSet(ctx *gin.Context) {
}

// 设置slot模式
func (c *GameController) SlotModeSet(ctx *gin.Context) {
}

// 执行slot旋转
func (c *GameController) SlotSpin(ctx *gin.Context) {
}

// 执行slot双倍
func (c *GameController) SlotDoubleup(ctx *gin.Context) {
}

// 执行slot收集
func (c *GameController) SlotCollect(ctx *gin.Context) {
}
