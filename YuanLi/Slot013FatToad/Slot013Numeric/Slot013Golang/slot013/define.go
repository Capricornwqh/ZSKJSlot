package slot013

const (
	BUY_NONE             = iota // 未購買，即一般投注
	BUY_EXTRA_BET               // 額外投注
	BUY_FREE_SPINS              // 購買免費遊戲
	BUY_SUPER_FREE_SPINS        // 購買超級免費遊戲
)

// 遊戲模式
const (
	GAME_MODE_NORMAL = 1 << iota // 一般模式
	GAME_MODE_BONUS              // 中副遊戲
	GAME_MODE_FREE               // 中免費遊戲
)

// 錯誤代碼
const (
	ERROR_CODE_OK = iota
	ERROR_CODE_UNMARSHAL
	ERROR_CODE_NO_SERVICE
	ERROR_CODE_NO_RESPONSE
	ERROR_CODE_NO_BALANCE
	ERROR_CODE_NO_TOKEN     // 找不到token
	ERROR_CODE_REPEAT_LOGIN // 重複登入
	ERROR_CODE_KICKED       // 後踢前或平台踢出
	ERROR_CODE_NO_GAME      // 找不到遊戲
	ERROR_CODE_NO_PLATFORM  // 無此平台
	ERROR_CODE_MAX_ODDS     // 超過最大贏分倍數
)

const (
	ERROR_CODE_PLATFORM_API = iota + 100 // 平台api錯誤
)

// 錯誤代碼 (遊戲相關)
const (
	ERROR_CODE_INVALID_RTP          = iota + 1000 // 無效的 RTP
	ERROR_CODE_INVALID_BUY_TYPE                   // 無效的購買類型
	ERROR_CODE_SETTING_NOT_FOUND                  // 未找到設定
	ERROR_CODE_ROULETTE_SPIN_FAILED               // 機率輪盤擲骰失敗
)

// 錯誤代碼 (測試相關)
const (
	ERROR_CODE_INVALID_DEBUG_CODE       = iota + 1100 // 無效的測試代碼
	ERROR_CODE_INVALID_DEBUG_SYMBOL                   // 無效的盤面
	ERROR_CODE_INVALID_DEBUG_FLAG                     // 無效的特別旗標
	ERROR_CODE_INVALID_DEBUG_REEL_INDEX               // 無效的停輪 index
	ERROR_CODE_INVALID_DEBUG_MULTIPLIER               // 無效的乘倍
)

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
	SpecialSpin  bool          `json:"-"`        // 特殊 Spin
	MaxWin       bool          `json:"MaxWin"`   // 最大贏分
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
	MGGroupIndex   int             `json:"MGGroupIndex"` // 主遊戲初始盤面轉輪群組 index
	MGTumbleList   []*TumbleResult `json:"MGTumbleList"` // 主遊戲盤面列表
	MainWin        uint64          `json:"MainWin"`      // 主遊戲贏分
	FreeGameType   int             `json:"FreeGameType"` // 免費遊戲類型 (即免費遊戲的轉輪群組 index)
	MGFeatureWin   uint64          `json:"MGFeatureWin"` // 主遊戲天降贏分
	MGFeatureCount int             `json:"-"`            // 天降次數
}

// FGResult 免費遊戲結果
type FGResult struct {
	FGSpinList     []*SpinResult `json:"FGSpinList"`     // 免費遊戲 Spin 結果列表
	FGSpinCount    int           `json:"FGSpinCount"`    // 免費遊戲 Spin 次數
	FGCumWildCount int           `json:"FGCumWildCount"` // 免費遊戲累積 Wild 個數
	FGMultiplier   int           `json:"FGMultiplier"`   // 免費遊戲乘倍
	FGIndex        int           `json:"-"`              // 免費遊戲配套index
	FreeWin        uint64        `json:"FreeWin"`        // 免費遊戲贏分
	FreeWinCount   int           `json:"FreeWinCount"`   // 免費遊戲贏分次數
	FGFeatureWin   uint64        `json:"FGFeatureWin"`   // 免費遊戲天降贏分
	FGFeatureCount int           `json:"-"`              // 天降次數
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
