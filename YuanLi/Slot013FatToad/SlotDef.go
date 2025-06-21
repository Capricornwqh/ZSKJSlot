package Slot013FatToad

type Symbol int32

func (s Symbol) String() string {
	if s < 0 {
		return "NN"
	}
	return [...]string{"WW", "H1", "H2", "H3", "H4", "LA", "LK", "LQ", "LJ", "SS", "SS2", "SS3", "SS5"}[s]
}

// ReelSymbol 轉輪帶
type ReelSymbol []Symbol

// GameSymbol 獎圖盤面
type GameSymbol []ReelSymbol

// TumbleResult 一次掉落的結果
type TumbleResult struct {
	TumbleSymbol      GameSymbol `json:"TumbleSymbol"`      // 當次獎圖盤面
	ReelIndex         []int      `json:"ReelIndex"`         // 初始盤面停輪 index
	MGPerformanceType int        `json:"MGPerformanceType"` // 主遊戲表演類型 (0: 無表演, 1: 天降横财, 2: 金蟾吃铜钱)
	LineSymbol        []Symbol   `json:"LineSymbol"`        // 連線獎圖陣列
	LineCount         []int      `json:"LineCount"`         // 連線個數陣列
	LineWin           []uint64   `json:"LineWin"`           // 連線贏分陣列
	Win               uint64     `json:"Win"`               // 當次盤面贏分
	SSCount           int        `json:"SSCount"`           // 铜钱数量
}

type SlotResult struct {
	Code    int    `json:"Code"`    // 回應代碼 (參考 ErrorCode.go)
	TraceId string `json:"TraceId"` // 追蹤 ID
	MGResult
	FGResult
	TotalBet     uint64        `json:"-"`        // 總投注額
	TotalWin     uint64        `json:"TotalWin"` // 總贏分
	GameMode     int           `json:"GameMode"` // 遊戲模式
	BuyType      int           `json:"BuyType"`  // 購買類型 (0: 未買, 1: Buy ExtraBet, 2: Buy FreeSpins, 3: Buy SuperFreeSpins)
	WWWild       []*WildStruct `json:"-"`        // 金蟾
	SSWild       []*WildStruct `json:"-"`        // 铜钱
	SSCount      int           `json:"-"`        // 铜钱数量
	WWLevel      int           `json:"-"`        // 金蟾等级
	WWMultiplier int           `json:"-"`        // 金蟾乘倍
}

// SpinResult 玩家 Spin 一次的結果 (與 MGResult 不同！)
type SpinResult struct {
	TumbleSymbol    GameSymbol `json:"TumbleSymbol"`    // 當次獎圖盤面
	ReelIndex       []int      `json:"ReelIndex"`       // 初始盤面停輪 index
	LineSymbol      []Symbol   `json:"LineSymbol"`      // 連線獎圖陣列
	LineCount       []int      `json:"LineCount"`       // 連線個數陣列
	LineWin         []uint64   `json:"LineWin"`         // 連線贏分陣列
	PerformanceType int        `json:"PerformanceType"` // 表演類型 (0: 無表演, 1: 天降横财, 2: 金蟾吃铜钱)
	CumWildCount    int        `json:"CumWildCount"`    // 累積 Wild 個數
	Multiplier      int        `json:"Multiplier"`      // 乘倍
	Stage           int        `json:"Stage"`           // 目前階段
	SpinWin         uint64     `json:"SpinWin"`         // Spin 一次總贏分
	SSCount         int        `json:"SSCount"`         // 铜钱数量
}

// MGResult 主遊戲結果
type MGResult struct {
	MGGroupIndex int             `json:"MGGroupIndex"` // 主遊戲初始盤面轉輪群組 index
	MGTumbleList []*TumbleResult `json:"MGTumbleList"` // 主遊戲盤面列表
	MainWin      uint64          `json:"MainWin"`      // 主遊戲贏分
	FreeGameType int             `json:"FreeGameType"` // 免費遊戲類型 (即免費遊戲的轉輪群組 index)
}

// FGResult 免費遊戲結果
type FGResult struct {
	FGSpinList     []*SpinResult `json:"FGSpinList"`     // 免費遊戲 Spin 結果列表
	FGSpinCount    int           `json:"FGSpinCount"`    // 免費遊戲 Spin 次數
	FGCumWildCount int           `json:"FGCumWildCount"` // 免費遊戲累積 Wild 個數
	FGMultiplier   int           `json:"FGMultiplier"`   // 免費遊戲乘倍
	FGIndex        int           `json:"-"`              // 免費遊戲配套index
	FreeWin        uint64        `json:"FreeWin"`        // 免費遊戲贏分
}

// DebugCmd 測試指令
type DebugCmd struct {
	DebugData []int
}

type SpinCmd struct {
	RTP          int        `json:"RTP,omitempty"` // RTP 編號，不是由 client 傳
	TotalBet     int        `json:"TotalBet"`      // 總投注額
	BuyType      int        `json:"BuyType"`       // 購買類型 (0: 未買, 1: Buy ExtraBet, 2: Buy FreeSpins, 3: Buy SuperFreeSpins)
	DebugCmdList []DebugCmd `json:"DebugCmdList"`  // 測試指令列表
}
