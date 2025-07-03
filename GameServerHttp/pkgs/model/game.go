package model

type GameArg struct {
	CID   uint64 `json:"cid" yaml:"cid" form:"cid" binding:"required"`
	UID   uint64 `json:"uid" yaml:"uid" form:"uid" binding:"required"`
	Alias string `json:"alias" yaml:"alias" form:"alias" binding:"required"`
}

type GameRet struct {
	GID    uint64  `json:"gid" yaml:"gid"`
	Game   any     `json:"game" yaml:"game"`
	Wallet float64 `json:"wallet" yaml:"wallet"`
}
