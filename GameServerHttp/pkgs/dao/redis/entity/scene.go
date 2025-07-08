package entity_redis

import entity_pgsql "SlotGameServer/pkgs/dao/postgresql/entity"

const (
	RedisScene = "scene" // 场景信息
)

// 表示游戏及其所有相关环境
type Scene struct {
	SId                uint64     `json:"sid"`  // 最新旋转ID
	Game               any        `json:"game"` // 游戏实例
	entity_pgsql.Story `json:"-"` // 嵌入游戏故事信息
}
