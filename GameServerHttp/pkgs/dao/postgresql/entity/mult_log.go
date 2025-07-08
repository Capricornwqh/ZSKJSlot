package entity_pgsql

import "time"

// 倍数日志记录，记录每次倍数游戏的详细信息
type Multlog struct {
	Id        uint64    `gorm:"column:id;type:bigserial;primaryKey;not null;comment:主键ID;" redis:"id" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:创建时间;" redis:"createdAt" json:"createdAt"`
	GId       uint64    `gorm:"column:gid;type:bigint;not null;comment:游戏ID;" redis:"gid" json:"gid"`
	MRTP      float64   `gorm:"column:mrtp;type:numeric(19,4);not null;default:0;comment:主要返还玩家比率;" redis:"mrtp" json:"mrtp"`
	Mult      int       `gorm:"column:mult;type:integer;not null;default:0;comment:倍数;" redis:"mult" json:"mult"`
	Risk      float64   `gorm:"column:risk;type:numeric(19,4);not null;default:0;comment:风险值;" redis:"risk" json:"risk"`
	Gain      float64   `gorm:"column:gain;type:numeric(19,4);not null;default:0;comment:总收益;" redis:"gain" json:"gain"`
	Wallet    float64   `gorm:"column:wallet;type:numeric(19,4);not null;default:0;comment:钱包余额;" redis:"wallet" json:"wallet"`
}

// 表名
func (Multlog) TableName() string {
	return "mult_log"
}

// 表注释
func (Multlog) Comment() string {
	return "倍数日志表"
}
