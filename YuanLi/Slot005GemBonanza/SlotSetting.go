package Slot005GemBonanza

import "Force/GameServer/Common"

// PROB_VERSION 機率版號
const PROB_VERSION = "slot005-97.00-v1.0.0"

// 常數定義
const (
	PAYLINE_TOTAL = 10 // 中獎線數

	SLOT_COL = 5 // 輪盤組數
	SLOT_ROW = 3 // 輪盤列數

	DEFAULT_RTP = 0 // 預設 RTP 編號
	RTP_TOTAL   = 1 // 共幾組 RTP

	WIN_SCATTER_COUNT     = 3 // 中 3 個 Scatter 可進免費遊戲
	MAX_WIN_SCATTER_COUNT = 5 // 中 5 個 Scatter 以上賠率相同、觸發免費遊戲次數相同
	MAX_COVER_WILD_COUNT  = 2 // 免費遊戲最大覆蓋 Wild 數量

	FREE_GAME_CONDITION_WILD_COUNT = 1 // 免費遊戲表演條件: 前三輪 Wild 獎圖數量
	FREE_GAME_CONDITION_H5_COUNT   = 3 // 免費遊戲表演條件: H5 獎圖數量

	MAX_ODDS = 5000 * PAYLINE_TOTAL // 最大賠率
)

// 測試指令資料 index
const (
	DEBUG_INDEX_GROUP_INDEX   = iota // 主遊戲轉輪群組 index (0~5)
	DEBUG_INDEX_REEL_INDEX_01        // 停輪位置
	DEBUG_INDEX_REEL_INDEX_02
	DEBUG_INDEX_REEL_INDEX_03
	DEBUG_INDEX_REEL_INDEX_04
	DEBUG_INDEX_REEL_INDEX_05
	DEBUG_INDEX_COVER_SCATTER_COUNT // 主遊戲覆蓋 Scatter 獎圖數量 (3~5)
	DEBUG_INDEX_FREE_GAME_TYPE      // 免費遊戲類型 (即免費遊戲的轉輪群組 index, 0~1)
	DEBUG_INDEX_H5_SCORE_01         // H5 獎圖分數編號 (0~12)
	DEBUG_INDEX_H5_SCORE_02
	DEBUG_INDEX_H5_SCORE_03
	DEBUG_INDEX_H5_SCORE_04
	DEBUG_INDEX_H5_SCORE_05
	DEBUG_INDEX_H5_SCORE_06
	DEBUG_INDEX_H5_SCORE_07
	DEBUG_INDEX_H5_SCORE_08
	DEBUG_INDEX_H5_SCORE_09
	DEBUG_INDEX_H5_SCORE_10
	DEBUG_INDEX_H5_SCORE_11
	DEBUG_INDEX_H5_SCORE_12
	DEBUG_INDEX_H5_SCORE_13
	DEBUG_INDEX_H5_SCORE_14
	DEBUG_INDEX_H5_SCORE_15
	DEBUG_INDEX_COVER_WILD_COUNT // 免費遊戲覆蓋 Wild 數量 (0~2)
	DEBUG_INDEX_PERFORMANCE_TYPE // 表演類型 (主遊戲: 0~1，免費遊戲: 0~3)
)

// BetRatio 投注額比例
var BetRatio = map[int]float64{
	Common.BUY_NONE:             1,
	Common.BUY_EXTRA_BET:        1.5,
	Common.BUY_FREE_SPINS:       100,
	Common.BUY_SUPER_FREE_SPINS: 270,
}

// Symbols
const (
	NN Symbol = iota - 1
	WW
	H1
	H2
	H3
	H4
	H5
	LA
	LK
	LQ
	LJ
	LT
	SS
	MAX_SYMBOL
)

// LineIndexArray 中獎線圖定義
//	R1  R2  R3  R4  ...
//	0   0   0   0   ...
//	1   1   1   1   ...
//	2   2   2   2   ...
//	...
var LineIndexArray = [][]int{
	{1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0},
	{2, 2, 2, 2, 2},
	{1, 0, 0, 0, 1},
	{1, 2, 2, 2, 1},
	{2, 1, 0, 1, 2},
	{0, 1, 2, 1, 0},
	{2, 2, 1, 0, 0},
	{0, 0, 1, 2, 2},
	{2, 1, 1, 1, 0},
}

// SymbolOdds 獎圖賠率 [獎圖][數量]賠率
var SymbolOdds = [][]int{
	{0, 0, 5, 50, 200, 2000}, // WW
	{0, 0, 5, 50, 200, 2000}, // H1
	{0, 0, 0, 30, 150, 1000}, // H2
	{0, 0, 0, 20, 100, 500},  // H3
	{0, 0, 0, 20, 100, 500},  // H4
	{0, 0, 0, 10, 50, 200},   // H5
	{0, 0, 0, 2, 25, 100},    // LA
	{0, 0, 0, 2, 25, 100},    // LK
	{0, 0, 0, 2, 10, 50},     // LQ
	{0, 0, 0, 2, 10, 50},     // LJ
	{0, 0, 0, 2, 10, 50},     // LT
	{0, 0, 0, 0, 0, 0, 0},    // SS
}

// WinSymbolCount 中獎獎圖連線個數 (達到才算中獎) [獎圖]個數
var WinSymbolCount = [MAX_SYMBOL]int{2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, SLOT_COL + 1}

// FreeGameType 免費遊戲類型
const (
	FREE_GAME_01 int = iota
	FREE_GAME_02
)

// MGReelGroup 主遊戲轉輪群組
var MGReelGroup = [][][][]Symbol{
	// RTP0 上線數值
	{
		// MAIN_GAME_01
		{
			{H5, H5, H5, H5, H5, LJ, H4, LT, H5, LJ, H3, LA, SS, LQ, LK, H1, LT, H2, LJ, H3, H4, H5, H5, H2, H4, LK, H3, H4, H2, H5, H5, LA, LQ, H4, LJ, H5, LK, LA, LJ, H4, LK, H3, LQ, LT, LA, H2, LK, H4, LT, SS, LA, LT, H1, LQ, LK, LA, LQ},
			{H5, H5, H5, H5, H5, LK, LQ, LJ, LK, H5, LT, H4, LJ, LQ, H5, H5, H4, LA, H1, LQ, H4, H2, H5, H5, H3, LA, LJ, LT, H2, H1, H4, LQ, LT, H3, LJ, H4, LA, LK, H4, LJ, H3, LT, SS, LA, LK, H2, LA, H3, LT, H5, LA, H2, H4, LQ},
			{H5, H5, H5, H5, H5, LK, LJ, H5, LA, LQ, H5, LK, LA, LT, H5, H5, H4, H2, LQ, H3, H4, H5, H5, LK, H1, LT, H2, H4, H3, LJ, H4, LA, LJ, LQ, LK, H2, LT, LK, H2, LQ, H4, LA, LJ, H4, LT, H3, LA, SS, LQ, LA, H1, H4, LQ, LT, LK, LA, LJ, H3},
			{H5, H5, H5, H5, H5, LJ, LA, H3, H5, LQ, H4, LK, LA, LJ, SS, LT, H4, LJ, H5, H5, LQ, H4, LK, H2, H1, H5, H5, LQ, LA, SS, LT, H4, H3, LA, H5, LK, H4, LJ, H3, LK, H2, LT, H4, LK, LQ, H2, LJ, H4, LA, H2, LT, H1, H3, LA, LK, LJ, LT},
			{H5, H5, H5, H5, H5, H4, H2, LA, H5, H4, LQ, H3, H5, H5, LK, H2, H4, LK, LJ, LT, LQ, H4, LA, H3, LK, LJ, H5, H5, LQ, H4, H3, LQ, SS, LT, H4, H3, LJ, H1, LA, H2, LT, H4, LJ, H1, LA, LK, LQ, LJ, SS, LT, H5, LA, LK, H2, LT, LK},
		},
		// MAIN_GAME_02
		{
			{H5, H5, H5, LQ, LT, LJ, H4, LT, H5, LJ, H3, LA, SS, LQ, LK, H1, LT, H2, LJ, H3, H4, H5, H5, H2, H4, LK, H3, H4, H2, H5, H5, LA, SS, LQ, H4, H5, LJ, LK, SS, LA, LJ, H4, LK, H3, LQ, LT, LA, H2, LK, H4, LT, SS, LA, LT, H1, LQ, LK, SS, LA, LQ},
			{H5, H5, H5, H5, H5, LK, LQ, SS, LJ, LK, H5, LT, H4, LJ, SS, LQ, H5, H5, H4, LA, H1, LQ, H4, H2, H5, H5, H3, LA, SS, LJ, LT, H2, H1, H4, LQ, SS, LT, H3, LJ, H4, LA, H5, LK, H4, LJ, H3, LT, LA, SS, LK, H2, LA, H3, LT, LA, H2, H4, LQ},
			{H5, H5, H5, H5, H5, LK, LJ, H5, LA, LQ, H5, LK, LA, LT, H5, H5, H4, H2, LQ, H3, H4, H5, H5, LK, H1, LT, H2, H4, H3, LJ, H4, LA, LJ, LQ, LK, H2, LT, LK, H2, LQ, H4, LA, LJ, H4, LT, LA, H3, LQ, LA, H1, H4, LQ, LT, LK, LA, LJ, H3},
			{H5, H5, H5, H5, H5, LJ, LA, H3, H5, LQ, H4, LK, LA, SS, LJ, LT, H4, LJ, H5, H5, LQ, H4, LK, H2, H1, H5, H5, LQ, LA, LT, H4, H3, LA, SS, LK, H4, LJ, H3, LK, H2, LT, H4, LK, LQ, H2, LJ, H4, LA, H2, LT, H1, H3, LA, H5, LK, LJ, LT},
			{H5, H5, H5, H5, H5, H4, H2, LA, H5, H4, LQ, H3, H5, H5, LK, H2, LJ, H4, LK, LT, LQ, H4, LA, H3, LK, LJ, H5, H5, LQ, H4, H3, LQ, H5, LT, H4, H3, LJ, H1, LA, H2, LT, H4, LJ, H1, LA, LK, LQ, LJ, LT, LA, LK, H2, LT, LK},
		},
		// MAIN_GAME_03
		{
			{H5, H5, H5, H5, H5, LJ, H4, LT, H5, LJ, H3, LA, SS, LQ, LK, H1, LT, H2, LJ, H3, H4, H5, H5, H2, H4, LK, H3, H4, H2, H5, H5, LA, SS, LQ, H4, LJ, LK, SS, LA, LJ, H4, LK, H3, LQ, SS, LT, LA, H2, LK, H4, LT, SS, LA, LT, H1, LQ, H5, LK, SS, LA, LQ},
			{LK, LT, H5, H5, H5, LK, LQ, LJ, LK, H5, LT, H4, LJ, LQ, H5, H5, H4, LA, H1, LQ, H4, H2, H5, H5, H3, LA, LJ, LT, H2, H1, H4, LQ, LT, H3, LJ, H4, LA, LK, H4, LJ, H3, LT, H5, LA, LK, H2, LA, H3, LT, LA, H2, H4, LQ},
			{H5, H5, H5, H5, H5, LK, LJ, H5, LA, SS, LQ, H5, LK, LA, LT, H5, H5, H4, H2, LQ, H3, H4, H5, H5, LK, H1, LT, H2, H4, H3, SS, LJ, H4, LA, SS, LJ, LQ, SS, LK, H2, LT, SS, LK, H2, LQ, H4, LA, LJ, H4, LT, H3, LA, SS, LQ, LA, H1, H4, LQ, LT, LK, LA, LJ, H3},
			{H5, H5, H5, H5, H5, LJ, LA, H3, H5, LQ, H4, LK, LA, LJ, LT, H4, LJ, H5, H5, LQ, H4, LK, H2, H1, H5, H5, LQ, LA, LT, H4, H3, LA, H5, LK, H4, LJ, H3, LK, H2, LT, H4, LK, LQ, H2, LJ, H4, LA, H2, LT, H1, H3, LA, LK, LJ, LT},
			{H5, H5, H5, H5, H5, H4, H2, LA, H5, H4, LQ, H3, H5, H5, LK, H2, LJ, H4, LK, SS, LT, LQ, H4, LA, H3, LK, LJ, H5, H5, LQ, H4, H3, LQ, LT, H4, H3, LJ, H1, LA, H2, LT, H4, LJ, H1, LA, LK, LQ, SS, LJ, LT, H5, LA, LK, H2, LT, LK},
		},
		// MAIN_GAME_04
		{
			{H5, H5, H5, H5, H5, LJ, H4, LT, H5, LJ, H3, LA, LQ, LK, H1, LT, H2, LJ, H3, H4, H5, H5, H2, H4, LK, H3, H4, H2, H5, H5, LA, LQ, H4, LJ, H5, LK, LA, LJ, H4, LK, H3, LQ, LT, LA, H2, LK, LT, H4, LA, LT, H1, LQ, LK, LA, LQ},
			{H5, H5, H5, H5, H5, LK, LQ, SS, LJ, LK, H5, LT, H4, LJ, SS, LQ, H5, H5, H4, LA, H1, LQ, H4, H2, H5, H5, H3, LA, SS, LJ, LT, H2, H1, H4, LQ, SS, LT, H3, LJ, H4, LA, H5, LK, H4, LJ, H3, LT, LA, SS, LK, H2, LA, H3, LT, LA, H2, H4, LQ},
			{LJ, LT, H5, H5, H5, LK, LJ, H5, LA, SS, LQ, H5, LK, LA, LT, H5, H5, H4, H2, LQ, H3, H4, H5, H5, LK, H1, LT, H2, H4, H3, SS, LJ, H4, LA, SS, LJ, LQ, SS, LK, H2, LT, LK, H2, LQ, H4, LA, LJ, H4, LT, H3, LA, SS, LQ, LA, H1, H4, LQ, LT, LK, LA, LJ, H3},
			{H5, H5, H5, H5, H5, LJ, LA, H3, H5, LQ, H4, LK, LA, LJ, LT, H4, LJ, H5, H5, LQ, H4, LK, H2, H1, H5, H5, LQ, LA, LT, H4, H3, LA, H5, LK, H4, LJ, H3, LK, H2, LT, H4, LK, LQ, H2, LJ, H4, LA, H2, LT, H1, H3, LA, LK, LJ, LT},
			{H5, H5, H5, H5, H5, H4, H2, LA, H5, H4, LQ, H3, H5, H5, LK, H2, LJ, H4, LK, SS, LT, LQ, H4, LA, H3, LK, LJ, H5, H5, LQ, H4, H3, LQ, LT, H4, H3, LJ, H1, LA, H2, LT, H4, LJ, H1, LA, LK, SS, LQ, LJ, H5, LT, LA, LK, H2, LT, LK},
		},
		// MAIN_GAME_05
		{
			{H5, H5, H5, LT, H3, LJ, H1, LQ, H4, LK, H3, LA, H1, LK, H2, LQ, H3, LA, H1},
			{H5, H5, H5, LT, H3, LJ, H4, LQ, H2, LK, H4, LA, H1, LK, LQ, H3, H2, H4, LA},
			{H5, H5, H5, LJ, H2, LT, H1, LK, H4, LQ, H3, LA, H1, LQ, LK, H1, H3, LA, H4},
			{H5, H5, H5, LJ, H1, LT, H2, LK, H4, LQ, H3, LA, H1, LK, LQ, H3, LJ, H2, H4, LA},
			{H5, H5, H5, LT, LA, H2, H4, LQ, H3, LK, H1, LJ, H2, LK, H3, LJ, LQ, H1, LA},
		},
		// MAIN_GAME_06
		{
			{H5, H5, H5, H5, H5, LJ, H4, LT, H5, LJ, H3, LA, LQ, LK, H1, LT, H2, LJ, H3, H4, H5, H5, H2, H4, LK, H3, H4, H2, H5, H5, LA, LQ, H4, LJ, H5, LK, LA, LJ, H4, LK, H3, LQ, LT, LA, H2, LK, H4, LT, LA, LT, H1, LQ, LK, LA, LQ},
			{H5, H5, H5, H5, H5, LK, LQ, LJ, LK, H5, LT, H4, LJ, LQ, H5, H5, H4, LA, H1, LQ, H4, H2, H5, H5, H3, LA, LJ, LT, H2, H1, H4, LQ, LT, H3, LJ, H4, LA, LK, H4, LJ, H3, LT, LA, LK, H2, LA, H3, LT, H5, LA, H2, H4, LQ},
			{H5, H5, H5, H5, H5, LK, LJ, H5, LA, LQ, H5, LK, LA, LT, H5, H5, H4, H2, LQ, H3, H4, H5, H5, LK, H1, LT, H2, H4, H3, LJ, H4, LA, LJ, LQ, LK, H2, LT, LK, H2, LQ, H4, LA, LJ, H4, LT, H3, LA, LQ, LA, H1, H4, LQ, LT, LK, LA, LJ, H3},
			{H5, H5, H5, H5, H5, LJ, LA, H3, H5, LQ, H4, LK, LA, LJ, LT, H4, LJ, H5, H5, LQ, H4, LK, H2, H1, H5, H5, LQ, LA, LT, H4, H3, LA, H5, LK, H4, LJ, H3, LK, H2, LT, H4, LK, LQ, H2, LJ, H4, LA, H2, LT, H1, H3, LA, LK, LJ, LT},
			{H5, H5, H5, H5, H5, H4, H2, LA, H5, H4, LQ, H3, H5, H5, LK, H2, H4, LK, LJ, LT, LQ, H4, LA, H3, LK, LJ, H5, H5, LQ, H4, H3, LQ, LT, H4, H3, LJ, H1, LA, H2, LT, H4, LJ, H1, LA, LK, LQ, LJ, LT, H5, LA, LK, H2, LT, LK},
		},
	},
}

// MGReelGroupWT 主遊戲轉輪群組權重
var MGReelGroupWT = []map[int][]uint{
	// RTP0 上線數值
	{
		Common.BUY_NONE:             {86, 229, 229, 229, 749, 0},
		Common.BUY_EXTRA_BET:        {86, 229, 229, 229, 749, 0},
		Common.BUY_FREE_SPINS:       {0, 0, 0, 0, 0, 1},
		Common.BUY_SUPER_FREE_SPINS: {0, 0, 0, 0, 0, 1},
	},
}

// FGReelGroup 免費遊戲轉輪群組
var FGReelGroup = [][][][]Symbol{
	// RTP0 上線數值
	{
		// FREE_GAME_01
		{
			{H5, H5, H4, LA, LQ, H4, H3, LT, H5, H5, LQ, H1, LT, LA, LQ, LJ, LA, LQ, LT, H2, LK, LQ, LT, LA, H2, LT, LQ, LA, LJ, LT, LA, LK, LT, LA, LQ},
			{H5, H5, H5, H3, LK, H5, H4, H1, LK, LJ, LT, H2, LK, H5, H5, LJ, LA, LK, LJ, LQ, H1, LK, H5, H5, LJ, H3, LK, LJ, LA, LT, LJ, LK, LQ, LJ},
			{H5, H5, H5, H4, H2, H3, H5, LA, H2, H4, H3, H5, LK, H1, H4, LJ, LT, H3, H2, H4, LT, LA, H2, LQ, LK, H3, LQ, LJ},
			{H5, H5, H4, H2, LA, H5, LT, H4, LQ, H1, H3, H4, LK, H1, H3, H2, H5, LJ, H3},
			{H5, H5, H2, H3, H5, H4, LA, LT, H5, H4, H3, H1, H2, H5, LJ, LK, LQ, LA, LT, LJ, LK, LQ},
		},
		// FREE_GAME_02
		{
			{H5, H5, H5, NN, NN, NN, NN, NN, NN, H5, H5, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN, NN},
			{H5, H5, H5, NN, NN, H5, H5, NN, NN, NN, H5, NN, NN, NN, NN, H5, H5, NN, NN, NN, NN, NN, NN, NN, H5, H5, NN, NN, NN, NN, NN, H5, NN, NN, NN, NN, NN},
			{H5, H5, H5, NN, NN, NN, H5, NN, NN, NN, NN, H5, H5, NN, NN, NN, NN, H5, NN, NN, NN, NN, NN, H5, NN, NN, NN, NN, NN, NN, NN},
			{H5, H5, H5, NN, NN, NN, H5, NN, NN, H5, NN, NN, NN, NN, H5, NN, NN, NN, NN, H5, H5, NN, NN},
			{H5, H5, H5, NN, NN, H5, NN, NN, NN, H5, H5, NN, NN, NN, NN, H5, NN, NN, NN, H5, NN, NN, NN, NN, NN},
		},
	},
}

// FGReelGroupWT 免費遊戲轉輪群組權重
var FGReelGroupWT = []map[int][]uint{
	// RTP0 上線數值
	{
		Common.BUY_NONE:             {11, 1},
		Common.BUY_EXTRA_BET:        {11, 1},
		Common.BUY_FREE_SPINS:       {9, 1},
		Common.BUY_SUPER_FREE_SPINS: {0, 1},
	},
}

// MGCoverScatterCountWT 覆蓋 Scatter 數量權重
var MGCoverScatterCountWT = []map[int]map[int]uint{
	// RTP0 上線數值
	{
		Common.BUY_NONE:             {5: 0, 4: 0, 3: 0, 0: 1},
		Common.BUY_EXTRA_BET:        {5: 0, 4: 0, 3: 43, 0: 5314},
		Common.BUY_FREE_SPINS:       {5: 15, 4: 51, 3: 429, 0: 0},
		Common.BUY_SUPER_FREE_SPINS: {5: 1, 4: 6, 3: 122, 0: 0},
	},
}

// FGInitSpinCount 免費遊戲初始場次 (index: Scatter 數量)
var FGInitSpinCount = [MAX_WIN_SCATTER_COUNT + 1]int{0, 0, 0, 10, 15, 20}

// H5ScoreList H5 獎圖分數列表
var H5ScoreList = []int{2, 5, 10, 15, 20, 25, 50, 100, 200, 500, 1666, 2500, 5000}

// MGH5ScoreWT 主遊戲 H5 獎圖分數權重
var MGH5ScoreWT = []map[int]uint{
	// RTP0 上線數值
	{2: 4000, 5: 3800, 10: 2500, 15: 1000, 20: 500, 25: 300, 50: 100, 100: 10, 200: 10, 500: 5, 1666: 2, 2500: 3, 5000: 5},
}

// FGH5ScoreWT 免費遊戲 H5 獎圖分數權重
var FGH5ScoreWT = [][]map[int]uint{
	// RTP0 上線數值
	{
		// FREE_GAME_01
		{2: 4000, 5: 3800, 10: 2500, 15: 1000, 20: 500, 25: 300, 50: 100, 100: 10, 200: 10, 500: 5, 1666: 2, 2500: 3, 5000: 5},
		// FREE_GAME_02
		{2: 0, 5: 1250, 10: 1250, 15: 1250, 20: 1250, 25: 1000, 50: 125, 100: 50, 200: 20, 500: 5, 1666: 2, 2500: 2, 5000: 3},
	},
}

// Wild Table (控制 Wild 個數的權重表)
const (
	WILD_TABLE_01 int = iota
	WILD_TABLE_02
	WILD_TABLE_03
)

// WildTableList Wild Table 列表
var WildTableList = []int{WILD_TABLE_01, WILD_TABLE_02, WILD_TABLE_03}

// FGWildTableMap Wild Table 個數配置表 [RTP][FREE_GAME_01 or 02][ScatterCount][Stage]Wild Table 個數列表
var FGWildTableMap = [][]map[int][][]int{
	// RTP0 上線數值
	{
		// FREE_GAME_01
		{
			3: {{5, 4, 1}, {6, 4, 0}, {9, 1, 0}, {9, 0, 1}},
			4: {{10, 4, 1}, {5, 5, 0}, {7, 3, 0}, {5, 4, 1}},
			5: {{15, 4, 1}, {6, 4, 0}, {7, 3, 0}, {3, 6, 1}},
		},
		// FREE_GAME_02
		{
			3: {{8, 1, 1}, {8, 2, 0}, {8, 2, 0}, {9, 0, 1}},
			4: {{10, 5, 0}, {7, 3, 0}, {7, 3, 0}, {5, 4, 1}},
			5: {{15, 5, 0}, {7, 3, 0}, {7, 3, 0}, {5, 4, 1}},
		},
	},
}

// WildCountList Wild 個數列表
var WildCountList = []int{0, 1, 2}

// FGWildCountWT Wild 個數權重 [RTP][WildTable][H5個數]權重列表
var FGWildCountWT = []map[int][][]uint{
	// RTP0 上線數值
	{
		WILD_TABLE_01: {
			{1000, 0, 0},
			{950, 50, 0},
			{800, 190, 10},
			{800, 190, 10},
		},
		WILD_TABLE_02: {
			{1000, 0, 0},
			{900, 100, 0},
			{670, 320, 10},
			{670, 300, 30},
		},
		WILD_TABLE_03: {
			{1000, 0, 0},
			{0, 1000, 0},
			{0, 1000, 0},
			{0, 1000, 0},
		},
	},
}

// 主遊戲 PerformanceType 表演類型
const (
	NONE            int = iota // 無表演
	REPLACE_SCATTER            // 隨機替換一個 Scatter
)

// 免費遊戲 PerformanceType 表演類型
const (
	REPLACE_WILD   int = iota + 1 // 出現 Wild 的整輪替換
	REPLACE_H5                    // 替換全部 H5 獎圖
	REPLACE_OTHERS                // 替換 Wild 以外的獎圖
)

// CoverSymbolList 覆蓋獎圖列表
var CoverSymbolList = []Symbol{H2, H3, H4, LA, LK, LQ, LJ, LT}

// MGPerformanceList 主遊戲表演類型列表
var MGPerformanceList = []int{NONE, REPLACE_SCATTER}

// MGPerformanceWT 主遊戲表演類型權重
var MGPerformanceWT = [][]uint{
	// RTP0 上線數值
	{150, 30},
}

// FGPerformanceList 表演類型列表
var FGPerformanceList = []int{NONE, REPLACE_WILD, REPLACE_H5, REPLACE_OTHERS}

// FGPerformanceWT 免費遊戲表演類型權重
var FGPerformanceWT = [][]uint{
	// RTP0 上線數值
	{16, 2, 2, 0},
}

// FGStageWildCount 免費遊戲進下一階所需的 Wild 個數
var FGStageWildCount = []int{4, 8, 12, 9999}

// FGStageMultiplier 免費遊戲各階段的乘倍
var FGStageMultiplier = []int{1, 2, 3, 10}

// FGAddSpinCount 免費遊戲再觸發增加場次 (index: Stage)
var FGAddSpinCount = []int{0, 10, 10, 10}
