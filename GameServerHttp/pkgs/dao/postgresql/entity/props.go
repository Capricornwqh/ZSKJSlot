package entity_pgsql

import "time"

// 包含用户在特定俱乐部的属性。
// 任何属性默认可以为零，或者在数据库中不存在该对象。
type Props struct {
	CId       uint64    `gorm:"column:cid;type:bigint;not null;index:idx_cid_uid,priority:1;comment:俱乐部ID;" json:"cid" redis:"cid"`
	UId       uint64    `gorm:"column:uid;type:bigint;not null;index:idx_cid_uid,priority:2;comment:用户ID;" json:"uid" redis:"uid"`
	CreatedAt time.Time `gorm:"column:createdAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:创建时间;" redis:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:修改时间;" redis:"updatedAt" json:"updatedAt"`
	Wallet    float64   `gorm:"column:wallet;type:numeric(19,4);not null;default:0;comment:钱包余额(以金币为单位);" json:"wallet" redis:"wallet"`
	Access    AL        `gorm:"column:access;type:smallint;not null;default:1;comment:用户访问权限(1-会员, 2-经销商, 4-簿记员, 8-大师, 16-管理员);" redis:"access" json:"access"`
	MRTP      float64   `gorm:"column:mrtp;type:numeric(19,4);not null;default:0;comment:主要返还玩家比率;" redis:"mrtp" json:"mrtp"`
}

// 表名
func (Props) TableName() string {
	return "props"
}

// 表注释
func (Props) Comment() string {
	return "属性表"
}
