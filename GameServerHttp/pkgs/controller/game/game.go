package controller_game

import (
	service "SlotGameServer/pkgs/service/game"

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

// 游戏列表
func (c *GameController) GameList(ctx *gin.Context) {
}

// 获取游戏信息
func (c *GameController) GameInfo(ctx *gin.Context) {
}
