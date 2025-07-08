package slot013

// 版本
const PROB_VERSION = "slot013-97.00-v1.0.0"

const (
	PAYLINE_TOTAL = 50 //中奖线数

	DEFAULT_RTP = 0 // 預設 RTP 編號

	SLOT_COL = 6 //轮盘列数
	SLOT_ROW = 6 //轮盘行数

	DEFAU_RTP = 0 //预设RTP编号
	RTP_TOTAL = 1 //RTP总数

	WIN_SCATTER_COUNT     = 3 // 中 3 個 Scatter 可進免費遊戲
	MAX_WIN_SCATTER_COUNT = 5 // 中 5 個 Scatter 以上賠率相同、觸發免費遊戲次數相同
	MAX_COVER_WILD_COUNT  = 2 // 免費遊戲最大覆蓋 Wild 數量

	MAX_ODDS = 25000 //最大赔率
)

const (
	SSWILD_MAIN_GAME_1 = 15 // 主遊戲盘面1的铜钱概率
	SSWILD_MAIN_GAME_2 = 20 // 主遊戲盘面2的铜钱概率
	SSWILD_BUY_FREE    = 20 // 主遊戲buy free的铜钱倍数概率
	SSWILD_BUY_SUPER   = 1  // 主遊戲buy super的铜钱倍数概率

	FEATURE_WILD_MAIN_GAME = 2 // 主遊戲天降横财概率

	FREE1_LEVEL1_1 = 100 // 免費遊戲配套1的free1概率
	FREE1_LEVEL2_1 = 100 // 免費遊戲配套1的free2概率
	FREE1_LEVEL1_2 = 100 // 免費遊戲配套2的free1概率
	FREE1_LEVEL2_2 = 100 // 免費遊戲配套2的free1概率
	FREE2_LEVEL3_2 = 100 // 免費遊戲配套2的free2概率
	FREE2_LEVEL4_2 = 100 // 免費遊戲配套2的free2概率

	SSWILD_FREE_GAME_2 = 1 // 免費遊戲盤面2的铜钱概率

	SSWILD_BUY_FREE_2 = 4 // 免費遊戲buy free配套2的铜钱倍数概率

	SSWILD_BUY_SUPER_2 = 5 // 免費遊戲buy super配套2的铜钱倍数概率

	FEATURE_BUY_NONE_LEVEL1_1 = 50 // 免費遊戲天降横财配套1概率
	FEATURE_BUY_NONE_LEVEL2_1 = 50 // 免費遊戲天降横财配套1概率
	FEATURE_BUY_NONE_LEVEL1_2 = 60 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_NONE_LEVEL2_2 = 30 // 免費遊戲天降横财配套2概率

	FEATURE_BUY_FREE_LEVEL1_1 = 70 // 免費遊戲天降横财配套1概率
	FEATURE_BUY_FREE_LEVEL2_1 = 20 // 免費遊戲天降横财配套1概率
	FEATURE_BUY_FREE_LEVEL3_1 = 10 // 免費遊戲天降横财配套1概率
	FEATURE_BUY_FREE_LEVEL1_2 = 50 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_FREE_LEVEL2_2 = 20 // 免費遊戲天降横财配套2概率

	FEATURE_BUY_SUPER_LEVEL2_1 = 30 // 免費遊戲天降横财配套1概率
	FEATURE_BUY_SUPER_LEVEL3_1 = 20 // 免費遊戲天降横财配套1概率
	FEATURE_BUY_SUPER_LEVEL4_1 = 20 // 免費遊戲天降横财配套1概率
	FEATURE_BUY_SUPER_LEVEL2_2 = 30 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_SUPER_LEVEL3_2 = 20 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_SUPER_LEVEL4_2 = 10 // 免費遊戲天降横财配套2概率
)

const (
	NN  Symbol = iota - 1
	WW         // 三足金蟾
	H1         // 金元宝
	H2         // 翠玉白菜
	H3         // 灯笼
	H4         // 玉葫芦
	LA         // A
	LK         // K
	LQ         // Q
	LJ         // J
	SS         // 铜钱x1
	SS2        // 铜钱x2
	SS3        // 铜钱x3
	SS5        // 铜钱x5
	MAX_SYMBOL
)

// 測試指令資料 index
const (
	DEBUG_MAIN_GROUP_INDEX         = iota // 主遊戲轉輪群組 index (0~3)
	DEBUG_MAIN_REEL_INDEX_01              // 主游戏1轴停輪位置
	DEBUG_MAIN_REEL_INDEX_02              // 主游戏2轴停輪位置
	DEBUG_MAIN_REEL_INDEX_03              // 主游戏3轴停輪位置
	DEBUG_MAIN_REEL_INDEX_04              // 主游戏4轴停輪位置
	DEBUG_MAIN_REEL_INDEX_05              // 主游戏5轴停輪位置
	DEBUG_MAIN_REEL_INDEX_06              // 主游戏6轴停輪位置
	DEBUG_MAIN_FEATURE_COUNT              // 主遊戲天降横财铜钱數量 (0~4)
	DEBUG_MAIN_FEATURE_MULTIPLE_01        // 主遊戲天降横财铜钱1倍数
	DEBUG_MAIN_FEATURE_MULTIPLE_02        // 主遊戲天降横财铜钱2倍数
	DEBUG_MAIN_FEATURE_MULTIPLE_03        // 主遊戲天降横财铜钱3倍数
	DEBUG_MAIN_FEATURE_MULTIPLE_04        // 主遊戲天降横财铜钱4倍数
	DEBUG_MAIN_FEATURE_MULTIPLE_05        // 主遊戲天降横财铜钱5倍数
	DEBUG_FREE_GROUP_INDEX                // 免費遊戲遊戲轉輪群組 index (0~4)
	DEBUG_FREE_REEL_INDEX_01              // 免費遊戲1轴停輪位置
	DEBUG_FREE_REEL_INDEX_02              // 免費遊戲2轴停輪位置
	DEBUG_FREE_REEL_INDEX_03              // 免費遊戲3轴停輪位置
	DEBUG_FREE_REEL_INDEX_04              // 免費遊戲4轴停輪位置
	DEBUG_FREE_REEL_INDEX_05              // 免費遊戲5轴停輪位置
	DEBUG_FREE_REEL_INDEX_06              // 免費遊戲6轴停輪位置
	DEBUG_FREE_FEATURE_COUNT              // 免费遊戲天降横财铜钱數量 (0~4)
	DEBUG_FREE_FEATURE_MULTIPLE_01        // 免費遊戲天降横财铜钱1倍数
	DEBUG_FREE_FEATURE_MULTIPLE_02        // 免費遊戲天降横财铜钱2倍数
	DEBUG_FREE_FEATURE_MULTIPLE_03        // 免費遊戲天降横财铜钱3倍数
	DEBUG_FREE_FEATURE_MULTIPLE_04        // 免費遊戲天降横财铜钱4倍数
	DEBUG_FREE_FEATURE_MULTIPLE_05        // 免費遊戲天降横财铜钱5倍数
	DEBUG_MY_SYMBOL_INDEX                 // My符號索引 (0~3)
)

// 主遊戲 PerformanceType 表演類型
const (
	NONE                int = iota // 無表演
	PERFORMANCE_FEATURE            // 天降横财
	PERFORMANCE_EAT                // 金蟾吃铜钱
	PERFORMANCE_LEVELUP            // 金蟾升级
)

const (
	MAIN_GAME_01    int = iota // 主遊戲轉輪群組 01
	MAIN_GAME_02               // 主遊戲轉輪群組 02
	MAIN_GAME_FREE             // buy free
	MAIN_GAME_SUPER            // buy super free
)

const (
	FREE_GAME_01 int = iota // 免費遊戲轉輪群組 01
	FREE_GAME_02            // 免費遊戲轉輪群組 02
	FREE_GAME_03            // 免費遊戲轉輪群組 03
	FREE_GAME_04            // 免費遊戲轉輪群組 04
	FREE_GAME_05            // 免費遊戲轉輪群組 05
)

const (
	WWLEVEL_SSWILD_1 = 5  // 金蟾等級1的铜钱數量
	WWLEVEL_SSWILD_2 = 9  // 金蟾等級2的铜钱數量
	WWLEVEL_SSWILD_3 = 13 // 金蟾等級3的铜钱數量
	WWLEVEL_SSWILD_4 = 16 // 金蟾等級4的铜钱數量
	WWLEVEL_SSWILD_5 = 19 // 金蟾等級5的铜钱數量
)

const (
	FREE_INDEX_1 = iota // 免費遊戲配套索引1
	FREE_INDEX_2        // 免費遊戲配套索引2
)

// FGInitSpinCount 免費遊戲初始場次
var FGInitSpinCount = []int{0, 5, 8, 11, 13, 15, 16}

// WinSymbolCount 中獎獎圖連線個數 (達到才算中獎) [獎圖]個數
var WinSymbolCount = [MAX_SYMBOL]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, SLOT_COL + 1}

// BetRatio 投注額比例
var BetRatio = map[int]float64{
	BUY_NONE:             1,
	BUY_EXTRA_BET:        1.5,
	BUY_FREE_SPINS:       100,
	BUY_SUPER_FREE_SPINS: 500,
}

// 中奖线
var LineIndexArray = [][]int{
	{0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 1},
	{2, 2, 2, 2, 2, 2},
	{3, 3, 3, 3, 3, 3},
	{4, 4, 4, 4, 4, 4},
	{5, 5, 5, 5, 5, 5},
	{0, 1, 2, 3, 4, 5},
	{5, 4, 3, 2, 1, 0},
	{0, 1, 0, 1, 0, 1},
	{1, 2, 1, 2, 1, 2},
	{2, 3, 2, 3, 2, 3},
	{3, 4, 3, 4, 3, 4},
	{4, 5, 4, 5, 4, 5},
	{1, 0, 1, 0, 1, 0},
	{2, 1, 2, 1, 2, 1},
	{3, 2, 3, 2, 3, 2},
	{4, 3, 4, 3, 4, 3},
	{5, 4, 5, 4, 5, 4},
	{0, 1, 2, 2, 1, 0},
	{1, 2, 3, 3, 2, 1},
	{2, 3, 4, 4, 3, 2},
	{3, 4, 5, 5, 4, 3},
	{5, 4, 3, 3, 4, 5},
	{4, 3, 2, 2, 3, 4},
	{3, 2, 1, 1, 2, 3},
	{2, 1, 0, 0, 1, 2},
	{0, 1, 1, 1, 1, 0},
	{1, 2, 2, 2, 2, 1},
	{2, 3, 3, 3, 3, 2},
	{3, 4, 4, 4, 4, 3},
	{4, 5, 5, 5, 5, 4},
	{1, 0, 0, 0, 0, 1},
	{2, 1, 1, 1, 1, 2},
	{3, 2, 2, 2, 2, 3},
	{4, 3, 3, 3, 3, 4},
	{5, 4, 4, 4, 4, 5},
	{0, 1, 1, 1, 1, 2},
	{1, 2, 2, 2, 2, 3},
	{2, 3, 3, 3, 3, 4},
	{3, 4, 4, 4, 4, 5},
	{5, 4, 4, 4, 4, 3},
	{4, 3, 3, 3, 3, 2},
	{3, 2, 2, 2, 2, 1},
	{2, 1, 1, 1, 1, 0},
	{0, 5, 0, 5, 0, 5},
	{5, 0, 5, 0, 5, 0},
	{0, 2, 0, 2, 0, 2},
	{2, 0, 2, 0, 2, 0},
	{3, 5, 3, 5, 3, 5},
	{5, 3, 5, 3, 5, 3},
}

// 赔率
var SymbolOdds = [][]int{
	{0, 0, 0, 50, 200, 800, 2000}, // WW
	{0, 0, 0, 30, 100, 500, 1000}, // H1
	{0, 0, 0, 25, 80, 400, 800},   // H2
	{0, 0, 0, 15, 50, 200, 500},   // H3
	{0, 0, 0, 15, 50, 200, 500},   // H4
	{0, 0, 0, 5, 10, 25, 50},      // LA
	{0, 0, 0, 5, 10, 25, 50},      // LK
	{0, 0, 0, 5, 10, 25, 50},      // LQ
	{0, 0, 0, 5, 10, 25, 50},      // LJ
	{0, 0, 0, 50, 200, 800, 2000}, // SS
	{0, 0, 0, 50, 200, 800, 2000}, // SS2
	{0, 0, 0, 50, 200, 800, 2000}, // SS3
	{0, 0, 0, 50, 200, 800, 2000}, // SS5
}

// MGReelGroupWT 主遊戲轉輪群組權重
var MGReelGroupWT = []map[int][]uint{
	// RTP0 上線數值
	{
		BUY_NONE:             {374, 1000, 0, 0},
		BUY_EXTRA_BET:        {374, 1000, 0, 0},
		BUY_FREE_SPINS:       {0, 0, 1, 0},
		BUY_SUPER_FREE_SPINS: {0, 0, 0, 1},
	},
}

// MGReelGroup 主遊戲轉輪群組
var MGReelGroup = [][][][]Symbol{
	// RTP0 上線數值
	{
		// MAIN_GAME_01
		{
			{LJ, LJ, LJ, H1, H1, H1, H4, H4, H4, LQ, LQ, LK, LK, LK, H3, H3, H3, H4, H4, H4, LA, LA, LA, WW, LJ, LJ, H2, H2, H2, LQ, LQ, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, H3, LQ, LQ, LQ, H2, H2, H2, LA, LA, LK, LK, H1, H1, LJ, LJ, LJ, H4, H4, LJ, LJ, LJ, H1, H1, H1, H4, H4, H4, LQ, LQ, LK, LK, LK, H3, H3, H3, H4, H4, H4, LA, LA, LA, WW, LJ, LJ, H2, H2, H2, LQ, LQ, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, H3, LQ, LQ, LQ, H2, H2, H2, LA, LA, LK, LK, H1, H1, LJ, LJ, LJ, H4, H4},
			{LK, LK, LK, H2, H2, H2, H3, H3, H3, LA, LA, LJ, LJ, LJ, H4, H4, H4, H1, H1, H1, LQ, LQ, LQ, LJ, LJ, H1, H1, LA, LA, H4, H4, H4, H4, H4, H4, LK, LK, LQ, LQ, LQ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H1, H1, H4, H4, LK, LK, LK, H3, H3, H4, H4, LK, LK, LK, H2, H2, H2, H3, H3, H3, LA, LA, LA, LJ, LJ, H4, H4, H4, H1, H1, H1, LQ, LQ, LQ, LJ, LJ, H1, H1, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LQ, LQ, LQ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H1, H1, H4, H4, LK, LK, LK, H3, H3, H4, H4, LK, LK, LK, H3, H3, H3, H3, H3, H3, LA, LA, LJ, LJ, SS, H4, H4, H4, H1, H1, H1, LQ, LQ, LQ, LJ, LJ, H1, H1, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LQ, LQ, LQ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H1, H1, H4, H4, LK, LK, LK, H3, H3, H4, H4, LK, LK, LK, H2, H2, H2, H3, H3, H3, LA, LA, LJ, LJ, LJ, H4, H4, H4, H1, H1, H1, LQ, LQ, LQ, LJ, LJ, H1, H1, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LQ, LQ, LQ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H1, H1, H4, H4, LK, LK, LK, H3, H3, H4, H4},
			{LA, LA, LA, H4, H4, H4, H1, H1, H1, LK, LK, LJ, LJ, SS, LQ, LQ, LQ, H3, H3, H3, H3, H3, H3, LK, LK, LK, H1, H1, H1, LA, LA, LJ, LJ, LJ, H2, H2, H2, LQ, LQ, LQ, LA, LA, LA, H3, H3, H3, H4, H4, H4, LK, LK, LJ, LJ, LJ, H4, H4, H4, LQ, LQ, LQ, LA, LA, LA, H4, H4, H4, H1, H1, H1, LK, LK, SS, LJ, LJ, LQ, LQ, LQ, H3, H3, H3, H2, H2, H2, LK, LK, LK, H1, H1, H1, LA, LA, LJ, LJ, LJ, H2, H2, H2, LQ, LQ, LQ, LA, LA, LA, H3, H3, H3, H4, H4, H4, LK, LK, LJ, LJ, LJ, H4, H4, H4, LQ, LQ, LQ, LA, LA, LA, H1, H1, H1, H1, H1, H1, LK, LK, LJ, LJ, SS, LQ, LQ, LQ, H3, H3, H3, H2, H2, H2, LK, LK, LK, H1, H1, H1, LA, LA, LJ, LJ, LJ, H2, H2, H2, LQ, LQ, LQ, LA, LA, LA, H3, H3, H3, H4, H4, H4, LK, LK, LJ, LJ, LJ, H4, H4, H4, LQ, LQ, LQ, LA, LA, LA, H4, H4, H4, H1, H1, H1, LK, LK, LJ, LJ, LJ, LQ, LQ, LQ, H3, H3, H3, H2, H2, H2, LK, LK, LK, H1, H1, H1, LA, LA, LJ, LJ, LJ, H2, H2, H2, LQ, LQ, LQ, LA, LA, LA, H3, H3, H3, H4, H4, H4, LK, LK, LJ, LJ, LJ, H4, H4, H4, LQ, LQ, LQ},
			{LK, LK, LK, H3, H3, H3, H2, H2, H2, LQ, LQ, LJ, LJ, SS, LJ, LJ, LJ, H1, H1, H1, LA, LA, LA, H3, H3, H3, H3, H3, H3, LQ, LQ, LQ, H3, H3, H4, H4, H4, LA, LA, LA, LK, LK, LK, H4, H4, H4, LJ, LJ, LK, LK, H1, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA, LK, LK, LK, H3, H3, H3, H2, H2, H2, LQ, LQ, LJ, LJ, LJ, LJ, LJ, LJ, H1, H1, H1, LA, LA, LA, H2, H2, H2, H3, H3, H3, LQ, LQ, LQ, H3, H3, H4, H4, H4, LA, LA, LA, LK, LK, LK, H4, H4, H4, LJ, LJ, LK, LK, H1, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA, LK, LK, LK, H2, H2, H2, H2, H2, H2, LQ, LQ, LJ, LJ, LJ, LJ, LJ, LJ, H1, H1, H1, LA, LA, LA, H2, H2, H2, H3, H3, H3, LQ, LQ, LQ, H3, H3, H4, H4, H4, LA, LA, LA, LK, LK, LK, H4, H4, H4, LJ, LJ, LK, LK, H1, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA, LK, LK, LK, H3, H3, H3, H2, H2, H2, LQ, LQ, LJ, LJ, LJ, LJ, LJ, LJ, H1, H1, H1, LA, LA, LA, H2, H2, H2, H3, H3, H3, LQ, LQ, LQ, H3, H3, H4, H4, H4, LA, LA, LA, LK, LK, LK, H4, H4, H4, LJ, LJ, LK, LK, H1, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA},
			{LA, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LQ, LQ, SS, LA, LA, LA, H3, H3, H3, LQ, LQ, LQ, H1, H1, H1, H1, H1, H1, H3, H3, H3, LK, LK, LK, H2, H2, LA, LA, LA, LQ, LQ, LQ, H1, H1, H1, H4, H4, H4, LJ, LJ, LJ, LK, LK, LK, H2, H2, H2, LJ, LJ, SS, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LQ, LQ, LQ, LA, LA, LA, H3, H3, H3, LQ, LQ, LQ, H1, H1, H1, LJ, LJ, LJ, H3, H3, H3, LK, LK, LK, H2, H2, LA, LA, LA, SS, LQ, LQ, H1, H1, H1, H4, H4, H4, LJ, LJ, LJ, LK, LK, LK, H2, H2, H2, LJ, LJ, LA, LA, LA, H4, H4, H4, H4, H4, H4, LK, LK, LQ, LQ, LQ, LA, LA, LA, H3, H3, H3, LQ, LQ, LQ, H1, H1, H1, LJ, LJ, LJ, H3, H3, H3, LK, LK, LK, H2, H2, LA, LA, LA, LQ, LQ, LQ, H1, H1, H1, H4, H4, H4, LJ, LJ, LJ, SS, LK, LK, H2, H2, H2, LJ, LJ, LA, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LQ, LQ, LQ, LA, LA, LA, H3, H3, H3, LQ, LQ, LQ, H1, H1, H1, LJ, LJ, LJ, H3, H3, H3, LK, LK, LK, H2, H2, LA, LA, LA, LQ, LQ, LQ, H1, H1, H1, H4, H4, H4, LJ, LJ, LJ, LK, LK, LK, H2, H2, H2, LJ, LJ},
			{LK, LK, LK, H1, H1, H1, H1, H1, H1, LA, LA, LQ, LQ, SS, LJ, LJ, LJ, H3, H3, H3, H3, H3, H3, LA, LA, LK, LK, H2, H2, H2, LQ, LQ, LQ, H1, H1, H1, H1, LJ, LJ, LJ, H2, H2, H2, H2, H2, H2, LQ, LQ, LQ, H4, H4, H4, LA, LA, LA, H3, H3, H3, LJ, LJ, LK, LK, LK, H1, H1, H1, H3, H3, H3, LA, LA, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, H4, H4, H4, H4, LA, LA, LK, LK, H2, H2, H2, LQ, LQ, LQ, H1, H1, H1, H1, LJ, LJ, LJ, LK, LK, LK, H2, H2, H2, LQ, LQ, LQ, H4, H4, H4, LA, LA, LA, H3, H3, H3, LJ, LJ, LK, LK, LK, H1, H1, H1, H3, H3, H3, LA, LA, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, H4, H3, H3, H3, LA, LA, LK, LK, H2, H2, H2, LQ, LQ, LQ, H1, H1, H1, H1, LJ, LJ, LJ, LK, LK, LK, H2, H2, H2, LQ, LQ, LQ, H4, H4, H4, LA, LA, LA, H3, H3, H3, LJ, LJ, LK, LK, LK, H1, H1, H1, H3, H3, H3, LA, LA, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, H4, H3, H3, H3, LA, LA, LK, LK, H2, H2, H2, LQ, LQ, LQ, H1, H1, H1, H1, LJ, LJ, LJ, LK, LK, LK, H2, H2, H2, LQ, LQ, LQ, H4, H4, H4, LA, LA, LA, H3, H3, H3, LJ, LJ},
		},
		// MAIN_GAME_02
		{
			{LJ, LJ, LJ, H1, H1, H1, H1, H1, H1, LQ, LQ, LK, LK, H3, H3, H3, H3, H3, H3, H3, LA, LA, LA, LJ, LJ, H2, H2, H2, H2, H2, H2, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, H3, LQ, LQ, LQ, H4, H4, H4, H4, H4, LK, LK, H1, H1, LJ, LJ, LJ, H4, H4},
			{LK, LK, LK, H2, H2, H2, H3, H3, H3, LA, LA, LJ, LJ, LJ, H4, H4, H4, H1, H1, H1, LJ, LJ, LJ, LJ, LJ, H1, H1, LA, LA, H2, H2, H2, H2, H2, H2, LK, LK, H3, H3, H3, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H1, H1, H1, H1, LK, LK, LK, H4, H4, H4, H4, LK, LK, LK, H2, H2, H2, H3, H3, H3, LA, LA, LJ, LJ, SS, H4, H4, H4, H1, H1, H1, LQ, LQ, LQ, LJ, LJ, H1, H1, LA, LA, H2, H2, H2, H2, H2, H2, LK, LK, SS, LQ, LQ, H3, H3, H3, H3, H3, H3, LJ, LJ, LJ, H1, H1, H1, H1, LK, LK, LK, H3, H3, H4, H4},
			{LA, LA, LA, H4, H4, H4, H1, H1, H1, LK, LK, LJ, LJ, LJ, LQ, LQ, LQ, H2, H2, H2, H2, H2, H2, LK, LK, LK, H1, H1, H1, LA, LA, LJ, LJ, LJ, H2, H2, H2, LQ, LQ, LQ, SS, LA, LA, H3, H3, H3, H3, H3, H3, LK, LK, LJ, LJ, LJ, H4, H4, H4, LQ, LQ, LQ, LA, LA, LA, H4, H4, H4, H4, H4, H4, LK, LK, LJ, LJ, SS, LQ, LQ, LQ, H3, H3, H3, H3, H3, H3, LK, LK, LK, H1, H1, H1, LA, LA, LJ, LJ, LJ, H2, H2, H2, LQ, LQ, LQ, SS, LA, LA, H3, H3, H3, H4, H4, H4, LK, LK, LJ, LJ, LJ, H4, H4, H4, LQ, LQ, LQ},
			{LK, LK, LK, H3, H3, H3, H2, H2, H2, LQ, LQ, LJ, LJ, LJ, LJ, LJ, LJ, H1, H1, H1, LA, LA, LA, H2, H2, H2, H3, H3, H3, LQ, LQ, LQ, H3, H3, H4, H4, H4, LA, LA, LA, SS, LK, LK, H4, H4, H4, LJ, LJ, LK, LK, H1, H1, H1, H1, H1, H1, H4, H4, LA, LA, SS, LK, LK, H2, H2, H2, H2, H2, H2, LQ, LQ, LJ, LJ, SS, LJ, LJ, LJ, H1, H1, H1, LA, LA, LA, H3, H3, H3, H3, H3, H3, LQ, LQ, LQ, H4, H4, H4, H4, H4, H4, LA, LA, SS, LK, LK, H4, H4, H4, LJ, LJ, LK, LK, H1, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA},
			{LA, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LQ, LQ, LQ, LA, LA, LA, H3, H3, H3, LQ, LQ, LQ, H1, H1, H1, LJ, LJ, LJ, H3, H3, H3, LK, LK, LK, H2, H2, LA, LA, LA, SS, LQ, LQ, H1, H1, H1, H4, H4, H4, LJ, LJ, LJ, SS, LK, LK, H2, H2, H2, LJ, LJ, LA, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LQ, LQ, SS, LA, LA, LA, H3, H3, H3, LQ, LQ, LQ, H1, H1, H1, LJ, LJ, LJ, H3, H3, H3, LK, LK, LK, H2, H2, LA, LA, LA, SS, LQ, LQ, H1, H1, H1, H1, H1, H1, LJ, LJ, LJ, H2, H2, H2, H2, H2, H2, LJ, LJ},
			{LK, LK, LK, H1, H1, H1, H1, H1, H1, LA, LA, LQ, LQ, LQ, LJ, LJ, LJ, H3, H3, H3, H3, H3, H3, LA, LA, LK, LK, LQ, LQ, LQ, LQ, LQ, LQ, H1, H1, H1, H1, LJ, LJ, LJ, SS, LK, LK, H2, H2, H2, LQ, LQ, LQ, H4, H4, H4, LA, LA, LA, H3, H3, H3, LJ, LJ, SS, LK, LK, H2, H2, H2, H2, H2, H2, LA, LA, LQ, LQ, SS, LJ, LJ, LJ, H4, H4, H4, H3, H3, H3, LA, LA, LK, LK, H2, H2, H2, LQ, LQ, LQ, H1, H1, H1, H1, LJ, LJ, LJ, SS, LK, LK, H2, H2, H2, LQ, LQ, LQ, H4, H4, H4, LA, LA, LA, H3, H3, H3, LJ, LJ},
		},
		// MAIN_GAME_FREE
		{
			{LJ, LJ, LJ, WW, H3, H3, LA, LA, LA, WW, H4, H4, LK, LK, LK, WW, LQ, LQ, LQ, H4, H4, WW, LA, LA, H1, H1, LK, WW, LJ, LJ, LJ, H3, H3, WW, LQ, LQ},
			{LK, LK, H4, H4, H4, SS, LJ, LJ, LJ, H3, H3, SS, LA, LA, LA, LQ, LQ, SS, LK, LK, LK, LJ, LJ, SS, LA, LA, LQ, LQ, LQ, SS, LJ, H4, H4, H2, LA, SS, LA, H1, H1, LQ, LQ, SS, LA, LA, LK, LQ, LQ, SS, LJ, LJ, LJ, H4, H4, SS},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H4, H4, LA, LA, LA, LK, H3, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LK, LK, SS, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, SS, LJ, LJ, LJ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, SS, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, H3, SS, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H4, H4, LA, LA, SS, LK, LK, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, LJ, LJ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, H3, SS, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4},
		},
		// MAIN_GAME_SUPER
		{
			{LJ, LJ, LJ, WW, H3, H3, LA, LA, LA, WW, H4, H4, LK, LK, LK, WW, LQ, LQ, LQ, H4, H4, WW, LA, LA, H1, H1, LK, WW, LJ, LJ, LJ, H3, H3, WW, LQ, LQ},
			{LK, LK, H4, H4, H4, SS, LJ, LJ, LJ, H3, H3, SS, LA, LA, LA, LQ, LQ, SS, LK, LK, LK, LJ, LJ, SS, LA, LA, LQ, LQ, LQ, SS, LJ, H2, H2, H2, LA, SS, LK, LK, LK, LQ, LQ, SS, LA, LA, LK, LQ, LQ, SS, LJ, LJ, LJ, H3, H3, SS},
			{LQ, LQ, H3, H3, H3, SS, LK, LK, LK, H4, H4, SS, LA, LA, LA, LQ, LQ, SS, LK, LK, LK, LJ, LJ, SS, LA, LA, LQ, LQ, LQ, SS, LJ, H3, H3, H3, LA, SS, LA, H1, H1, LQ, LQ, SS, LA, LA, LK, LQ, LQ, SS, LJ, LJ, LJ, LA, LA, SS},
			{LK, LK, H4, H4, H4, SS, LJ, LJ, LJ, H3, H3, SS, LA, LA, LA, LQ, LQ, SS, LK, LK, LK, LJ, LJ, SS, LA, LA, LQ, LQ, LQ, SS, LJ, H2, H2, H2, LA, SS, LA, LA, LA, LQ, LQ, SS, LA, LA, LK, LQ, LQ, SS, LJ, LJ, LJ, H4, H4, SS},
			{LK, LK, H4, H4, H4, SS, LJ, LJ, LJ, H3, H3, SS, LA, LA, LA, LQ, LQ, SS, LK, LK, LK, LJ, LJ, SS, LA, LA, LQ, LQ, LQ, SS, LJ, H3, H3, H3, LA, SS, LQ, LQ, LQ, LJ, LJ, SS, LA, LA, LK, LQ, LQ, SS, LJ, LJ, LJ, LQ, LQ, SS},
			{LK, LK, H4, H4, H4, SS, LJ, LJ, LJ, H3, H3, SS, LA, LA, LA, LQ, LQ, SS, LK, LK, LK, LJ, LJ, SS, LA, LA, LQ, LQ, LQ, SS, LJ, H1, H1, H1, LA, SS, LA, LA, LA, LQ, LQ, SS, LA, LA, LK, LQ, LQ, SS, LJ, LJ, LJ, LK, LK, SS},
		},
	},
}

// FGReelGroup 免費遊戲轉輪群組
var FGReelGroup = [][][][]Symbol{
	// RTP0 上線數值
	{
		// FREE_GAME_01
		{
			{H2, H2, H2, LQ, LQ, LQ, NN, NN, NN, LQ, LQ, LK, LK, LK, H1, H1, H1, NN, NN, NN, LA, LA, LA, LQ, LQ, LQ, NN, NN, NN, NN, NN, NN, LK, LK, LK, H4, H4, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, NN, NN, NN, NN, NN, NN, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, LQ, LQ, LQ, SS, LK, LK, NN, NN, NN, LQ, LQ, LQ, LQ, H1, H1, LA, LA, LK, LK, LK, NN, NN, NN, LQ, LQ, LQ, H4, H4, H4, LK, LK, NN, NN, NN, NN, LK, LK, LK, H4, H4},
			{LK, LK, LK, LJ, LJ, LJ, NN, NN, NN, LA, LA, LQ, LQ, LQ, H1, H1, H1, NN, NN, NN, LJ, LJ, LJ, LQ, LQ, H3, H3, LA, LA, NN, NN, NN, NN, NN, NN, LA, LA, SS, LK, LK, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, NN, NN, NN, NN, LA, LA, LA, H3, H3, H3, H3, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, NN, NN, NN, H2, H2, H2, LA, LA, LJ, LJ, LJ, NN, NN, NN, LA, LA, LA, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, H3, H3, H3, H3},
			{LK, LK, LK, LQ, LQ, LQ, NN, NN, NN, LQ, LQ, LK, LK, LK, H1, H1, H1, NN, NN, NN, LK, LK, LK, LQ, LQ, LQ, NN, NN, NN, LK, LK, LK, LK, LK, NN, NN, NN, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, NN, NN, NN, LK, LK, LK, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, LA, LA, LA, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, LQ, H1, H1, LA, LA, SS, LK, LK, NN, NN, NN, LQ, LQ, LQ, H3, H3, H3, LA, LA, LQ, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4},
			{LA, LA, LA, LJ, LJ, LJ, H1, H1, H1, LA, LA, SS, LJ, LJ, H2, H2, H2, NN, NN, NN, LJ, LJ, LJ, LJ, LJ, NN, NN, NN, NN, LK, LK, LK, LQ, LQ, LQ, LA, LA, LJ, LJ, LJ, NN, NN, NN, NN, NN, NN, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, H3, H3, H3, SS, LJ, LJ, LJ, LJ, NN, NN, NN, LA, LA, H2, H2, H2, H2, H2, H2, LA, LA, SS, LJ, LJ, H3, H3, H3, LA, LA, LA, SS, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, H3, H3, H3},
			{LK, LK, LK, LQ, LQ, LQ, NN, NN, NN, LQ, LQ, SS, LK, LK, H2, H2, H2, H4, H4, H4, LK, LK, LK, SS, LQ, LQ, H4, H4, H4, LK, LK, LK, LK, LK, SS, H4, H4, LK, LK, LK, H3, H3, H3, LQ, LQ, LQ, NN, NN, NN, LK, LK, LK, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, LA, LA, LA, SS, LK, LK, H1, H1, H1, LQ, LQ, LQ, LQ, H1, H1, LA, LA, SS, LK, LK, H1, H1, H1, LQ, LQ, SS, H4, H4, H4, LA, LA, SS, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4},
			{LA, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, SS, LJ, LJ, H2, H2, H2, H3, H3, H3, SS, LJ, LJ, LJ, H3, H3, H3, LA, LA, H2, H2, H2, H3, H3, H3, LA, LA, LJ, LJ, LJ, NN, NN, NN, LA, LA, SS, LJ, LJ, LJ, H3, H3, H2, H2, H2, LA, LA, LA, NN, NN, NN, LJ, LJ, LQ, LQ, LQ, H3, H3, LA, LA, SS, H2, H2, H2, H2, H2, LA, LA, SS, LJ, LJ, H3, H3, H3, LA, LA, SS, LK, LK, LK, H3, H3, H2, H2, SS, LA, LA, LA, H3, H3, H3},
		},
		// FREE_GAME_02
		{
			{LJ, LJ, LJ, LA, LA, LA, H4, H4, H4, LQ, LQ, LK, LK, LK, H2, H2, H2, H4, H4, H4, LK, LK, LK, LA, LA, LA, H4, H4, H4, H3, H3, H3, LA, LA, LA, H3, H3, H3, LJ, LJ, H1, H1, H1, LQ, LQ, LQ, H2, H2, H2, LK, LK, LK, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, LQ, LQ, LQ, LK, LK, LK, NN, NN, NN, LJ, LJ, LJ, LJ, H1, H1, LA, LA, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, NN, NN, NN, NN, NN, LQ, LQ, LQ, LQ, LK, LK, LK, H4, H4},
			{LA, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LQ, LQ, H1, H1, H1, H3, H3, H3, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, H2, H2, H2, H3, H3, H3, LA, LA, LK, LK, LK, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, NN, NN, NN, NN, LA, LA, LA, H3, H3, H3, H3, LJ, LJ, LJ, H1, H1, H1, LA, LA, LA, H2, H2, H2, H2, H2, H2, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, NN, NN, NN, H3, H3, H2, H2, LA, LA, LA, H3, H3, H3, H3},
			{LK, LK, LK, LQ, LQ, LQ, NN, NN, NN, LQ, LQ, LK, LK, LK, H1, H1, H1, H2, H2, H2, LK, LK, LK, SS, LQ, LQ, H4, H4, H4, LK, LK, LK, LK, LK, LK, NN, NN, NN, NN, NN, H1, H1, H1, LQ, LQ, LQ, H4, H4, H4, LK, LK, LK, NN, NN, NN, LJ, LJ, LJ, H4, H4, LA, LA, LA, SS, LK, LK, NN, NN, NN, LQ, LQ, LQ, LQ, H1, H1, LA, LA, LK, LK, LK, H1, H1, H1, NN, NN, NN, H3, H3, H3, LA, LA, LQ, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4},
			{LA, LA, LA, LJ, LJ, LJ, H1, H1, H1, LA, LA, LJ, LJ, SS, H2, H2, H2, H4, H4, H4, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, LK, LK, LK, LQ, LQ, LQ, LA, LA, LJ, LJ, LJ, NN, NN, NN, LA, LA, LA, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, SS, H3, H3, H3, LJ, LJ, LJ, LJ, LJ, NN, NN, NN, NN, H2, H2, H2, H2, H2, H2, LA, LA, SS, LJ, LJ, NN, NN, NN, LA, LA, LA, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, H3, H3, H3, H3},
			{LK, LK, LK, LQ, LQ, LQ, H4, H4, H4, LQ, LQ, SS, LK, LK, H2, H2, H2, H4, H4, H4, LK, LK, LK, NN, NN, NN, NN, NN, NN, LK, LK, LK, LK, LK, H4, H4, H4, LK, LK, LK, H3, H3, H3, LQ, LQ, LQ, H4, H4, H4, LK, LK, LK, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, LA, LA, LA, LK, LK, LK, H1, H1, H1, SS, LQ, LQ, LQ, H1, H1, LA, LA, SS, LK, LK, H1, H1, H1, LQ, LQ, LQ, H4, H4, H4, LA, LA, SS, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4},
			{LA, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, SS, LJ, LJ, H2, H2, H2, H3, H3, H3, LJ, LJ, LJ, LJ, SS, H3, H3, LA, LA, H2, H2, H2, H3, H3, H3, LA, LA, LJ, LJ, LJ, NN, NN, NN, NN, NN, NN, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, H4, H4, H4, H4, LJ, LJ, LQ, LQ, LQ, H3, H3, LA, LA, H2, H2, H2, NN, NN, NN, LA, LA, SS, LJ, LJ, H3, H3, H3, LA, LA, SS, LK, LK, LK, NN, NN, NN, H2, LA, LA, LA, H3, H3, H3, H3},
		},
		// FREE_GAME_03
		{
			{LA, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LA, LA, H4, H4, H4, H4, H4, H4, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, H4, H4, H4, H3, H3, H3, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H3, H3, H4, H4, LA, LA, LA, H3, H3, H3, H3, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, H4, H4, H4, H4, H4, H4, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H3, H3, H4, H4, LA, LA, LA, H3, H3, H3, H3},
			{LK, LK, LK, LQ, LQ, LQ, H2, H2, H2, LQ, LQ, LK, LK, LK, H1, H1, H1, H1, H1, H1, LK, LK, LK, LQ, LQ, LQ, H2, H2, H2, LA, LA, LA, LA, LA, LA, H2, H2, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, H2, H2, H2, LA, LA, LA, LQ, LQ, LQ, LA, LA, LA, H2, H2, LQ, LQ, LQ, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, LQ, H1, H1, LA, LA, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, H2, H2, H2, LK, LK, LQ, LQ, LQ, LQ, LK, LK, LK, H2, H2},
			{LK, LK, LK, LQ, LQ, LQ, H4, H4, H4, LQ, LQ, LJ, LJ, LJ, H1, H1, H1, H2, H2, H2, LJ, LJ, LJ, LQ, LQ, LQ, H4, H4, H4, LK, LK, LK, LK, LK, LK, H4, H4, LJ, LJ, LJ, H1, H1, H1, LQ, LQ, LQ, H4, H4, H4, LJ, LJ, LJ, LQ, LQ, LQ, H2, H2, H2, H4, H4, H3, H3, H3, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, LQ, H1, H1, H1, H1, LJ, LJ, LJ, H1, H1, H1, LQ, LQ, LQ, H3, H3, H3, H4, H4, LQ, LQ, LQ, LQ, H2, H2, H2, H4, H4},
			{LA, LA, LA, LJ, LJ, LJ, H1, H1, H1, LA, LA, SS, LJ, LJ, H2, H2, H2, H4, H4, H4, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, LK, LK, LK, LQ, LQ, SS, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, H3, H3, H3, H3, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, H2, H2, H2, H2, H2, H2, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, H3, H3, H3, H3},
			{LK, LK, LK, LQ, LQ, LQ, H4, H4, H4, LQ, LQ, SS, LK, LK, H2, H2, H2, H4, H4, H4, LK, LK, LK, LQ, LQ, LQ, H4, H4, H4, SS, LK, LK, LK, LK, LK, H4, H4, H4, LK, LK, H3, H3, H3, SS, LQ, LQ, H4, H4, H4, LK, LK, SS, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, LA, LA, SS, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, SS, H1, H1, LA, LA, LA, LK, LK, H1, H1, H1, SS, LQ, LQ, H4, H4, H4, LA, LA, SS, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4},
			{LA, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, SS, LJ, LJ, H2, H2, H2, H3, H3, H3, LJ, LJ, LJ, LJ, H3, H3, H3, LA, LA, H2, H2, H2, H3, H3, H3, LA, LA, SS, LJ, LJ, H1, H1, H1, LA, LA, SS, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, SS, H4, H4, H4, LJ, LJ, LQ, LQ, SS, H3, H3, LA, LA, H2, H2, H2, H2, H2, H2, LA, LA, SS, LJ, LJ, H3, H3, H3, LA, LA, SS, LK, LK, LK, H3, H3, H2, SS, LA, LA, LA, H3, H3, H3, H3},
		},
		// FREE_GAME_04
		{
			{NN, NN, NN, NN, NN, NN, LA, LA, LA, NN, NN, NN, NN, LA, LA, LA, NN, NN, NN, NN, NN, NN, LK, LK, LK, NN, NN, NN, LQ, LQ, LQ, NN, NN, NN, NN, NN, NN, LJ, LJ, LJ},
			{NN, NN, NN, NN, NN, NN, LA, LA, LA, NN, NN, NN, NN, LA, LA, LA, NN, NN, NN, NN, NN, NN, LK, LK, LK, NN, NN, NN, LQ, LQ, LQ, NN, NN, NN, NN, NN, NN, LJ, LJ, LJ},
			{NN, NN, NN, NN, NN, NN, LA, LA, LA, NN, NN, NN, NN, LA, LA, LA, NN, NN, NN, NN, NN, NN, LK, LK, LK, NN, NN, NN, LQ, LQ, LQ, NN, NN, NN, NN, NN, NN, LJ, LJ, LJ},
			{H2, H2, H2, H1, H1, H1, LA, LA, LA, NN, NN, NN, NN, LA, LA, LA, NN, NN, NN, NN, NN, NN, LK, LK, LK, NN, NN, NN, LQ, LQ, LQ, H3, H3, H3, H4, H4, H4, LJ, LJ, LJ},
			{H2, H2, H2, LA, LA, H1, H1, H1, H2, H2, H2, LQ, LQ, H1, H1, H1, LK, LK, H4, H4, LA, LA, LA, NN, NN, NN, LJ, LJ, LK, LK, H2, H2, H2, LQ, LQ, H1, H1, H1, LJ, LJ},
			{H2, H2, H2, LA, LA, H1, H1, H1, H2, H2, H2, LQ, LQ, H1, H1, H1, LK, LK, H4, H4, LA, LA, LA, H3, H3, H3, LJ, LJ, LK, LK, NN, NN, NN, NN, NN, NN, H4, H4, LJ, LJ},
		},
		// FREE_GAME_05
		{
			{LA, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LA, LA, H1, H1, H1, H4, H4, H4, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, H2, H2, H2, H3, H3, H3, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H3, H3, H4, H4, LA, LA, LA, H3, H3, H3, H3, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, H4, H4, H4, H4, H4, H4, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H3, H3, H4, H4, LA, LA, LA, H3, H3, H3, H3},
			{LK, LK, LK, LQ, LQ, LQ, H2, H2, H2, LQ, LQ, LK, LK, LK, H1, H1, H1, H1, H1, H1, LK, LK, LK, H4, H4, H4, H2, H2, H2, LA, LA, LA, LA, LA, LA, H2, H2, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, H3, H3, H3, LA, LA, LA, LQ, LQ, LQ, LA, LA, LA, H2, H2, LQ, LQ, LQ, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, LQ, H1, H1, LA, LA, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, H2, H2, H2, LK, LK, LQ, LQ, LQ, LQ, LK, LK, LK, H2, H2},
			{LK, LK, LK, LQ, LQ, LQ, H4, H4, H4, LQ, LQ, LJ, LJ, LJ, H1, H1, H1, H2, H2, H2, LJ, LJ, LJ, LQ, LQ, LQ, H4, H4, H4, LK, LK, LK, LK, LK, LK, H4, H4, LJ, LJ, LJ, H1, H1, H1, LQ, LQ, LQ, H4, H4, H4, LJ, LJ, LJ, LQ, LQ, LQ, H2, H2, H2, H4, H4, H3, H3, H3, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, LQ, H1, H1, H1, H1, LJ, LJ, LJ, H1, H1, H1, LQ, LQ, LQ, H3, H3, H3, H4, H4, LQ, LQ, LQ, LQ, H2, H2, H2, H4, H4},
			{LA, LA, LA, LJ, LJ, LJ, H1, H1, H1, LA, LA, LA, LJ, LJ, H2, H2, H2, H4, H4, H4, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, LK, LK, LK, LQ, LQ, SS, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, H3, H3, H3, H3, LJ, LJ, LJ, LJ, LJ, H3, H3, LA, LA, H2, H2, H2, H2, H2, H2, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, H3, H3, H3, H3},
			{LK, LK, LK, LQ, LQ, LQ, H4, H4, H4, LQ, LQ, LQ, LK, LK, H2, H2, H2, H4, H4, H4, LK, LK, LK, LQ, LQ, LQ, H4, H4, H4, H4, LK, LK, LK, LK, LK, H4, H4, LK, LK, LK, H3, H3, H3, LQ, LQ, LQ, H4, H4, H4, LK, LK, LK, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, LA, LA, LA, LK, LK, LK, H1, H1, H1, LQ, LQ, LQ, H1, H1, H1, LA, LA, LA, LK, LK, H1, H1, H1, LQ, LQ, LQ, H4, H4, H4, LA, LA, SS, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4},
			{LA, LA, LA, LJ, LJ, LJ, H3, H3, H3, LA, LA, LA, LJ, LJ, H2, H2, H2, H3, H3, H3, LJ, LJ, LJ, LJ, H3, H3, H3, LA, LA, H2, H2, H2, H3, H3, H3, LA, LA, LA, LJ, LJ, H1, H1, H1, LA, LA, LA, LJ, LJ, LJ, H3, H3, H2, H2, LA, LA, LA, SS, H4, H4, H4, LJ, LJ, LQ, LQ, LQ, H3, H3, LA, LA, H2, H2, H2, H2, H2, H2, LA, LA, LA, LJ, LJ, H3, H3, H3, LA, LA, LA, LK, LK, LK, H3, H3, H2, H2, LA, LA, LA, H3, H3, H3, H3},
		},
	},
}

// MysteryList Mystery 轴图标
var MysteryList = []Symbol{H1, H2, H3, H4}

// MysteryColWT Mystery 轴图标权重
var MysteryColWT = []uint{40, 40, 40, 40}

// MysteryGloballWT Mystery 轴图标权重
var MysteryGloballWT = []uint{60, 40, 0, 0}

// MGSSMulti  主游戏铜钱倍数列表
var MGSSMultiList = []Symbol{SS2, SS3, SS5}

// MGSSMultiWT 主游戏铜钱倍数权重
var MGSSMultiWT = []uint{100, 50, 10}

// MGBuySSMultiWT 主游戏铜钱倍数权重2
var MGBuySSMultiWT = []uint{100, 60, 10}

// MGFeatureSSList 主游戏天降横财铜钱个数
var MGFeatureSSList = []uint{1, 2, 3, 4, 5}

// MGFeatureSSWT 主游戏天降横财铜钱个数权重
var MGFeatureSSWT = []uint{0, 820, 300, 100, 50}

// FGIndexList 免費遊戲配套索引列表
var FGIndexList = []int{FREE_INDEX_1, FREE_INDEX_2}

// FGIndexWT 免費遊戲配套索引權重
var FGIndexWT = []uint{130, 110}

// FGBuyFreeIndexWT 免費遊戲buy配套索引權重
var FGBuyFreeIndexWT = []uint{815, 200}

// FGBuySuperIndexWT 免費遊戲buy配套索引權重
var FGBuySuperIndexWT = []uint{760, 240}

// FGSSMultiList 免費遊戲銅錢倍數列表
var FGSSMultiList = []Symbol{SS2, SS3, SS5}

// FGSSMultiWT 免費遊戲銅錢倍數權重
var FGSSMultiWT = []uint{100, 40, 10}

// FGGroupIndexList 免費遊戲轉輪群組索引列表
var FGGroupIndexList = []int{FREE_GAME_01, FREE_GAME_02, FREE_GAME_03, FREE_GAME_04, FREE_GAME_05}

// FGLevelGroupIndexWT 免費遊戲轉輪群組索引權重
var FGLevelGroupIndexWT = [][2][7][]uint{
	BUY_NONE: {
		0: {
			0: {0, 0, 0, 0, 0},
			1: {100, 0, 100, 0, 0},
			2: {0, 100, 0, 0, 100},
			3: {0, 0, 100, 0, 100},
			4: {0, 0, 100, 0, 100},
			5: {0, 0, 100, 0, 100},
			6: {0, 0, 100, 0, 100},
		},
		1: {
			0: {0, 0, 0, 0, 0},
			1: {100, 0, 100, 0, 0},
			2: {0, 100, 100, 0, 0},
			3: {0, 100, 100, 0, 0},
			4: {0, 100, 100, 0, 0},
			5: {0, 100, 0, 0, 0},
			6: {0, 100, 0, 0, 0},
		},
	},
	BUY_EXTRA_BET: {},
	BUY_FREE_SPINS: {
		0: {
			0: {0, 0, 0, 0, 0},
			1: {100, 0, 100, 0, 0},
			2: {0, 200, 100, 0, 0},
			3: {0, 200, 100, 0, 0},
			4: {0, 0, 100, 0, 0},
			5: {0, 0, 100, 0, 0},
			6: {0, 0, 100, 0, 0},
		},
		1: {
			0: {0, 0, 0, 0, 0},
			1: {100, 0, 100, 0, 0},
			2: {100, 0, 100, 0, 0},
			3: {100, 0, 100, 0, 0},
			4: {0, 100, 0, 0, 0},
			5: {0, 100, 0, 0, 0},
			6: {0, 100, 0, 0, 0},
		},
	},
	BUY_SUPER_FREE_SPINS: {
		0: {
			0: {0, 0, 0, 0, 0},
			1: {0, 200, 50, 0, 0},
			2: {0, 200, 50, 0, 0},
			3: {0, 100, 0, 0, 0},
			4: {0, 100, 0, 0, 0},
			5: {0, 100, 0, 0, 0},
			6: {0, 100, 0, 0, 0},
		},
		1: {
			0: {0, 0, 0, 0, 0},
			1: {100, 0, 100, 0, 0},
			2: {100, 0, 100, 0, 0},
			3: {0, 100, 0, 0, 0},
			4: {0, 100, 0, 0, 0},
			5: {0, 100, 0, 0, 0},
			6: {0, 100, 0, 0, 0},
		},
	},
}
