package entity_pgsql

import "time"

// 银行日志，记录俱乐部资金变动
type Banklog struct {
	Id        uint64    `gorm:"column:id;type:bigserial;primaryKey;not null;comment:主键ID;" redis:"id" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:创建时间;" redis:"createdAt" json:"createdAt"`
	Bank      float64   `gorm:"column:bank;type:numeric(19,4);not null;default:0;comment:银行余额" redis:"bank" json:"bank"`
	Fund      float64   `gorm:"column:fund;type:numeric(19,4);not null;default:0;comment:奖池余额" redis:"fund" json:"fund"`
	Lock      float64   `gorm:"column:lock;type:numeric(19,4);not null;default:0;comment:锁定资金;" redis:"lock" json:"lock"`
	BankSum   float64   `gorm:"column:banksum;type:numeric(19,4);not null;default:0;comment:银行变动金额" redis:"banksum" json:"banksum"`
	FundSum   float64   `gorm:"column:fundsum;type:numeric(19,4);not null;default:0;comment:奖池变动金额" redis:"fundsum" json:"fundsum"`
	LockSum   float64   `gorm:"column:locksum;type:numeric(19,4);not null;default:0;comment:锁定资金变动金额;" redis:"locksum" json:"locksum"`
}

// 表名
func (Banklog) TableName() string {
	return "bank_log"
}

// 表注释
func (Banklog) Comment() string {
	return "银行日志表"
}
