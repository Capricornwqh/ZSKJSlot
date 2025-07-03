package pgsql_entity

import (
	"time"
)

// 操作类型枚举
type OperationType string

const (
	OpTypeGamePlay     OperationType = "play"     // 游戏玩法记录
	OpTypeBet          OperationType = "bet"      // 投注记录
	OpTypeWin          OperationType = "win"      // 中奖记录
	OpTypeShopPurchase OperationType = "purchase" // 商店购买记录
	OpTypeBonus        OperationType = "bonus"    // 奖励记录
	OpTypeRecharge     OperationType = "recharge" // 充值记录
)

// 记录表
type Record struct {
	Id              int64         `gorm:"column:id;type:bigserial;primaryKey;not null;comment:主键ID;" redis:"id" json:"id"`
	UId             int64         `gorm:"column:uId;type:bigint;not null;index;comment:玩家ID;" redis:"uId" json:"uId"`
	GId             int64         `gorm:"column:gId;type:bigint;index;comment:游戏ID;" redis:"gId" json:"gId"`
	OperationType   OperationType `gorm:"column:operationType;type:varchar(50);not null;index;comment:操作类型;" redis:"operationType" json:"operationType"`
	Details         JSONB         `gorm:"column:details;type:jsonb;comment:操作详情(JSON格式);" redis:"details" json:"details"`
	CurrencyChange  float64       `gorm:"column:currencyChange;type:decimal(5,2);comment:货币变化(正数为增加，负数为减少);" redis:"currencyChange" json:"currencyChange"`
	CurrencyBalance float64       `gorm:"column:currencyBalance;type:decimal(5,2);comment:操作后的货币余额;" redis:"currencyBalance" json:"currencyBalance"`
	SessionId       string        `gorm:"column:sessionId;type:varchar(100);index;comment:会话ID,关联同一游戏会话的多个操作;" redis:"sessionId" json:"sessionId"`
	DeviceInfo      string        `gorm:"column:deviceInfo;type:varchar(255);comment:设备信息;" redis:"deviceInfo" json:"deviceInfo"`
	IPAddress       string        `gorm:"column:ipAddress;type:varchar(50);comment:IP地址;" redis:"ipAddress" json:"ipAddress"`
	CreatedAt       time.Time     `gorm:"column:createdAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:创建时间;" redis:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time     `gorm:"column:updatedAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:修改时间;" redis:"updatedAt" json:"updatedAt"`
}

// 表名
func (Record) TableName() string {
	return "record"
}

// 表注释
func (Record) Comment() string {
	return "玩家操作记录表"
}
