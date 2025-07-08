package entity_pgsql

import "time"

// 钱包日志，记录用户钱包余额变动
type Walletlog struct {
	Id        uint64    `gorm:"column:id;type:bigserial;primaryKey;not null;comment:主键ID;" redis:"id" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:创建时间;" redis:"createdAt" json:"createdAt"`
	CId       uint64    `gorm:"column:cid;type:bigint;not null;default:0;index:idx_cid_uid,priority:1;comment:俱乐部ID;" json:"cid" redis:"cid"`
	UId       uint64    `gorm:"column:uid;type:bigint;not null;default:0;index:idx_cid_uid,priority:2;comment:用户ID;" json:"uid" redis:"uid"`
	AId       uint64    `gorm:"column:aid;type:bigint;not null;default:0;comment:管理员ID;" json:"aid" redis:"aid"`
	Wallet    float64   `gorm:"column:wallet;type:numeric(19,4);not null;default:0;comment:新的钱包余额(以金币为单位);" json:"wallet" redis:"wallet"`
	Sum       float64   `gorm:"column:sum;type:numeric(19,4);not null;default:0;comment:变动金额;" json:"sum" redis:"sum"`
}

// 表名
func (Walletlog) TableName() string {
	return "wallet_log"
}

// 表注释
func (Walletlog) Comment() string {
	return "钱包日志表"
}
