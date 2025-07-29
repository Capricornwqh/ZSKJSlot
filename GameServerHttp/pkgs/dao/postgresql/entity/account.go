package entity_pgsql

import (
	"strconv"
	"time"
)

const (
	EmailStateAvailable    = 1
	EmailStateToBeVerified = 2
)

// 用户状态
type UF uint

const (
	UFactivated UF = 1 << iota // 用户已激活
	UFsigncode                 // 用户已签约
	UFsuspended                // 用户已冻结
	UFdeleted                  // 用户已删除
)

func (u UF) MarshalBinary() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(u), 10)), nil
}

func (u UF) MarshalText() ([]byte, error) {
	return u.MarshalBinary()
}

// 用户访问权限
type AL uint

const (
	ALmember AL = 1 << iota // 用户有权访问俱乐部
	ALdealer                // 可以更改俱乐部游戏设置和用户游戏玩法
	ALbooker                // 可以更改用户属性并将用户资金转入/转出俱乐部存款
	ALmaster                // 可以更改俱乐部银行账户、资金账户和存款账户
	ALadmin                 // 可以更改其他用户的相同访问权限
	ALall    = ALmember | ALdealer | ALbooker | ALmaster | ALadmin
)

func (a AL) MarshalBinary() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(a), 10)), nil
}

func (a AL) MarshalText() ([]byte, error) {
	return a.MarshalBinary()
}

// 用户表
type Account struct {
	Id            uint64    `gorm:"column:id;type:bigserial;primaryKey;not null;comment:主键ID;" redis:"id" json:"id"`
	PlatformId    string    `gorm:"column:platformId;type:varchar(100);not null;default:'';comment:平台ID;" redis:"platformId" json:"platformId"`
	PlatformType  string    `gorm:"column:platformType;type:varchar(50);not null;default:'';comment:平台类型;" redis:"platformType" json:"platformType"`
	CreatedAt     time.Time `gorm:"column:createdAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:创建时间;" redis:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updatedAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:修改时间;" redis:"updatedAt" json:"updatedAt"`
	DeletedAt     time.Time `gorm:"column:deletedAt;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:删除时间;" redis:"deletedAt" json:"deletedAt"`
	LastLoginDate time.Time `gorm:"column:lastLoginDate;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:最后登录时间;" redis:"lastLoginDate" json:"lastLoginDate"`
	Username      string    `gorm:"column:username;type:varchar(50);not null;default:'';comment:用户名;" redis:"username" json:"username"`
	Pass          string    `gorm:"column:pass;type:varchar(255);not null;default:'';comment:密码;" redis:"pass" json:"pass"`
	Email         string    `gorm:"column:email;type:varchar(100);not null;default:'';comment:邮箱;" redis:"email" json:"email"`
	Rank          int32     `gorm:"column:rank;type:int;not null;default:0;comment:排名;" redis:"rank" json:"rank"`
	Gender        int32     `gorm:"column:gender;type:smallint;not null;default:0;comment:性别 0-保密 1-男 2-女;" redis:"gender" json:"gender"`
	IsAdmin       int32     `gorm:"column:isAdmin;type:smallint;not null;default:0;comment:是否管理员(0-非管理员, 1-管理员);" redis:"isAdmin" json:"isAdmin"`
	Birthday      time.Time `gorm:"column:birthday;type:timestamptz;not null;default:CURRENT_TIMESTAMP;comment:生日;" redis:"birthday" json:"birthday"`
	Avatar        string    `gorm:"column:avatar;type:varchar(1024);not null;default:'';comment:头像;" redis:"avatar" json:"avatar"`
	Mobile        string    `gorm:"column:mobile;type:varchar(20);not null;default:'';comment:手机;" redis:"mobile" json:"mobile"`
	Address       string    `gorm:"column:address;type:varchar(255);not null;default:'';comment:地址;" redis:"address" json:"address"`
	IPInfo        string    `gorm:"column:ipInfo;type:varchar(255);not null;default:'';comment:IP信息;" redis:"ipInfo" json:"ipInfo"`
	Country       string    `gorm:"column:country;type:varchar(20);not null;default:'CN';comment:国家;" redis:"country" json:"country"`
	Language      string    `gorm:"column:language;type:varchar(20);not null;default:'zh';comment:语言;" redis:"language" json:"language"`
	Device        string    `gorm:"column:device;type:varchar(255);not null;default:'';comment:设备;" redis:"device" json:"device"`
	Terminal      string    `gorm:"column:terminal;type:varchar(255);not null;default:'';comment:终端;" redis:"terminal" json:"terminal"`
	Status        UF        `gorm:"column:status;type:smallint;not null;default:1;comment:用户状态(1-已激活, 2-已签约, 3-冻结, 10-已删除);" redis:"status" json:"status"`
	GAL           AL        `gorm:"column:gal;type:smallint;not null;default:1;comment:用户访问权限(1-会员, 2-经销商, 4-簿记员, 8-大师, 16-管理员);" redis:"gal" json:"gal"`
}

// 表名
func (Account) TableName() string {
	return "account"
}

// 表注释
func (Account) Comment() string {
	return "账号表"
}

// 设置accountId的起始值
func (Account) SetAccountIdStartValue() string {
	return "ALTER SEQUENCE account_id_seq RESTART WITH 10000000;"
}
