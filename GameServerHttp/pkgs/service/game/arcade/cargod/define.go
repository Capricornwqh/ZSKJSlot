package service_cargod

// 事件ID常量定义
const (
	EventFortune     = 0  // 福气奖
	EventDouble      = 1  // 加倍奖
	EventChainLight1 = 2  // 连灯奖1
	EventChainLight2 = 3  // 连灯奖2
	EventChainLight3 = 4  // 连灯奖3
	EventExtraBonus1 = 5  // 锦上添花1
	EventExtraBonus2 = 6  // 锦上添花2
	EventExtraBonus3 = 7  // 锦上添花3
	EventWorldTravel = 8  // 云游四海
	EventRedCars     = 9  // 四大名车-红
	EventGreenCars   = 10 // 四大名车-绿
	EventYellowCars  = 11 // 四大名车-黄
	EventFerrari     = 12 // 品牌奖-法拉利
	EventBenz        = 13 // 品牌奖-奔驰
	EventBMW         = 14 // 品牌奖-宝马
	EventAudi        = 15 // 品牌奖-奥迪
	EventRedLight    = 16 // 红灯奖
	EventGreenLight  = 17 // 绿灯奖
	EventYellowLight = 18 // 黄灯奖
	EventGrandSlam   = 20 // 满贯列车
	EventJackpot     = 21 // Jackpot
)

// Car 车神游戏核心结构体
type Car struct {
	GroupSet *GroupSet   `json:"groupSet"` // 组合设置：定义车辆位置、事件和奖池配置
	LevelSet []*LevelSet `json:"levelSet"` // 等级设置：定义赔率、位置权重和特殊率
	BetTable []int64     `json:"betTable"` // 投注表：可选的投注金额档次
	CarType  []string    `json:"carType"`  // 车辆类型：不同类型的车辆品牌
	CarColor []string    `json:"carColor"` // 车辆颜色：红、蓝、绿、黄等颜色配置
}

// GroupSet 组合设置
type GroupSet struct {
	SymbolPosition []int `json:"symbolPosition"` // 符号位置：车辆在轮盘上的位置索引
	EventID        []int `json:"eventId"`        // 事件ID：触发的特殊事件类型
	EventWeight    []int `json:"eventWeight"`    // 事件权重：各事件的触发概率权重
	JackpotOdds    []int `json:"jackpotOdds"`    // 奖池赔率：累积奖池的赔率倍数
	JackpotWeight  []int `json:"jackpotWeight"`  // 奖池权重：触发奖池的概率权重
}

// LevelSet 等级设置
type LevelSet struct {
	Odds           []int `json:"odds"`           // 赔率：各等级的基础赔率倍数
	PositionWeight []int `json:"positionWeight"` // 位置权重：各位置的命中概率权重
	SpecialRate    []int `json:"specialRate"`    // 特殊率：触发特殊车辆的概率
}

// GameResult 游戏结果
type GameResult struct {
	LevelIndex         int            `json:"levelIndex"`         // 等级索引：当前游戏使用的等级配置索引
	MainCarPositionIdx int            `json:"mainCarPositionIdx"` // 主车位置索引：获胜车辆在轮盘上的位置
	HitCar             int            `json:"hitCar"`             // 命中车辆：获胜车辆的编号（1-8）
	MainGameOdds       int            `json:"mainGameOdds"`       // 主游戏赔率：基础投注的赔率倍数
	MainGameWin        int64          `json:"mainGameWin"`        // 主游戏获胜：基础投注的获胜金额
	EventWin           int64          `json:"eventWin"`           // 事件获胜：特殊事件的获胜金额
	JackpotWin         int64          `json:"jackpotWin"`         // 累积奖池获胜：奖池奖励金额
	TotalWin           int64          `json:"totalWin"`           // 总获胜：本局游戏的总获胜金额
	EventID            int            `json:"eventId"`            // 事件ID：触发的特殊事件类型
	EventResult        []*EventResult `json:"eventResult"`        // 事件结果：特殊事件的详细结果
}

// EventResult 事件结果
type EventResult struct {
	BonusHits    int `json:"bonusHits"`    // 奖励命中：特殊事件命中的位置或结果
	BonusHitCar  int `json:"bonusHitCar"`  // 奖励命中车辆：特殊事件命中的车辆信息
	BonusWinOdds int `json:"bonusWinOdds"` // 奖励获胜赔率：特殊事件的赔率倍数
}
