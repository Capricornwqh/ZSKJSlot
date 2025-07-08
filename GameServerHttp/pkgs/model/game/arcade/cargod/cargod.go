package model_cargod

// 定义特殊车辆的属性
type SpecialItem struct {
	LightCarColor int `json:"lightCarColor"` // 轻车颜色：1-红色，2-蓝色，3-绿色，4-黄色
	LightCarBrand int `json:"lightCarBrand"` // 轻车品牌：不同品牌代表不同的赔率档次
	SpecialCarNum int `json:"specialCarNum"` // 特殊车编号：用于标识特定的特殊车辆
}

// 每局游戏的开奖结果
type GameResult struct {
	LightCarBrand        int            `json:"lightCarBrand"`        // 获胜轻车品牌
	LightCarColor        int            `json:"lightCarColor"`        // 获胜轻车颜色
	LightCarNum          int            `json:"lightCarNum"`          // 获胜轻车编号
	LightCarWin          int            `json:"lightCarWin"`          // 轻车获胜标识：1-获胜，0-未获胜
	SpecialLightcar      int            `json:"specialLightcar"`      // 特殊轻车标识：是否为特殊车辆
	SpecialFlag          int            `json:"specialFlag"`          // 特殊标志：触发特殊玩法的标识
	JackpotOdds          int            `json:"jackpotOdds"`          // 累积奖池赔率：特殊奖励的倍数
	ZhuangXian           int            `json:"zhuangXian"`           // 庄闲结果：1-庄赢，2-闲赢，3-和局（如果启用百家乐玩法）
	SpecialColorBrandArr []*SpecialItem `json:"specialColorBrandArr"` // 特殊颜色品牌数组：本局出现的所有特殊车辆
}

// 玩家加入游戏时的下注配置
type JoinBetInfo struct {
	MinTotalBet    int            `json:"minTotalBet"`    // 最小总下注额：单局游戏最少下注金额
	MaxTotalBet    int            `json:"maxTotalBet"`    // 最大总下注额：单局游戏最多下注金额
	MaxSingleBet   int            `json:"maxSingleBet"`   // 最大单项下注：单个投注项最大金额
	MinSingleBet   int            `json:"minSingleBet"`   // 最小单项下注：单个投注项最小金额
	EnableBaccarat bool           `json:"enableBaccarat"` // 是否启用百家乐：启用后可投注庄闲
	BetList        []int          `json:"betList"`        // 可选下注金额列表：预设的下注金额选项
	WheelColor     []int          `json:"wheelColor"`     // 轮盘颜色配置：定义轮盘上的颜色分布
	OddsTable      *JoinOddsTable `json:"oddsTable"`      // 赔率表：各投注项的赔率配置
}

// 玩家结果
type ResultPlayerResult struct {
	Nickname string            `json:"nickname"` // 玩家昵称
	BetItems map[string][2]int `json:",inline"`  // 投注项目：key为投注类型，value为[投注金额, 获胜金额]
}

// 获取信息请求
type GetInfoRequest struct {
	GameName string `json:"gameName" binding:"required,min=1"` // 游戏名称：必须指定要查询的游戏
}

// 获取信息响应
type GetInfoResponse struct {
	RoomID      int       `json:"roomID,omitempty"`      // 房间ID：当前游戏房间编号
	RefreshTime int       `json:"refreshTime,omitempty"` // 刷新时间：数据刷新间隔（秒）
	TS          int       `json:"ts,omitempty"`          // 时间戳：服务器当前时间
	RoomInfo    *RoomInfo `json:"roomInfo,omitempty"`    // 房间信息：包含游戏状态和历史记录
	BetInfo     *BetInfo  `json:"betInfo,omitempty"`     // 下注信息：包含下注限制和赔率配置
}

// 游戏客制用内容
type JoinExtraData struct {
	AnnounceWinOdds []int `json:"AnnounceWinOdds"` // 公告获胜赔率：用于显示特殊赔率信息
}

// 加入游戏请求
type JoinRequest struct {
	GameName string `json:"gameName"` // 游戏名称：要加入的游戏类型
	RoomID   int    `json:"roomID"`   // 房间ID：要加入的房间编号
	History  int    `json:"history"`  // 是否获取历史记录：1-获取，0-不获取
}

// 加入游戏响应
type JoinResponse struct {
	RoomID    int            `json:"roomID,omitempty"`    // 房间ID：成功加入的房间编号
	RoundNo   string         `json:"roundNo,omitempty"`   // 局数：当前游戏局次编号
	StartTime int            `json:"startTime,omitempty"` // 开始时间：本局游戏开始时间戳
	EndTime   int            `json:"endTime,omitempty"`   // 结束时间：本局游戏结束时间戳
	TS        int            `json:"ts,omitempty"`        // 时间戳：服务器当前时间
	BetInfo   *JoinBetInfo   `json:"betInfo,omitempty"`   // 下注信息：包含下注限制和配置
	History   *History       `json:"history,omitempty"`   // 历史记录：近期游戏结果
	ExtraData *JoinExtraData `json:"extraData,omitempty"` // 额外数据：游戏特定信息
}

// 历史记录请求
type HistoryRequest struct {
	GameName string `json:"gameName"` // 游戏名称：要查询的游戏类型
	RoomID   int    `json:"roomID"`   // 房间ID：要查询的房间编号
	Num      int    `json:"num"`      // 获取的历史记录数量：最多返回的记录条数
}

// 刷新请求
type RefreshRequest struct {
	GameName string `json:"gameName" binding:"required"` // 游戏名称：必须指定游戏类型
	RoundNo  string `json:"roundNo" binding:"required"`  // 局数：必须指定当前局次
}

type RefreshResponse struct {
	PlayerNum     int                       `json:"playerNum,omitempty"`     // 玩家数量：当前房间内玩家总数
	BetData       map[string]int            `json:"betData,omitempty"`       // 下注数据：各投注项的总下注金额
	PlayerBetList map[string]map[string]int `json:"playerBetList,omitempty"` // 玩家下注列表：每个玩家的具体下注情况
	PlayerList    []string                  `json:"playerList,omitempty"`    // 玩家列表：当前房间内所有玩家昵称
	TS            int                       `json:"ts,omitempty"`            // 时间戳：数据更新时间
}

// 下注请求
type BetRequest struct {
	GameName string         `json:"gameName"` // 游戏名称：投注的游戏类型
	RoundNo  string         `json:"roundNo"`  // 局数：投注的游戏局次
	Confirm  int            `json:"confirm"`  // 确认标识：1-确认投注，0-取消投注
	BetData  map[string]int `json:"betData"`  // 投注数据：key为投注类型，value为投注金额
}

// 结果请求
type ResultRequest struct {
	GameName string `json:"gameName"` // 游戏名称：要查询结果的游戏类型
	RoundNo  string `json:"roundNo"`  // 局数：要查询结果的游戏局次
}

// 结果响应
type ResultResponse struct {
	TotalWin        int                            `json:"totalWin,omitempty"`        // 总获胜金额：玩家本局总收益
	Balance         int                            `json:"balance,omitempty"`         // 余额：玩家当前账户余额
	GameResult      *GameResult                    `json:"gameResult,omitempty"`      // 游戏结果：本局的开奖结果
	PlayerResult    *ResultPlayerResult            `json:"playerResult,omitempty"`    // 玩家结果：当前玩家的投注和收益情况
	AllPlayerResult map[string]*ResultPlayerResult `json:"allPlayerResult,omitempty"` // 所有玩家结果：房间内所有玩家的结果
}

// 历史记录
type History struct {
	DisplayedHistory []*HistoryItem `json:"displayedHistory"` // 显示最近出球：用于展示给玩家的历史开奖记录
}

// 历史记录项 - 单次游戏的历史记录
type HistoryItem struct {
	LightCarBrand   int   `json:"lightCarBrand"`   // 轻车品牌：获胜车辆的品牌
	LightCarColor   int   `json:"lightCarColor"`   // 轻车颜色：获胜车辆的颜色
	LightCarNum     int   `json:"lightCarNum"`     // 轻车号码：获胜车辆的编号
	SpecialFlag     int   `json:"specialFlag"`     // 特殊标志：是否触发特殊玩法
	SpecialLightcar []int `json:"specialLightcar"` // 特殊车颜色品牌：本局出现的特殊车辆信息
}

// 房间信息
type RoomInfo struct {
	RoomID          int      `json:"roomID"`          // 房间ID：房间唯一标识
	RoundNo         string   `json:"roundNo"`         // 局数：当前游戏局次编号
	StartTime       int      `json:"startTime"`       // 开始时间：本局游戏开始时间戳
	EndTime         int      `json:"endTime"`         // 结束时间：本局游戏结束时间戳
	AnnounceWinOdds []int    `json:"announceWinOdds"` // 公告获胜赔率：特殊赔率公告信息
	History         *History `json:"history"`         // 历史记录：近期游戏开奖历史
}

// 下注信息
type BetInfo struct {
	MaxTotalBet  int            `json:"maxTotalBet"`  // 最大总下注：单局最大下注总额
	MinTotalBet  int            `json:"minTotalBet"`  // 最小总下注：单局最小下注总额
	MaxSingleBet int            `json:"maxSingleBet"` // 最大单项下注：单个投注项最大金额
	MinSingleBet int            `json:"minSingleBet"` // 最小单项下注：单个投注项最小金额
	BetList      []int          `json:"betList"`      // 可选下注金额：预设的下注金额档次
	DefaultBet   int            `json:"defaultBet"`   // 默认下注：系统默认的下注金额
	OddsTable    *OddsTableItem `json:"oddsTable"`    // 赔率表：投注项的赔率配置
}

// 赔率表项
type OddsTableItem struct {
	Odds     int   `json:"odds,omitempty"`     // 赔率：该投注项的赔率倍数
	BetRange []int `json:"betRange,omitempty"` // 投注范围：该投注项允许的下注金额范围
}

// 赔率表
// key为投注类型（如：color_red、brand_1、num_1等）
type OddsTable map[string]OddsTableItem

// 加入游戏赔率表项
type JoinOddsTableItem struct {
	Odds            *int  `json:"odds"`            // 赔率：该投注项的赔率倍数（指针类型，可为空）
	BetRange        []int `json:"betRange"`        // 投注范围：允许的下注金额范围
	AnnounceWinodds []int `json:"announceWinOdds"` // 公告获胜赔率：特殊赔率公告
}

// 加入游戏赔率表 - 加入游戏时的完整赔率配置
// 投注类型说明：
// - color_red/blue/green/yellow：投注车辆颜色
// - brand_1/2/3/4：投注车辆品牌
// - num_1/2/3/4/5/6/7/8：投注车辆编号
// - special：投注特殊车辆
// - zhuang/xian：投注庄闲（如果启用百家乐）
type JoinOddsTable map[string]JoinOddsTableItem
