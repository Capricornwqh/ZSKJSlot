package pgsql_entity

import "time"

// 俱乐部数据
type Club struct {
	CId       uint64    `gorm:"primaryKey;autoIncrement;comment:俱乐部ID;" json:"cid" yaml:"cid"`
	CreatedAt time.Time `gorm:"column:createdAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:创建时间;" redis:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:修改时间;" redis:"updatedAt" json:"updatedAt"`
	Name      string    `gorm:"column:name;type:varchar(150);not null;default:'';comment:俱乐部名称;" json:"name,omitempty" yaml:"name,omitempty"`
	Bank      float64   `gorm:"column:bank;type:numeric(19,4);not null;default:0;comment:用户输赢余额，以金币为单位;" json:"bank" yaml:"bank"`
	Fund      float64   `gorm:"column:fund;type:numeric(19,4);not null;default:0;comment:奖池资金，以金币为单位;" json:"fund" yaml:"fund"`
	Lock      float64   `gorm:"column:lock;type:numeric(19,4);not null;default:0;comment:游戏中不变的存款;" json:"lock" yaml:"lock"`
	Rate      float64   `gorm:"column:rate;type:numeric(19,4);not null;default:2.5;comment:累积奖池游戏的奖池率;" json:"rate" yaml:"rate"`
	MRTP      float64   `gorm:"column:mrtp;type:numeric(19,4);not null;default:0;comment:主要返还玩家比率(Master RTP);" json:"mrtp" yaml:"mrtp"`
}

// 表名
func (Club) TableName() string {
	return "club"
}

// 表注释
func (Club) Comment() string {
	return "俱乐部表"
}
