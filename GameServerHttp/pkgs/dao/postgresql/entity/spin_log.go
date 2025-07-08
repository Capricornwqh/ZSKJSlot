package entity_pgsql

import "time"

// 旋转日志记录，记录每次旋转的详细信息
type Spinlog struct {
	Id        uint64    `gorm:"column:id;type:bigserial;primaryKey;not null;comment:主键ID;" redis:"id" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:创建时间;" redis:"createdAt" json:"createdAt"`
	GId       uint64    `gorm:"column:gid;type:bigint;not null;default:0;comment:游戏ID" json:"gid" redis:"gid"`
	MRTP      float64   `gorm:"column:mrtp;type:numeric(19,4);not null;default:0;comment:主要返还玩家比率" json:"mrtp" redis:"mrtp"`
	Game      string    `gorm:"column:game;type:text;not null;comment:游戏数据" json:"game" redis:"game"`
	Wins      string    `gorm:"column:wins;type:text;comment:JSON序列化的获胜列表" json:"wins,omitempty" redis:"wins"`
	Gain      float64   `gorm:"column:gain;type:numeric(19,4);not null;default:0;comment:最后旋转的总收益" json:"gain" redis:"gain"`
	Wallet    float64   `gorm:"column:wallet;type:numeric(19,4);not null;default:0;comment:钱包余额" json:"wallet" redis:"wallet"`
}

// 表名
func (Spinlog) TableName() string {
	return "spin_log"
}

// 表注释
func (Spinlog) Comment() string {
	return "旋转日志表"
}
