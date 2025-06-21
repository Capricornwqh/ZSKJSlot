package Slot013FatToad

import "Force/GameServer/Common"

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
	SSWILD_MAIN_GAME_1 = 5  // 主遊戲盘面1的铜钱概率
	SSWILD_MAIN_GAME_2 = 15 // 主遊戲盘面2的铜钱概率
	SSWILD_BUY_SUPER   = 1  // 主遊戲buy super的铜钱倍数概率

	FEATURE_WILD_MAIN_GAME = 15 // 主遊戲天降横财概率

	FREE_MYSTERY_1 = 50 // 免費遊戲配套1的保底概率
	FREE_MYSTERY_2 = 20 // 免費遊戲配套2的保底概率
	FREE_MYSTERY_3 = 20 // 免費遊戲配套3的保底概率

	SSWILD_FREE_GAME_2 = 10 // 免費遊戲盤面2的铜钱概率

	SSWILD_BUY_FREE_1 = 0 // 免費遊戲buy free配套1的铜钱倍数概率
	SSWILD_BUY_FREE_2 = 1 // 免費遊戲buy free配套2的铜钱倍数概率
	SSWILD_BUY_FREE_3 = 0 // 免費遊戲buy free配套3的铜钱倍数概率

	SSWILD_BUY_SUPER_1 = 0 // 免費遊戲buy super配套1的铜钱倍数概率
	SSWILD_BUY_SUPER_2 = 2 // 免費遊戲buy super配套2的铜钱倍数概率
	SSWILD_BUY_SUPER_3 = 2 // 免費遊戲buy super配套3的铜钱倍数概率

	FEATURE_BUY_NONE_LEVEL1_1 = 15 // 免費遊戲天降横财配套1概率
	FEATURE_BUY_NONE_LEVEL2_1 = 5  // 免費遊戲天降横财配套1概率
	FEATURE_BUY_NONE_LEVEL1_2 = 22 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_NONE_LEVEL2_2 = 12 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_NONE_LEVEL1_3 = 25 // 免費遊戲天降横财配套3概率
	FEATURE_BUY_NONE_LEVEL2_3 = 15 // 免費遊戲天降横财配套3概率
	FEATURE_BUY_NONE_LEVEL3_3 = 10 // 免費遊戲天降横财配套3概率

	FEATURE_BUY_FREE_LEVEL1_1 = 5  // 免費遊戲天降横财配套1概率
	FEATURE_BUY_FREE_LEVEL2_1 = 5  // 免費遊戲天降横财配套1概率
	FEATURE_BUY_FREE_LEVEL1_2 = 10 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_FREE_LEVEL2_2 = 10 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_FREE_LEVEL1_3 = 10 // 免費遊戲天降横财配套3概率
	FEATURE_BUY_FREE_LEVEL2_3 = 10 // 免費遊戲天降横财配套3概率

	FEATURE_BUY_SUPER_LEVEL1_1 = 15 // 免費遊戲天降横财配套1概率
	FEATURE_BUY_SUPER_LEVEL2_1 = 5  // 免費遊戲天降横财配套1概率
	FEATURE_BUY_SUPER_LEVEL1_2 = 25 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_SUPER_LEVEL2_2 = 20 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_SUPER_LEVEL3_2 = 20 // 免費遊戲天降横财配套2概率
	FEATURE_BUY_SUPER_LEVEL1_3 = 25 // 免費遊戲天降横财配套3概率
	FEATURE_BUY_SUPER_LEVEL2_3 = 15 // 免費遊戲天降横财配套3概率
	FEATURE_BUY_SUPER_LEVEL3_3 = 10 // 免費遊戲天降横财配套3概率
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
	DEBUG_INDEX_GROUP_INDEX   = iota // 主遊戲轉輪群組 index (0~5)
	DEBUG_INDEX_REEL_INDEX_01        // 停輪位置
	DEBUG_INDEX_REEL_INDEX_02
	DEBUG_INDEX_REEL_INDEX_03
	DEBUG_INDEX_REEL_INDEX_04
	DEBUG_INDEX_REEL_INDEX_05
	DEBUG_INDEX_REEL_INDEX_06
	DEBUG_INDEX_COVER_SCATTER_COUNT // 主遊戲覆蓋 Scatter 獎圖數量 (3~5)
	DEBUG_INDEX_FREE_GAME_TYPE      // 免費遊戲類型 (即免費遊戲的轉輪群組 index, 0~1)
	DEBUG_INDEX_COVER_WILD_COUNT    // 免費遊戲覆蓋 Wild 數量 (0~2)
	DEBUG_INDEX_PERFORMANCE_TYPE    // 表演類型 (主遊戲: 0~1，免費遊戲: 0~3)
)

// 主遊戲 PerformanceType 表演類型
const (
	NONE                int = iota // 無表演
	PERFORMANCE_FEATURE            // 天降横财
	PERFORMANCE_EAT                // 金蟾吃铜钱
	PERFORMANCE_LEVELUP            // 金蟾升级
)

const (
	MAIN_GAME_01      int = iota // 主遊戲轉輪群組 01
	MAIN_GAME_02                 // 主遊戲轉輪群組 02
	MAIN_GAME_MYSTERY            // 保底轉輪群組 (隨機選擇轉輪群組)
	MAIN_GAME_FREE               // buy free
	MAIN_GAME_SUPER              // buy super free
)

const (
	FREE_GAME_01      int = iota // 免費遊戲轉輪群組 01
	FREE_GAME_02                 // 免費遊戲轉輪群組 02
	FREE_GAME_MYSTERY            // 免費遊戲保底轉輪群組 (隨機選擇轉輪群組)
	BUY_FREE_GAME_01             // buy free和buy super
	BUY_FREE_GAME_02             // buy free和buy super
	BUY_FREE_GAME_03             // buy free和buy super
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
	FREE_INDEX_3        // 免費遊戲配套索引3
)

// FGInitSpinCount 免費遊戲初始場次
var FGInitSpinCount = []int{0, 5, 8, 11, 13, 15, 16}

// WinSymbolCount 中獎獎圖連線個數 (達到才算中獎) [獎圖]個數
var WinSymbolCount = [MAX_SYMBOL]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, SLOT_COL + 1}

// BetRatio 投注額比例
var BetRatio = map[int]float64{
	Common.BUY_NONE:             1,
	Common.BUY_EXTRA_BET:        1.5,
	Common.BUY_FREE_SPINS:       100,
	Common.BUY_SUPER_FREE_SPINS: 500,
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
		Common.BUY_NONE:             {200, 94, 5, 0, 0},
		Common.BUY_EXTRA_BET:        {200, 94, 5, 0, 0},
		Common.BUY_FREE_SPINS:       {0, 0, 0, 1, 0},
		Common.BUY_SUPER_FREE_SPINS: {0, 0, 0, 0, 1},
	},
}

// MGReelGroup 主遊戲轉輪群組
var MGReelGroup = [][][][]Symbol{
	// RTP0 上線數值
	{
		// MAIN_GAME_01
		{
			{LJ, LJ, LJ, H3, H3, LA, LA, H2, H2, H4, H4, LK, LK, LA, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA, H3, H3, LK, LA, WW, LJ, LJ, H2, H2, H3, LQ, LQ, H1, H1, H3, H3, LA, LA, H2, H2, H4, H4, LK, LK, LQ, LQ, H1, H1, LJ, LJ, LJ, H2, H2, LQ, LQ, LK, LK, H3, H3, H4, LA, LA, H2, H2, H4, H4, LK, LK, H3, H3, H3, LQ, LQ, LQ, LJ, LJ, H1, H1, H1, LA, LK, LQ, H2, H2, LJ, LJ, H3, H4, LK, LK, LK, H2, H2, LQ, LJ, H4, H3, LQ, LQ, LJ, LJ, LJ, H2, H2, LQ, LQ, LK, LK, H3, H3, H4, LA, LA, H2, H2, H4, H4, LK, LK, H3, H3, H3, LQ, LQ, LJ, LJ, H1, H1, LA, LK, LQ, H2, H2, LJ, LJ, H3, H4, LK, LK, LK, H2, H2, H4, H4, H4, LQ, LQ, LQ, H1, H1, H1},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, H3, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LQ, LQ, LQ, H2, H2, H2, H4, H4, H4, LA, LA, H1, H1, H1, H3, H3, H3, LQ, LQ, LQ, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H3, H3, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, H4, SS, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, LQ, LQ, H3, H3, H3, H4, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H4},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H4, H4, LA, LA, LA, LK, H3, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, H4, LJ, LJ, LJ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK, LQ, LQ, H1, H1, H2, H2, LA, LA, SS, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, LJ, LJ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK, H3, H3, H3},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H3, H3, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H4, H4, LA, LA, SS, LK, LK, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, LJ, LJ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK, LQ, LQ, H1, H1, H2, H2, LA, LA, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, LJ, LJ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK, H2, H2, H2},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, LQ, SS, H3, H3, H3, H4, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H4},
		},
		// MAIN_GAME_02
		{
			{LJ, LJ, LJ, H3, H3, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LA, H1, H1, LQ, LQ, LQ, H4, H4, H4, LA, LA, LA, H3, H3, H3, LA, LA, LJ, LJ, H2, H2, H3, LQ, LQ, H1, H1, H3, H3, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LK, H1, H1, H1, LQ, LQ, LJ, LJ, LJ, H2, H2, H2, LQ, LQ, LK, LK, H3, H3, H4, LA, LA, H2, H2, H4, H4, LK, LK, H3, H3, H3, LQ, LQ, LJ, LJ, H1, H1, H1, LA, LK, LQ, H2, H2, LJ, LJ, H3, H4, LK, LK, LK, H2, H2, LQ, LJ, H4, H3, LQ, LQ, LJ, LJ, LJ},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, LJ, H3, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LJ, LJ, LJ, H2, H2, H2, LA, LA, LA, H1, H1, H1, LQ, LQ, LQ, H3, H4, LA, LK, LQ, LQ, LQ, H2, H2, H2, LJ, LJ, H3, H3, H4, LA, LA, H1, H1, LJ, LJ, SS, H2, H3, LK, LQ, H4, H4, LK, LK, SS, LA, LA, H2, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H4, H3, H3, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LJ, LJ, LJ, H1, H1, H1, LK, LJ, SS, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, LK, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H2, H4, H4, H4, LA, LA, LA, SS, LK, LK, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LA, SS, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, SS, LK, LQ, H4, H4, H4, LJ, LJ, H1, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, H4, LJ, LJ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, H1, LK, LK, LK, LQ},
			{LK, LK, H4, H4, H4, LQ, LQ, LJ, LJ, H3, H3, H3, LA, LA, H1, H1, H2, H2, H2, LK, LK, LK, H3, H3, LJ, LJ, H1, H1, H1, H2, LQ, LQ, SS, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, SS, LQ, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, SS, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, LA, SS, LQ, LQ, H3, H3, H3, H4, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, LJ, H2, H3, H1, LK, LK, LK, H4, H4, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, SS, LJ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, LK, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H4, H4, LA, LA, SS, LK, LK, LK, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, SS, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, SS, LQ, H4, H4, LJ, LJ, LJ, H1, H2, LK, LK, LK, H3, H3, H3, LA, LA, H4, H4, H4, LJ, LJ, LJ, SS, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK},
			{LK, LK, H4, H4, H4, LQ, LQ, LQ, SS, LJ, LJ, H3, H3, LA, LA, H1, H1, H1, H2, H2, H2, LK, LK, LK, SS, H3, H3, LJ, LJ, H2, H2, H2, LQ, LQ, LQ, LJ, LJ, H4, H4, H4, LA, LA, LA, H1, H1, H1, LQ, LQ, H3, H4, LA, LK, SS, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, SS, LA, H2, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H3},
		},
		// MAIN_GAME_MYSTERY
		{
			{NN, NN, NN, LA, LA, NN, NN, NN, H2, H2, NN, NN, NN, H1, H1, NN, NN, NN, H4, H4, NN, NN, NN, H3, H3, NN, NN, NN, LK, LK, NN, NN, NN, LQ, LQ, NN, NN, NN, LJ, LJ},
			{NN, NN, NN, LA, LA, NN, NN, NN, H2, H2, NN, NN, NN, H1, H1, NN, NN, NN, H4, H4, NN, NN, NN, H3, H3, NN, NN, NN, LK, LK, NN, NN, NN, LQ, LQ, NN, NN, NN, LJ, LJ},
			{NN, NN, NN, LA, LA, NN, NN, NN, H2, H2, NN, NN, NN, H1, H1, NN, NN, NN, H4, H4, NN, NN, NN, H3, H3, NN, NN, NN, LK, LK, NN, NN, NN, LQ, LQ, NN, NN, NN, LJ, LJ},
			{H2, H2, H2, LA, LA, H1, H1, H1, NN, NN, NN, LQ, LQ, H1, H1, H1, LK, LK, H4, H4, LA, LA, LA, NN, NN, NN, LJ, LJ, LK, LK, H3, H3, H3, LQ, LQ, H4, H4, H4, LJ, LJ},
			{H2, H2, H2, LA, LA, H1, H1, H1, H2, H2, H2, LQ, LQ, H1, H1, H1, LK, LK, H4, H4, LA, LA, LA, NN, NN, NN, LJ, LJ, LK, LK, H2, H2, H2, LQ, LQ, H1, H1, H1, LJ, LJ},
			{H2, H2, H2, LA, LA, H1, H1, H1, H2, H2, H2, LQ, LQ, H1, H1, H1, LK, LK, H4, H4, LA, LA, LA, H3, H3, H3, LJ, LJ, LK, LK, H2, H2, H2, LQ, LQ, H4, H4, H4, LJ, LJ},
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
			{LJ, LJ, LJ, H3, H3, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LA, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA, H3, H3, SS, LA, LA, LA, LJ, LJ, H2, H2, H3, LQ, LQ, H1, H1, H3, H3, LA, LA, H2, H2, H4, H4, LK, LK, LQ, LQ, H1, H1, LJ, LJ, LJ, H2, H2, LQ, LQ, SS, LK, LK, H3, H3, H4, LA, LA, H2, H2, H4, H4, LK, LK, H3, H3, H3, LQ, LQ, LJ, LJ, H1, H1, H1, LA, LK, LQ, H2, H2, H2, LJ, LJ, H3, H4, LK, LK, LK, H2, H2, LQ, LJ, H4, H3, LQ, LQ, LJ},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, H3, H3, H3, LA, LA, H1, H1, H2, H2, LQ, LQ, LK, LK, H3, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, SS, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, H2, H2, LJ, LJ, LJ, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LK, LK, SS, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H4, H3, H3, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, SS, LK, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H2, H4, H4, H4, LA, LA, LA, LK, LK, H3, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LA, SS, LK, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LA, LK, LQ, H4, H4, LJ, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, SS, LJ, LJ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, H1, LK, LK, LK, LQ},
			{LK, LK, H4, H4, LQ, LQ, LQ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LJ, LJ, LK, LK, H3, H3, LJ, LJ, H1, H1, H1, H2, LQ, LQ, SS, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H3, H4, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, LK, H4, H4, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, SS, LJ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, H2, LQ, LQ, LQ, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H4, H4, LA, LA, LA, LK, LK, LK, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, SS, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, LJ, LJ, LQ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LQ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H3},
		},
		// FREE_GAME_02
		{
			{LJ, LJ, LJ, H3, H3, LA, LA, H2, H2, H4, H4, LK, LK, LA, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA, H3, H3, LK, LA, LA, LA, LJ, LJ, H2, H2, H3, LQ, LQ, H1, H1, H3, H3, LA, LA, H2, H2, H4, H4, LK, LK, LQ, LQ, H1, H1, LJ, LJ, LJ, H2, H2, LQ, LQ, SS, LK, LK, H3, H3, H4, LA, LA, H2, H2, H4, H4, LK, LK, H3, H3, H3, LQ, LQ, LJ, LJ, H1, H1, LA, LK, LQ, H2, H2, LJ, LJ, H3, H4, LK, LK, LK, H2, H2, LQ, LJ, H4, H3, LQ, LQ, LJ},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, H3, H3, H3, LA, LA, H1, H1, H2, H2, LQ, LQ, LK, LK, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, H2, H2, LJ, LJ, LJ, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LK, LK, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H4, H3, H3, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, LK, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H2, H4, H4, H4, LA, LA, LA, LK, LK, H3, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LA, LK, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LA, LK, LQ, H4, H4, LJ, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, SS, LJ, LJ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, H1, LK, LK, LK, LQ},
			{LK, LK, H4, H4, LQ, LQ, LQ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LJ, LJ, LK, LK, H3, H3, LJ, LJ, H1, H1, H1, H2, LQ, LQ, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H3, H4, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, LK, H4, H4, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, LJ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, H2, LQ, LQ, LQ, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H4, H4, LA, LA, LA, LK, LK, LK, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, SS, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, LJ, LJ, LQ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LQ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H3},
		},
		// FREE_GAME_MYSTERY
		{
			{NN, NN, NN, LA, LA, NN, NN, NN, H2, H2, NN, NN, NN, H1, H1, NN, NN, NN, H4, H4, NN, NN, NN, H3, H3, NN, NN, NN, LK, LK, NN, NN, NN, LQ, LQ, NN, NN, NN, LJ, LJ},
			{NN, NN, NN, LA, LA, NN, NN, NN, H2, H2, NN, NN, NN, H1, H1, NN, NN, NN, H4, H4, NN, NN, NN, H3, H3, NN, NN, NN, LK, LK, NN, NN, NN, LQ, LQ, NN, NN, NN, LJ, LJ},
			{NN, NN, NN, LA, LA, NN, NN, NN, H2, H2, NN, NN, NN, H1, H1, NN, NN, NN, H4, H4, NN, NN, NN, H3, H3, NN, NN, NN, LK, LK, NN, NN, NN, LQ, LQ, NN, NN, NN, LJ, LJ},
			{H2, H2, H2, LA, LA, H1, H1, H1, NN, NN, NN, LQ, LQ, H1, H1, H1, LK, LK, H4, H4, LA, LA, LA, NN, NN, NN, LJ, LJ, LK, LK, H3, H3, H3, LQ, LQ, H4, H4, H4, LJ, LJ},
			{H2, H2, H2, LA, LA, H1, H1, H1, H2, H2, H2, LQ, LQ, H1, H1, H1, LK, LK, H4, H4, LA, LA, LA, NN, NN, NN, LJ, LJ, LK, LK, H2, H2, H2, LQ, LQ, H1, H1, H1, LJ, LJ},
			{H2, H2, H2, LA, LA, H1, H1, H1, H2, H2, H2, LQ, LQ, H1, H1, H1, LK, LK, H4, H4, LA, LA, LA, H3, H3, H3, LJ, LJ, LK, LK, H2, H2, H2, LQ, LQ, H4, H4, H4, LJ, LJ},
		},
		// BUY_FREE_GAME_01
		{
			{LJ, LJ, LJ, H3, H3, SS, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LA, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA, H3, H3, SS, LA, LA, LA, LJ, LJ, H2, H2, H3, LQ, LQ, H1, H1, H3, H3, LA, LA, H2, H2, H4, H4, LK, LK, SS, LQ, LQ, H1, H1, LJ, LJ, LJ, H2, H2, LQ, LQ, SS, LK, LK, H3, H3, H4, LA, LA, H2, H2, H4, H4, LK, LK, H3, H3, H3, LQ, LQ, LJ, LJ, H1, H1, H1, LA, LK, LQ, H2, H2, H2, LJ, LJ, H3, H4, LK, LK, LK, H2, H2, LQ, LJ, H4, H3, LQ, LQ, LJ},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, H3, H3, H3, LA, LA, H1, H1, H2, H2, LQ, LQ, LK, LK, H3, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, SS, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, SS, LQ, LQ, H2, H2, LJ, LJ, LJ, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LK, LK, SS, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H4, H3, H3, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, SS, LK, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H2, H4, H4, H4, LA, LA, LA, LK, LK, H3, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LA, SS, LK, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LA, LK, LQ, H4, H4, LJ, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, SS, LJ, LJ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, H1, LK, LK, LK, LQ},
			{LK, LK, H4, H4, LQ, LQ, LQ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LJ, LJ, LK, LK, H3, H3, LJ, LJ, H1, H1, H1, H2, LQ, LQ, SS, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H3, H4, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, LK, H4, H4, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, SS, LJ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, H2, LQ, LQ, LQ, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H4, H4, LA, LA, LA, LK, LK, LK, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, SS, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, LJ, LJ, LQ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LQ, SS, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, H2, H2, LJ, LJ, H3, H3, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H3},
		},
		// BUY_FREE_GAME_02
		{
			{LK, LK, H4, H4, LQ, LQ, LQ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LJ, LJ, LK, LK, H3, H3, LJ, LJ, H1, H1, H1, H2, LQ, LQ, SS, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H3, H4, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, LK, H4, H4, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, SS, LJ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, H2, LQ, LQ, LQ, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H4, H4, LA, LA, LA, LK, LK, LK, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, SS, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, LJ, LJ, LQ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LQ, SS, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, H2, H2, LJ, LJ, H3, H3, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H3},
			{LJ, LJ, LJ, H3, H3, SS, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LA, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA, H3, H3, SS, LA, LA, LA, LJ, LJ, H2, H2, H3, LQ, LQ, H1, H1, H3, H3, LA, LA, H2, H2, H4, H4, LK, LK, LQ, LQ, LQ, H1, H1, LJ, LJ, LJ, H2, H2, LQ, LQ, LK, LK, LK, H3, H3, H4, LA, LA, H2, H2, H4, H4, LK, LK, H3, H3, H3, LQ, LQ, LJ, LJ, H1, H1, H1, LA, LK, LQ, H2, H2, H2, LJ, LJ, H3, H4, LK, LK, LK, H2, H2, LQ, LJ, H4, H3, LQ, LQ, LJ},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, H3, H3, H3, LA, LA, H1, H1, H2, H2, LQ, LQ, LK, LK, H3, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, SS, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, LJ, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LK, LK, LK, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H4, H3, H3, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, SS, LK, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H2, H4, H4, H4, LA, LA, LA, LK, LK, H3, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LA, LK, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LA, LK, LQ, H4, H4, LJ, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, SS, LJ, LJ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, H1, LK, LK, LK, LQ},
		},
		// BUY_FREE_GAME_03
		{
			{LK, LK, H4, H4, LQ, LQ, LQ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LJ, LJ, LK, LK, H3, H3, LJ, LJ, H1, H1, H1, H2, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, H3, H4, LA, LA, H1, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H3, H4, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, LK, H4, H4, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, SS, LJ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, H2, LQ, LQ, LQ, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H4, H4, LA, LA, LA, LK, LK, LK, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LK, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LK, LQ, LQ, H4, H4, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, LJ, LJ, LQ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, LK, LK, LK},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, LJ, H3, H3, LA, LA, H1, H1, H2, H2, LK, LK, LK, H3, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, H2, H2, LJ, LJ, H3, H3, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LA, LA, LA, H2, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H3},
			{LJ, LJ, LJ, H3, H3, H3, LA, LA, H2, H2, H2, H4, H4, H4, LK, LK, LA, H1, H1, LQ, LQ, LQ, H4, H4, LA, LA, H3, H3, SS, LA, LA, LA, LJ, LJ, H2, H2, H3, LQ, LQ, H1, H1, H3, H3, LA, LA, H2, H2, H4, H4, LK, LK, LQ, LQ, LQ, H1, H1, LJ, LJ, LJ, H2, H2, LQ, LQ, LK, LK, LK, H3, H3, H4, LA, LA, H2, H2, H4, H4, LK, LK, H3, H3, H3, LQ, LQ, LJ, LJ, H1, H1, H1, LA, LK, LQ, H2, H2, H2, LJ, LJ, H3, H4, LK, LK, LK, H2, H2, LQ, LJ, H4, H3, LQ, LQ, LJ},
			{LK, LK, H4, H4, LQ, LQ, LJ, LJ, H3, H3, H3, LA, LA, H1, H1, H2, H2, LQ, LQ, LK, LK, H3, H3, H3, LJ, LJ, H1, H1, H2, LQ, LQ, LQ, LJ, LJ, LJ, H4, H4, H2, LA, LA, LA, H1, H1, H1, LQ, LQ, H3, H4, LA, LK, LK, LQ, LQ, LQ, H2, H2, LJ, LJ, LJ, H3, H4, LA, LA, H1, H1, LJ, LJ, H2, H3, LK, LQ, H4, H4, LK, LK, LK, LA, LA, LA, H2, H2, LK, LQ, H4, H4, LJ, LJ, H1, H2, LA, LA, SS, LQ, LQ, H3, H3, H4, H4, LK, LK, H1, H1, H1, LJ, LJ, H2, H3, H1, LK, LK, H4, H4, H4, H3, H3, H3},
			{LA, LA, H3, H3, LK, LQ, H4, H4, H4, LA, LK, H2, H2, LQ, LQ, LQ, LJ, LJ, H1, H1, H1, LK, LJ, H3, H3, H4, LA, LK, H2, H2, LQ, LQ, LQ, LK, LK, LK, H1, H1, H3, H3, LJ, LJ, LJ, H2, H2, H2, H4, H4, H4, LA, LA, LA, LK, LK, H3, H3, H3, LQ, LQ, H1, H1, H2, H2, LA, LA, LA, LK, LK, LK, H3, H4, LQ, LQ, H2, H2, H2, LA, LA, LK, LQ, H4, H4, LJ, LJ, LJ, H1, H2, LK, LK, H3, H3, LA, LA, H4, H4, SS, LJ, LJ, LQ, H2, H2, H3, H3, LA, LA, LA, H1, H1, H1, LK, LK, LK, LQ},
		},
	},
}

// FGReelGroupWT 免費遊戲轉輪群組權重
var FGReelGroupWT = []map[int][]uint{
	// RTP0 上線數值
	{
		Common.BUY_NONE:             {500, 200, 300},
		Common.BUY_EXTRA_BET:        {500, 200, 300},
		Common.BUY_FREE_SPINS:       {0, 0, 0},
		Common.BUY_SUPER_FREE_SPINS: {0, 0, 0},
	},
}

// MysteryList Mystery 轴图标
var MysteryList = []Symbol{H1, H2, H3, H4}

// MysteryWT Mystery 轴图标权重
var MysteryWT = []uint{20, 30, 40, 40}

// MGSSMulti  主游戏铜钱倍数列表
var MGSSMultiList = []Symbol{SS2, SS3, SS5}

// MGSSMultiWT 主游戏铜钱倍数权重
var MGSSMultiWT = []uint{100, 60, 10}

// MGFeatureSSList 主游戏天降横财铜钱个数
var MGFeatureSSList = []uint{1, 2}

// MGFeatureSSWT 主游戏天降横财铜钱个数权重
var MGFeatureSSWT = []uint{200, 100}

// FGIndexList 免費遊戲配套索引列表
var FGIndexList = []int{FREE_INDEX_1, FREE_INDEX_2, FREE_INDEX_3}

// FGIndexWT 免費遊戲配套索引權重
var FGIndexWT = []uint{500, 200, 300}

// FGBuyFreeIndexWT 免費遊戲buy配套索引權重
var FGBuyFreeIndexWT = []uint{543, 307, 150}

// FGBuySuperIndexWT 免費遊戲buy配套索引權重
var FGBuySuperIndexWT = []uint{670, 230, 100}

// FGSSMultiList 免費遊戲銅錢倍數列表
var FGSSMultiList = []Symbol{SS2, SS3, SS5}

// FGSSMultiWT 免費遊戲銅錢倍數權重
var FGSSMultiWT = []uint{104, 41, 10}
