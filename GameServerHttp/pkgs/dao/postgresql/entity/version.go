package pgsql_entity

// 版本表
type Version struct {
	Id            int64 `gorm:"column:id;type:bigserial;primaryKey;not null;comment:主键ID;" redis:"id" json:"id"`
	VersionNumber int64 `gorm:"column:versionNumber;type:bigint;not null;default:0;comment:版本号;" redis:"versionNumber" json:"versionNumber"`
}

// 表名
func (Version) TableName() string {
	return "version"
}

// 表注释
func (Version) Comment() string {
	return "版本表"
}
