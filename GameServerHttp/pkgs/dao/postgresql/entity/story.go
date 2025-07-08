package entity_pgsql

import "time"

// Story 代表用户在俱乐部中打开的游戏实例。
// 每个游戏实例拥有自己的GID。Alias是游戏类型标识符。
type Story struct {
	GId       uint64    `gorm:"column:gid;type:bigserial;primaryKey;not null;comment:主键ID;" redis:"gid" json:"gid"`
	CreatedAt time.Time `gorm:"column:createdAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:创建时间;" redis:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:修改时间;" redis:"updatedAt" json:"updatedAt"`
	Alias     string    `gorm:"column:alias;type:varchar(150);not null;default:'';comment:游戏类型标识符;" json:"alias" redis:"alias"`
	CId       uint64    `gorm:"column:cid;type:bigint;not null;default:0;index:idx_cid_uid,priority:1;comment:俱乐部ID;" json:"cid" redis:"cid"`
	UId       uint64    `gorm:"column:uid;type:bigint;not null;default:0;index:idx_cid_uid,priority:2;comment:用户ID;" json:"uid" redis:"uid"`
}

// 表名
func (Story) TableName() string {
	return "story"
}

// 表注释
func (Story) Comment() string {
	return "游戏实例表"
}
