package Slot005GemBonanza

type Symbol int32

func (s Symbol) String() string {
	if s < 0 {
		return "NN"
	}
	return [...]string{"WW", "H1", "H2", "H3", "H4", "H5", "LA", "LK", "LQ", "LJ", "LT", "SS"}[s]
}

// ReelSymbol 轉輪帶
type ReelSymbol []Symbol

// GameSymbol 獎圖盤面
type GameSymbol []ReelSymbol

// TumbleResult 一次掉落的結果
type TumbleResult struct {
	TumbleSymbol GameSymbol `json:"TumbleSymbol"` // 當次獎圖盤面
	ReelIndex    []int      `json:"ReelIndex"`    // 初始盤面停輪 index
	H5ScoreArray [][]int    `json:"H5ScoreArray"` // H5 獎圖分數二維陣列
	LineSymbol   []Symbol   `json:"LineSymbol"`   // 連線獎圖陣列
	LineCount    []int      `json:"LineCount"`    // 連線個數陣列
	LineWin      []uint64   `json:"LineWin"`      // 連線贏分陣列
	Win          uint64     `json:"Win"`          // 當次盤面贏分
}

// MGResult 主遊戲結果
type MGResult struct {
	MGTumbleList        []TumbleResult `json:"MGTumbleList"`        // 主遊戲盤面列表
	MGGroupIndex        int            `json:"MGGroupIndex"`        // 主遊戲初始盤面轉輪群組 index
	MGPerformanceSymbol GameSymbol     `json:"MGPerformanceSymbol"` // 主遊戲表演盤面
	MGPerformanceType   int            `json:"MGPerformanceType"`   // 主遊戲表演類型 (0: 無表演, 1: 隨機替換一個 Scatter)
	MainWin             uint64         `json:"MainWin"`             // 主遊戲贏分
	FreeGameType        int            `json:"FreeGameType"`        // 免費遊戲類型 (即免費遊戲的轉輪群組 index)
}

// SpinResult 玩家 Spin 一次的結果 (與 MGResult 不同！)
type SpinResult struct {
	TumbleSymbol      GameSymbol `json:"TumbleSymbol"`      // 當次獎圖盤面
	ReelIndex         []int      `json:"ReelIndex"`         // 初始盤面停輪 index
	H5ScoreArray      [][]int    `json:"H5ScoreArray"`      // H5 獎圖分數二維陣列
	LineSymbol        []Symbol   `json:"LineSymbol"`        // 連線獎圖陣列
	LineCount         []int      `json:"LineCount"`         // 連線個數陣列
	LineWin           []uint64   `json:"LineWin"`           // 連線贏分陣列
	H5Win             uint64     `json:"H5Win"`             // H5 獎圖總贏分
	PerformanceSymbol GameSymbol `json:"PerformanceSymbol"` // 表演盤面
	PerformanceType   int        `json:"PerformanceType"`   // 表演類型 (0: 無表演, 1: 出現 Wild 的整輪替換, 2: 替換全部 H5 獎圖, 3: 替換 Wild 以外的獎圖)
	CumWildCount      int        `json:"CumWildCount"`      // 累積 Wild 個數
	Multiplier        int        `json:"Multiplier"`        // 乘倍
	Stage             int        `json:"Stage"`             // 目前階段
	SpinWin           uint64     `json:"SpinWin"`           // Spin 一次總贏分
}

// FGResult 免費遊戲結果
type FGResult struct {
	FGSpinList     []SpinResult `json:"FGSpinList"`     // 免費遊戲 Spin 結果列表
	FGSpinCount    int          `json:"FGSpinCount"`    // 免費遊戲 Spin 次數
	FGCumWildCount int          `json:"FGCumWildCount"` // 免費遊戲累積 Wild 個數
	FGMultiplier   int          `json:"FGMultiplier"`   // 免費遊戲乘倍
	FGStage        int          `json:"FGStage"`        // 免費遊戲最後階段
	FreeWin        uint64       `json:"FreeWin"`        // 免費遊戲贏分
}

type SlotResult struct {
	Code int `json:"Code"` // 回應代碼 (參考 ErrorCode.go)
	MGResult
	FGResult
	TotalWin uint64 `json:"TotalWin"` // 總贏分
	GameMode int    `json:"GameMode"` // 遊戲模式
	BuyType  int    `json:"BuyType"`  // 購買類型 (0: 未買, 1: Buy ExtraBet, 2: Buy FreeSpins, 3: Buy SuperFreeSpins)
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
