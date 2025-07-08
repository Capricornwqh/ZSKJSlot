package model_user

import (
	"time"
)

// 账号类型
type AccountType int32

const (
	ACCOUNTTYPE_NOACCOUNT AccountType = 0
	ACCOUNTTYPE_ACCOUNT   AccountType = 1 //账号
	ACCOUNTTYPE_PHONE     AccountType = 2 //手机
	ACCOUNTTYPE_EMAIL     AccountType = 3 //邮箱
)

type UserVerifyRequest struct {
	Language string      `json:"language,omitempty"`                   //语言 zh_CN, en_US等
	Typed    AccountType `json:"typed" binding:"required,oneof=1 2 3"` //类型 2-手机 3-邮箱
	Object   string      `json:"object" binding:"required,min=1"`      //手机号或邮箱
}

// 用户请求
type UserSignRequest struct {
	Typed    AccountType `json:"typed" binding:"required,oneof=1 2 3"` //类型 1-账号登录 2-手机 3-邮箱
	Object   string      `json:"object" binding:"required,min=1"`      //手机号或邮箱或id
	Passwd   string      `json:"passwd" binding:"required,min=1"`      //密码
	Username string      `json:"username"`                             //用户名
	Token    string      `json:"token"`                                //token
	Device   string      `json:"device"`                               //设备
	Terminal string      `json:"terminal"`                             //终端
	Country  string      `json:"country"`                              //国家
	Language string      `json:"language"`                             //语言
}

// 用户基本信息
type UserBaseInfo struct {
	UserId        uint64    `json:"userId,omitempty"`        //主键ID
	CreatedAt     time.Time `json:"createdAt,omitempty"`     //创建时间
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`     //修改时间
	SuspendedAt   time.Time `json:"suspendedAt,omitempty"`   //冻住时间
	DeletedAt     time.Time `json:"deletedAt,omitempty"`     //删除时间
	LastLoginDate time.Time `json:"lastLoginDate,omitempty"` //最后登录时间
	Username      string    `json:"username,omitempty"`      //用户名
	Pass          string    `json:"pass,omitempty"`          //密码
	EMail         string    `json:"eMail,omitempty"`         //邮箱
	Status        uint32    `json:"status,omitempty"`        //用户状态(1-已激活, 2-已签约, 3-冻结, 10-已删除)
	Gender        int32     `json:"gender,omitempty"`        //性别 0-保密 1-男 2-女
	Birthday      time.Time `json:"birthday,omitempty"`      //生日
	DisplayName   string    `json:"displayName,omitempty"`   //显示名
	Avatar        string    `json:"avatar,omitempty"`        //头像
	Mobile        string    `json:"mobile,omitempty"`        //手机
	IPInfo        string    `json:"iPInfo,omitempty"`        //IP信息
	IsAdmin       bool      `json:"isAdmin,omitempty"`       //是否管理员(0-非管理员, 1-管理员)
	Country       string    `json:"country,omitempty"`       //国家
	Language      string    `json:"language,omitempty"`      //语言
	Token         string    `json:"token,omitempty"`         //token
}

// 用户profile请求
type UserProfile struct {
	DisplayName string     `json:"displayName,omitempty"` //显示名
	Birthday    *time.Time `json:"birthday,omitempty"`    //生日
	Country     string     `json:"country,omitempty"`     //国家
	Language    string     `json:"language,omitempty"`    //语言
	Avatar      string     `json:"avatar,omitempty"`      //头像
	Gender      int32      `json:"gender,omitempty"`      //性别 0-保密 1-男 2-女
	NoticeState int32      `json:"noticeState,omitempty"` //通知状态(1-开启, 2-关闭)
}
