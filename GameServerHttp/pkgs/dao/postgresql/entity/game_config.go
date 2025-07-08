package entity_pgsql

import "time"

// 游戏配置表
type GameConfig struct {
	Id              int64     `gorm:"column:id;type:bigserial;primaryKey;not null;comment:主键ID;" json:"id"`
	GameName        string    `gorm:"column:gameName;type:varchar(100);not null;comment:游戏名称;" json:"gameName"`
	GameDescription string    `gorm:"column:gameDescription;type:text;comment:游戏描述;" json:"gameDescription"`
	GameType        string    `gorm:"column:gameType;type:varchar(50);not null;index;comment:游戏类型(slot/table/etc);" json:"gameType"`
	GameCategory    string    `gorm:"column:gameCategory;type:varchar(50);index;comment:游戏分类;" json:"gameCategory"`
	GameThumbnail   string    `gorm:"column:gameThumbnail;type:varchar(255);comment:游戏缩略图URL;" json:"gameThumbnail"`
	GameFeatures    JSONB     `gorm:"column:gameFeatures;type:jsonb;comment:游戏特性(JSON格式);" json:"gameFeatures"`
	GameSettings    JSONB     `gorm:"column:gameSettings;type:jsonb;comment:游戏设置(JSON格式);" json:"gameSettings"`
	MinBet          int64     `gorm:"column:minBet;type:bigint;not null;default:1;comment:最小下注额;" json:"minBet"`
	MaxBet          int64     `gorm:"column:maxBet;type:bigint;not null;comment:最大下注额;" json:"maxBet"`
	RTP             float64   `gorm:"column:rtp;type:decimal(5,2);comment:理论返还率(%);" json:"rtp"`
	DisplayOrder    int32     `gorm:"column:displayOrder;type:int;not null;default:0;comment:显示顺序;" json:"displayOrder"`
	IsActive        bool      `gorm:"column:isActive;type:boolean;not null;default:true;index;comment:是否激活;" json:"isActive"`
	IsHot           bool      `gorm:"column:isHot;type:boolean;not null;default:false;comment:是否热门游戏;" json:"isHot"`
	IsNew           bool      `gorm:"column:isNew;type:boolean;not null;default:false;comment:是否新游戏;" json:"isNew"`
	CreatedAt       time.Time `gorm:"column:createdAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:创建时间;" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"column:updatedAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:修改时间;" json:"updatedAt"`
}

// 表名
func (GameConfig) TableName() string {
	return "game_config"
}

// 表注释
func (GameConfig) Comment() string {
	return "游戏配置表"
}
