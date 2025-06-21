package Slot005GemBonanza

import (
	"Force/GameServer/Common"
	"Force/GameServer/Utils"
	"fmt"
	"math/rand/v2"
)

type SlotProb struct {
	MGReelGroupRouletteMaps    [RTP_TOTAL]map[int]Utils.Roulette[int] // 主遊戲轉輪群組 Roulette Map
	FGReelGroupRouletteMaps    [RTP_TOTAL]map[int]Utils.Roulette[int] // 免費遊戲轉輪群組 Roulette Map
	MGCoverScatterRouletteMaps [RTP_TOTAL]map[int]Utils.Roulette[int] // 主遊戲覆蓋 Scatter 數量 Roulette Map
	MGH5ScoreRoulette          [RTP_TOTAL]Utils.Roulette[int]         // 主遊戲 H5 獎圖分數 Roulette
	FGH5ScoreRouletteList      [RTP_TOTAL][]Utils.Roulette[int]       // 免費遊戲 H5 獎圖分數 Roulette List
}

// Init 初始化
func (slot *SlotProb) Init() {
	// 建立 Roulette 資料
	slot.SetRTPRoulette()
}

// Run 進行遊戲
func (slot *SlotProb) Run(rtp int, lineBet int, result *SlotResult, debugCmdList []DebugCmd) (gameMode int) {
	gameMode = Common.GAME_MODE_NORMAL

	var buyType = result.BuyType
	var mainWin, freeWin uint64 = 0, 0
	var isMaxWin = false                    // 是否達到最大贏分
	var maxWin = uint64(lineBet * MAX_ODDS) // 最大贏分
	var reSpin = true
	var isPerformance = true
	var lastMGSymbol GameSymbol
	var lastReelIndex []int
	var lastScoreArray [][]int
	for reSpin {
		var tumble = TumbleResult{}
		// 第一轉盤面亂數產生
		if len(result.MGTumbleList) == 0 {
			var debugCmd *DebugCmd
			if len(debugCmdList) > 0 {
				debugCmd = &debugCmdList[0]
			}
			tumble.TumbleSymbol, result.MGGroupIndex, tumble.ReelIndex, tumble.H5ScoreArray, result.Code = slot.RandMGSymbol(rtp, buyType, lineBet, debugCmd)
			// 有發生錯誤則直接結束
			if result.Code != Common.ERROR_CODE_OK {
				return
			}
		} else {
			// 若符合重轉條件，則產生重轉盤面
			tumble.TumbleSymbol, tumble.ReelIndex, tumble.H5ScoreArray = slot.RespinSymbol(lastMGSymbol, rtp, result.MGGroupIndex, lineBet, lastReelIndex, lastScoreArray)
		}

		// 計算該盤面贏分
		tumble.Win, tumble.LineSymbol, tumble.LineCount, tumble.LineWin = slot.CalculateWin(lineBet, tumble.TumbleSymbol)

		// 紀錄結果
		mainWin += tumble.Win
		lastMGSymbol = tumble.TumbleSymbol
		lastReelIndex = tumble.ReelIndex
		lastScoreArray = tumble.H5ScoreArray
		result.MGTumbleList = append(result.MGTumbleList, tumble)

		// 達到最大贏分則結束
		mainWin, isMaxWin = Utils.CheckMaxWin(mainWin, maxWin)
		if isMaxWin {
			reSpin = false
			break // 直接中斷，不重轉
			// TODO: Log?
		}

		// 檢查是否符合重轉條件
		reSpin = slot.IsRespin(tumble.TumbleSymbol)
		// 有重轉就不表演
		if reSpin {
			isPerformance = false
		}
	}

	// 判斷是否進免費遊戲 (需未達最大贏分)
	var isWinFree = false
	var scatterCount = slot.GetSymbolCount(lastMGSymbol, SS)
	if scatterCount >= WIN_SCATTER_COUNT && !isMaxWin {
		if scatterCount >= MAX_WIN_SCATTER_COUNT {
			scatterCount = MAX_WIN_SCATTER_COUNT
		}
		isWinFree = true
		// 設定免費遊戲場次
		result.FGSpinCount = FGInitSpinCount[scatterCount]
	}

	// 儲存結果
	result.MainWin = mainWin

	// 免費遊戲相關處理 (需未達最大贏分)
	if isWinFree && !isMaxWin {
		// 決定表演類型
		if isPerformance {
			result.MGPerformanceType = Utils.RandChoiceByWeight(MGPerformanceList, MGPerformanceWT[rtp])
			// 處理測試指令
			if len(debugCmdList) > 0 && len(debugCmdList[0].DebugData) > DEBUG_INDEX_PERFORMANCE_TYPE {
				// 取得表演類型，並檢查是否合法
				var debugPerformanceType = debugCmdList[0].DebugData[DEBUG_INDEX_PERFORMANCE_TYPE]
				if 0 <= debugPerformanceType && debugPerformanceType <= REPLACE_SCATTER {
					result.MGPerformanceType = debugPerformanceType
				}
			}
			// 取得表演盤面
			if result.MGPerformanceType == REPLACE_SCATTER {
				result.MGPerformanceSymbol = slot.ReplaceScatter(lastMGSymbol)
			}
		}
		// 進行免費遊戲
		gameMode |= Common.GAME_MODE_FREE
		freeWin = slot.RunFreeGame(rtp, lineBet, scatterCount, result, debugCmdList)
	}

	result.TotalWin = result.MainWin + freeWin
	result.GameMode = gameMode
	return
}

// RunFreeGame 進行免費遊戲
func (slot *SlotProb) RunFreeGame(rtp int, lineBet int, scatterCount int, result *SlotResult, debugCmdList []DebugCmd) uint64 {
	var buyType = result.BuyType
	var freeWin uint64 = 0
	var isMaxWin = false                                       // 是否達到最大贏分
	var maxFreeWin = uint64(lineBet*MAX_ODDS) - result.MainWin // 免費遊戲的最大贏分
	var spinCount = 0                                          // 免費遊戲次數
	var stage = 0                                              // 免費遊戲階段
	var cumWildCount = 0                                       // 免費遊戲累積 Wild 個數
	var multiplier = FGStageMultiplier[stage]                  // 免費遊戲乘倍
	// 決定免費遊戲類型
	result.FreeGameType = slot.RandFreeGameType(rtp, buyType, debugCmdList)
	// 取得轉輪群組
	var reelGroup = FGReelGroup[rtp][result.FreeGameType]
	// 生成 WildTableId 列表
	var wildTableIdList = slot.CreateWildTableIdList(rtp, result.FreeGameType, scatterCount, stage)
	// 還有免費遊戲次數
	for spinCount < result.FGSpinCount {
		var spinWin uint64 = 0
		var spin = SpinResult{}
		var wildTableId = wildTableIdList[spinCount]
		spinCount++

		// 產生盤面
		var debugCmd *DebugCmd
		if len(debugCmdList) > spinCount {
			debugCmd = &debugCmdList[spinCount]
		}
		spin.TumbleSymbol, spin.ReelIndex, spin.H5ScoreArray = slot.RandFGSymbol(reelGroup, rtp, result.FreeGameType, lineBet, wildTableId, debugCmd)

		// 計算該盤面贏分
		spinWin, spin.LineSymbol, spin.LineCount, spin.LineWin = slot.CalculateWin(lineBet, spin.TumbleSymbol)

		// 計算 H5 獎圖總贏分
		var sumH5Score = 0
		for col := 0; col < len(spin.H5ScoreArray); col++ {
			sumH5Score += Utils.Sum(spin.H5ScoreArray[col])
		}
		var addWildCount = slot.GetSymbolCount(spin.TumbleSymbol, WW)
		cumWildCount += addWildCount
		spin.H5Win = uint64(sumH5Score * addWildCount * multiplier)
		spinWin += spin.H5Win

		// 檢查表演條件
		if slot.IsFGPerformance(spin.TumbleSymbol) {
			// 決定表演類型
			spin.PerformanceType = Utils.RandChoiceByWeight(FGPerformanceList, FGPerformanceWT[rtp])
			// 處理測試指令
			if debugCmd != nil && len(debugCmd.DebugData) > DEBUG_INDEX_PERFORMANCE_TYPE {
				// 取得表演類型，並檢查是否合法
				var debugPerformanceType = debugCmd.DebugData[DEBUG_INDEX_PERFORMANCE_TYPE]
				if 0 <= debugPerformanceType && debugPerformanceType <= REPLACE_OTHERS {
					spin.PerformanceType = debugPerformanceType
				}
			}
			// 取得表演盤面
			if spin.PerformanceType == REPLACE_WILD {
				spin.PerformanceSymbol = slot.ReplaceWild(spin.TumbleSymbol, result.FreeGameType)
			} else if spin.PerformanceType == REPLACE_H5 {
				spin.PerformanceSymbol = slot.ReplaceH5(spin.TumbleSymbol, result.FreeGameType)
			} else if spin.PerformanceType == REPLACE_OTHERS {
				spin.PerformanceSymbol = slot.ReplaceOthers(spin.TumbleSymbol, result.FreeGameType)
			}
		}

		// 達到最大贏分則結束
		spinWin, isMaxWin = Utils.CheckMaxWin(spinWin, maxFreeWin)
		if isMaxWin {
			// TODO: Log?
			result.FGSpinCount = spinCount // 結束免費遊戲
		}

		// 儲存結果
		spin.SpinWin = spinWin
		spin.CumWildCount = cumWildCount
		spin.Multiplier = multiplier
		spin.Stage = stage
		freeWin += spinWin
		maxFreeWin -= spinWin
		result.FGSpinList = append(result.FGSpinList, spin)

		// 檢查是否進階 更新階段及免費遊戲場次 (需未達最大贏分)
		if !isMaxWin && spinCount == result.FGSpinCount && cumWildCount >= FGStageWildCount[stage] {
			// 升階
			stage++
			// 更新乘倍
			multiplier = FGStageMultiplier[stage]
			// 增加免費遊戲場次
			result.FGSpinCount += FGAddSpinCount[stage]
			// 增加 WildTableId 列表
			var addWildTableIdList = slot.CreateWildTableIdList(rtp, result.FreeGameType, scatterCount, stage)
			wildTableIdList = append(wildTableIdList, addWildTableIdList...)
		}
	}

	result.FGCumWildCount = cumWildCount
	result.FGMultiplier = multiplier
	result.FGStage = stage
	result.FreeWin = freeWin
	return result.FreeWin
}

// RandMGSymbol 亂數產生主遊戲獎圖盤面
func (slot *SlotProb) RandMGSymbol(rtp int, buyType int, lineBet int, debugCmd *DebugCmd) (GameSymbol, int, []int, [][]int, int) {
	// 決定轉輪群組
	var roulette = slot.MGReelGroupRouletteMaps[rtp][buyType]
	var groupIdx, ok = roulette.Spin()
	if !ok {
		fmt.Printf("[ERROR] RandMGSymbol: RTP = %d, BuyType = %d, MGReelGroupRoulette spin failed.\n", rtp, buyType)
		return nil, -1, nil, nil, Common.ERROR_CODE_ROULETTE_SPIN_FAILED
	}
	// 處理測試指令
	var debugReelIndex []int
	var debugScoreIndices []int
	if debugCmd != nil {
		// 取得轉輪群組 index，並檢查是否合法
		if len(debugCmd.DebugData) > DEBUG_INDEX_GROUP_INDEX && 0 <= debugCmd.DebugData[DEBUG_INDEX_GROUP_INDEX] && debugCmd.DebugData[DEBUG_INDEX_GROUP_INDEX] < len(MGReelGroup[rtp]) {
			groupIdx = debugCmd.DebugData[DEBUG_INDEX_GROUP_INDEX]
		}
		// 取得停輪位置
		if len(debugCmd.DebugData) > DEBUG_INDEX_REEL_INDEX_05 {
			debugReelIndex = debugCmd.DebugData[DEBUG_INDEX_REEL_INDEX_01 : DEBUG_INDEX_REEL_INDEX_05+1]
		}
		// 取得 H5 獎圖分數編號(索引)列表
		if len(debugCmd.DebugData) > DEBUG_INDEX_H5_SCORE_15 {
			debugScoreIndices = make([]int, SLOT_COL*SLOT_ROW)
			// 後續處理過程會被異動，因此需要進行複製
			copy(debugScoreIndices, debugCmd.DebugData[DEBUG_INDEX_H5_SCORE_01:DEBUG_INDEX_H5_SCORE_15+1])
		}
	}
	var reelGroup = MGReelGroup[rtp][groupIdx]

	// 產出盤面
	var gameSymbol = make(GameSymbol, SLOT_COL)
	var reelIndex = make([]int, SLOT_COL)
	for col := 0; col < SLOT_COL; col++ {
		// 產出獎圖
		var reel = reelGroup[col] // 轉輪帶
		var reelLength = len(reel)
		var dice = rand.IntN(reelLength)
		// 處理測試指令中的停輪位置，並檢查是否合法
		if debugReelIndex != nil && 0 <= debugReelIndex[col] && debugReelIndex[col] < reelLength {
			dice = debugReelIndex[col]
		}
		var columnSymbol = make(ReelSymbol, SLOT_ROW)
		for row := 0; row < SLOT_ROW; row++ {
			var idx = dice + row
			if idx < reelLength {
				columnSymbol[row] = reel[idx]
			} else {
				columnSymbol[row] = reel[idx-reelLength]
			}
		}
		gameSymbol[col] = columnSymbol
		reelIndex[col] = dice
	}

	// 若 Scatter 獎圖數量為 0，則進入覆蓋 Scatter 獎圖流程
	if slot.GetSymbolCount(gameSymbol, SS) == 0 {
		// 決定覆蓋數量
		var scatterRoulette = slot.MGCoverScatterRouletteMaps[rtp][buyType]
		var count, ok = scatterRoulette.Spin()
		if !ok {
			fmt.Printf("[ERROR] RandMGSymbol: RTP = %d, BuyType = %d, MGCoverScatterRoulette spin failed.\n", rtp, buyType)
			return nil, -1, nil, nil, Common.ERROR_CODE_ROULETTE_SPIN_FAILED
		}
		// 處理測試指令
		if debugCmd != nil {
			// 取得覆蓋 Scatter 獎圖數量，並檢查是否合法
			if len(debugCmd.DebugData) > DEBUG_INDEX_COVER_SCATTER_COUNT && WIN_SCATTER_COUNT <= debugCmd.DebugData[DEBUG_INDEX_COVER_SCATTER_COUNT] && debugCmd.DebugData[DEBUG_INDEX_COVER_SCATTER_COUNT] <= MAX_WIN_SCATTER_COUNT {
				count = debugCmd.DebugData[DEBUG_INDEX_COVER_SCATTER_COUNT]
			}
		}
		// 覆蓋 Scatter 獎圖
		gameSymbol = slot.CoverScatter(gameSymbol, count)
	}

	// 亂數產生 H5 分數陣列 (主遊戲不計 H5 分數，僅供 Client 顯示，所以可以最後產出)
	var H5CountList = slot.GetH5CountList(gameSymbol)
	var scoreArray = slot.RandH5ScoreArray(rtp, -1, lineBet, H5CountList, debugScoreIndices)
	return gameSymbol, groupIdx, reelIndex, scoreArray, Common.ERROR_CODE_OK
}

// RandFreeGameType 亂數決定免費遊戲類型
func (slot *SlotProb) RandFreeGameType(rtp int, buyType int, debugCmdList []DebugCmd) int {
	// 處理測試指令
	if len(debugCmdList) > 0 {
		// 取得免費遊戲類型，並檢查是否合法
		if len(debugCmdList[0].DebugData) > DEBUG_INDEX_FREE_GAME_TYPE {
			var freeGameType = debugCmdList[0].DebugData[DEBUG_INDEX_FREE_GAME_TYPE]
			if (freeGameType == FREE_GAME_01 && buyType != Common.BUY_SUPER_FREE_SPINS) || freeGameType == FREE_GAME_02 {
				return freeGameType
			}
		}
	}
	// 決定轉輪群組 index (免費遊戲類型)
	var roulette = slot.FGReelGroupRouletteMaps[rtp][buyType]
	var groupIdx, ok = roulette.Spin()
	if !ok {
		fmt.Printf("[ERROR] RandFreeGameType: RTP = %d, BuyType = %d, FGReelGroupRouletteMaps spin failed.\n", rtp, buyType)
		return -1
	}
	return groupIdx
}

// CreateWildTableIdList 生成 WildTableId 列表
func (slot *SlotProb) CreateWildTableIdList(rtp int, freeGameType int, scatterCount int, stage int) []int {
	// 取得 Wild Table 個數列表
	var wildTableArray, ok = FGWildTableMap[rtp][freeGameType][scatterCount]
	if !ok {
		fmt.Printf("[ERROR] CreateWildTableIdList: Failed to get WildTableArray, RTP = %d, FreeGameType = %d, ScatterCount = %d.\n", rtp, freeGameType, scatterCount)
		return nil
	}
	if stage < 0 || stage > len(wildTableArray) {
		fmt.Printf("[ERROR] CreateWildTableIdList: Stage[%d] is invalid.\n", stage)
		return nil
	}
	var countList = wildTableArray[stage]

	// 根據個數列表產生 Id 列表
	var idList []int
	for i, count := range countList {
		for j := 0; j < count; j++ {
			idList = append(idList, WildTableList[i])
		}
	}
	// fmt.Printf("CreateWildTableIdList: idList = %v\n", idList)

	// 隨機排序 Id 列表
	rand.Shuffle(len(idList), func(i, j int) { idList[i], idList[j] = idList[j], idList[i] })
	return idList
}

// RandFGSymbol 亂數產生免費遊戲獎圖盤面
func (slot *SlotProb) RandFGSymbol(reelGroup [][]Symbol, rtp int, freeGameType int, lineBet int, wildTableId int, debugCmd *DebugCmd) (GameSymbol, []int, [][]int) {
	// 處理測試指令
	var debugReelIndex []int
	var debugScoreIndices []int
	if debugCmd != nil {
		// 取得停輪位置
		if len(debugCmd.DebugData) > DEBUG_INDEX_REEL_INDEX_05 {
			debugReelIndex = debugCmd.DebugData[DEBUG_INDEX_REEL_INDEX_01 : DEBUG_INDEX_REEL_INDEX_05+1]
		}
		// 取得 H5 獎圖分數編號(索引)列表
		if len(debugCmd.DebugData) > DEBUG_INDEX_H5_SCORE_15 {
			debugScoreIndices = make([]int, SLOT_COL*SLOT_ROW)
			// 後續處理過程會被異動，因此需要進行複製
			copy(debugScoreIndices, debugCmd.DebugData[DEBUG_INDEX_H5_SCORE_01:DEBUG_INDEX_H5_SCORE_15+1])
		}
	}
	// 產出盤面
	var gameSymbol = make(GameSymbol, SLOT_COL)
	var reelIndex = make([]int, SLOT_COL)
	for col := 0; col < SLOT_COL; col++ {
		// 產出獎圖
		var reel = reelGroup[col] // 轉輪帶
		var reelLength = len(reel)
		var dice = rand.IntN(reelLength)
		// 處理測試指令中的停輪位置，並檢查是否合法
		if debugReelIndex != nil && 0 <= debugReelIndex[col] && debugReelIndex[col] < reelLength {
			dice = debugReelIndex[col]
		}
		var columnSymbol = make(ReelSymbol, SLOT_ROW)
		for row := 0; row < SLOT_ROW; row++ {
			var idx = dice + row
			if idx < reelLength {
				columnSymbol[row] = reel[idx]
			} else {
				columnSymbol[row] = reel[idx-reelLength]
			}
		}
		gameSymbol[col] = columnSymbol
		reelIndex[col] = dice
	}

	// 亂數產生 H5 分數陣列
	var H5CountList = slot.GetH5CountList(gameSymbol)
	var scoreArray = slot.RandH5ScoreArray(rtp, freeGameType, lineBet, H5CountList, debugScoreIndices)

	// 覆蓋 Wild 獎圖流程，決定要覆蓋的 Wild 數量
	var wildCount = slot.RandWildCount(rtp, wildTableId, Utils.Sum(H5CountList))
	// 處理測試指令
	if debugCmd != nil {
		// 取得覆蓋 Wild 獎圖數量，並檢查是否合法
		if len(debugCmd.DebugData) > DEBUG_INDEX_COVER_WILD_COUNT && 0 <= debugCmd.DebugData[DEBUG_INDEX_COVER_WILD_COUNT] && debugCmd.DebugData[DEBUG_INDEX_COVER_WILD_COUNT] <= MAX_COVER_WILD_COUNT {
			wildCount = debugCmd.DebugData[DEBUG_INDEX_COVER_WILD_COUNT]
		}
	}
	if wildCount > 0 {
		// 覆蓋 Wild 獎圖
		gameSymbol, scoreArray = slot.CoverWild(gameSymbol, scoreArray, wildCount)
	}
	return gameSymbol, reelIndex, scoreArray
}

// RespinSymbol 重轉盤面
func (slot *SlotProb) RespinSymbol(lastSymbol GameSymbol, rtp int, groupIdx int, lineBet int, lastReelIndex []int, lastScoreArray [][]int) (GameSymbol, []int, [][]int) {
	// 取得含有 Scatter 的 cols
	var hasScatterCols = make([]bool, len(lastSymbol))
	var scatterCount = 0
	for col := 0; col < len(lastSymbol); col++ {
		for row := 0; row < len(lastSymbol[col]); row++ {
			if lastSymbol[col][row] == SS {
				hasScatterCols[col] = true
				// 順便檢查 Scatter 是否在最底部
				if row == len(lastSymbol[col])-1 {
					fmt.Printf("[WARN] RespinSymbol: Scatter is in the bottom of the column.\n")
					return lastSymbol, lastReelIndex, lastScoreArray
				}
				scatterCount++
				break
			}
		}
	}
	// 檢查 Scatter 數量
	if scatterCount != WIN_SCATTER_COUNT-1 {
		fmt.Printf("[WARN] RespinSymbol: Scatter count = %d, should be %d.\n", scatterCount, WIN_SCATTER_COUNT-1)
		return lastSymbol, lastReelIndex, lastScoreArray
	}

	// 取得轉輪群組
	var reelGroup = MGReelGroup[rtp][groupIdx]
	// 產出新盤面
	var gameSymbol = make(GameSymbol, SLOT_COL)
	var reelIndex = make([]int, SLOT_COL)
	var H5CountList = make([]int, SLOT_COL)
	for col := 0; col < SLOT_COL; col++ {
		// 產出獎圖
		var reel = reelGroup[col] // 轉輪帶
		var reelLength = len(reel)
		var dice = lastReelIndex[col] - 1
		if dice < 0 {
			dice += reelLength
		}
		// 不含 Scatter 的 cols 要重骰
		if !hasScatterCols[col] {
			dice = rand.IntN(reelLength)
		}
		var columnSymbol = make(ReelSymbol, SLOT_ROW)
		var H5Count = 0
		for row := 0; row < SLOT_ROW; row++ {
			var idx = dice + row
			if idx < reelLength {
				columnSymbol[row] = reel[idx]
			} else {
				columnSymbol[row] = reel[idx-reelLength]
			}
			if columnSymbol[row] == H5 {
				H5Count++
			}
		}
		gameSymbol[col] = columnSymbol
		reelIndex[col] = dice
		H5CountList[col] = H5Count
	}

	// 處理 H5 分數陣列
	var scoreArray = slot.RandH5ScoreArray(rtp, -1, lineBet, H5CountList, nil)
	for col := 0; col < SLOT_COL; col++ {
		// 重骰的 cols 不需處理
		if !hasScatterCols[col] {
			continue
		}
		// 根據 H5 獎圖數量，決定保留 H5 分數個數
		var lastScoreCount = len(lastScoreArray[col])
		if H5CountList[col] >= lastScoreCount {
			scoreArray[col] = append(lastScoreArray[col], scoreArray[col][lastScoreCount:]...)
		} else {
			scoreArray[col] = lastScoreArray[col][:H5CountList[col]]
		}
	}
	return gameSymbol, reelIndex, scoreArray
}

// GetSymbolCount 計算指定獎圖數量
func (slot *SlotProb) GetSymbolCount(gameSymbol GameSymbol, symbol Symbol) int {
	var count = 0
	for col := 0; col < len(gameSymbol); col++ {
		for row := 0; row < len(gameSymbol[col]); row++ {
			if gameSymbol[col][row] == symbol {
				count++
			}
		}
	}
	return count
}

// CoverScatter 覆蓋 Scatter 獎圖
func (slot *SlotProb) CoverScatter(gameSymbol GameSymbol, count int) GameSymbol {
	if count <= 0 {
		return gameSymbol
	} else if count > SLOT_COL {
		count = SLOT_COL
	}
	// 決定哪幾輪要覆蓋 Scatter 獎圖
	var allCols = [SLOT_COL]int{0, 1, 2, 3, 4}
	rand.Shuffle(len(allCols), func(i, j int) { allCols[i], allCols[j] = allCols[j], allCols[i] })
	// fmt.Printf("CoverScatter: shuffle allCols = %v\n", allCols)
	var colList = allCols[:count]
	// fmt.Printf("CoverScatter: colList = %v\n", colList)

	// 覆蓋獎圖
	for _, col := range colList {
		var row = rand.IntN(SLOT_ROW)
		// fmt.Printf("CoverScatter: [%d, %d] %v -> %v\n", col, row, gameSymbol[col][row], SS)
		gameSymbol[col][row] = SS
	}
	return gameSymbol
}

// GetH5CountList 計算各 reel(col) 中 H5 獎圖數量
func (slot *SlotProb) GetH5CountList(gameSymbol GameSymbol) []int {
	var countList = make([]int, len(gameSymbol))
	for col := 0; col < len(gameSymbol); col++ {
		var count = 0
		for row := 0; row < len(gameSymbol[col]); row++ {
			if gameSymbol[col][row] == H5 {
				count++
			}
		}
		countList[col] = count
	}
	return countList
}

// RandH5ScoreArray 依 H5 獎圖數量陣列，呼叫 RandH5ScoreList 產生 H5 獎圖分數二維陣列
func (slot *SlotProb) RandH5ScoreArray(rtp int, freeGameType int, lineBet int, countList []int, debugScoreIndices []int) [][]int {
	var scoreArray = make([][]int, len(countList))
	for col := 0; col < len(countList); col++ {
		scoreArray[col], debugScoreIndices = slot.RandH5ScoreList(rtp, freeGameType, lineBet, countList[col], debugScoreIndices)
	}
	return scoreArray
}

// RandH5ScoreList 亂數產生 H5 獎圖分數列表
func (slot *SlotProb) RandH5ScoreList(rtp int, freeGameType int, lineBet int, count int, debugScoreIndices []int) ([]int, []int) {
	// 取得 Roulette
	var roulette = slot.MGH5ScoreRoulette[rtp]
	if freeGameType >= FREE_GAME_01 && freeGameType <= FREE_GAME_02 {
		roulette = slot.FGH5ScoreRouletteList[rtp][freeGameType]
	}

	// 產生 H5 獎圖分數列表
	var totalBet = lineBet * PAYLINE_TOTAL
	var scoreList = make([]int, count)
	for i := 0; i < count; i++ {
		var score, ok = roulette.Spin()
		if !ok {
			fmt.Printf("[ERROR] RandH5ScoreList: RTP = %d, FreeGameType = %d, H5ScoreRoulette spin failed.\n", rtp, freeGameType)
			return scoreList, debugScoreIndices
		}
		// 處理測試指令：取得 H5 獎圖分數編號(索引)，並檢查是否合法
		if i < len(debugScoreIndices) && 0 <= debugScoreIndices[i] && debugScoreIndices[i] < len(H5ScoreList) {
			score = H5ScoreList[debugScoreIndices[i]]
		}
		// 實際分數為 score 乘 TotalBet
		scoreList[i] = score * totalBet
	}
	// 移除使用過的 H5 獎圖分數編號(索引)
	if len(debugScoreIndices) > count {
		debugScoreIndices = debugScoreIndices[count:]
	}
	return scoreList, debugScoreIndices
}

// RandWildCount 亂數決定覆蓋 Wild 數量
func (slot *SlotProb) RandWildCount(rtp int, wildTableId int, H5Count int) int {
	if H5Count <= 0 {
		return 0
	}
	// 取得 Wild 個數權重列表
	var wildCountWT = FGWildCountWT[rtp][wildTableId]
	if H5Count >= len(wildCountWT) {
		H5Count = len(wildCountWT) - 1
	}
	var weightList = wildCountWT[H5Count]
	// 亂數決定覆蓋 Wild 數量
	var count = Utils.RandChoiceByWeight(WildCountList, weightList)
	// fmt.Printf("RandWildCount: Count = %d\n", count)
	return count
}

// CoverWild 覆蓋 Wild 獎圖
func (slot *SlotProb) CoverWild(gameSymbol GameSymbol, scoreArray [][]int, count int) (GameSymbol, [][]int) {
	if count <= 0 {
		return gameSymbol, scoreArray
	}

	// 將最高分的 H5 覆蓋為 Wild 獎圖
	for i := 0; i < count; i++ {
		// 移除最高的 H5 分數
		// fmt.Printf("Before H5ScoreArray: %v\n", scoreArray)
		var targetCol = -1
		scoreArray, targetCol = slot.RemoveMaxH5Score(scoreArray)
		if targetCol < 0 {
			fmt.Printf("[ERROR] CoverWild: RemoveMaxH5Score failed.\n")
			return gameSymbol, scoreArray
		}
		// fmt.Printf("After H5ScoreArray: %v\n", scoreArray)

		// 取得 targetCol 中的 H5 位置列表
		var H5Rows []int
		for row := 0; row < len(gameSymbol[targetCol]); row++ {
			if gameSymbol[targetCol][row] == H5 {
				H5Rows = append(H5Rows, row)
			}
		}

		// 覆蓋獎圖
		var targetRow = H5Rows[rand.IntN(len(H5Rows))]
		// fmt.Printf("CoverWild: [%d, %d] %v -> %v\n", targetCol, targetRow, gameSymbol[targetCol][targetRow], WW)
		gameSymbol[targetCol][targetRow] = WW
	}
	return gameSymbol, scoreArray
}

// RemoveMaxH5Score 移除 H5 最高分
func (slot *SlotProb) RemoveMaxH5Score(scoreArray [][]int) ([][]int, int) {
	// 找出最高分所在的 col
	var maxScore int
	var targetCol = -1
	for col := 0; col < len(scoreArray); col++ {
		for row := 0; row < len(scoreArray[col]); row++ {
			if scoreArray[col][row] > maxScore {
				maxScore = scoreArray[col][row]
				targetCol = col
			}
		}
	}
	if targetCol < 0 {
		return scoreArray, targetCol
	}
	scoreArray[targetCol] = Utils.RemoveMaximum(scoreArray[targetCol])
	return scoreArray, targetCol
}

// CalculateWin 計算連線獎金
func (slot *SlotProb) CalculateWin(lineBet int, gameSymbol GameSymbol) (totalWin uint64, lineSymbolList []Symbol, lineCountList []int, lineWinList []uint64) {
	totalWin = 0
	lineSymbolList = make([]Symbol, len(LineIndexArray))
	lineCountList = make([]int, len(LineIndexArray))
	lineWinList = make([]uint64, len(LineIndexArray))

	for i, lineIndex := range LineIndexArray {
		var lineSymbol = make([]Symbol, len(lineIndex))
		for col, row := range lineIndex {
			lineSymbol[col] = gameSymbol[col][row]
		}
		var symbol, count = slot.CheckLine(lineSymbol)
		if symbol >= WW {
			var lineWin = uint64(lineBet * SymbolOdds[int(symbol)][count])
			totalWin += lineWin
			lineSymbolList[i] = symbol
			lineCountList[i] = count
			lineWinList[i] = lineWin
		} else {
			lineSymbolList[i] = NN
		}
		// fmt.Printf("CalculateWin: %d) %v Win: %d\n", i, lineSymbol, lineWin)
	}
	return
}

// CheckLine 檢查連線及獎圖
func (slot *SlotProb) CheckLine(lineSymbol []Symbol) (symbol Symbol, count int) {
	if len(lineSymbol) < SLOT_COL {
		fmt.Printf("[ERROR] CheckLine: LineSymbol length is invalid.\n")
		return NN, 0
	}

	symbol, count = NN, 0
	var findSymbol = lineSymbol[0]
	var linkCount = 0
	// 比對獎圖
	for i, s := range lineSymbol {
		if s == SS {
			break
		}
		if !(s == findSymbol || s == WW || findSymbol == WW) {
			break
		}
		if s != WW {
			findSymbol = s
		}
		linkCount = i + 1
	}
	// 檢查獎圖連線個數是否中獎
	if findSymbol != NN && linkCount >= WinSymbolCount[findSymbol] {
		symbol = findSymbol
		count = linkCount
		// 若第一個獎圖為 Wild，則檢查連續 Wild 是否有更高賠率
		if lineSymbol[0] == WW {
			// 計算連續 Wild 數
			var wildCount = 1
			for i := 1; i < len(lineSymbol); i++ {
				if lineSymbol[i] == WW && lineSymbol[i-1] == WW {
					wildCount++
				} else {
					// 不連續則中斷計算
					break
				}
			}
			// 選擇賠率較高者
			if wildCount >= WinSymbolCount[WW] && wildCount < SLOT_COL {
				if SymbolOdds[WW][wildCount] >= SymbolOdds[symbol][count] {
					symbol = WW
					count = wildCount
				}
			} else if wildCount == SLOT_COL {
				// 全 Wild
				symbol = WW
				count = SLOT_COL
			}
		}
	}
	return symbol, count
}

// IsRespin 檢查是否符合重轉條件
func (slot *SlotProb) IsRespin(gameSymbol GameSymbol) bool {
	// 檢查 Scatter 數量
	if slot.GetSymbolCount(gameSymbol, SS) != WIN_SCATTER_COUNT-1 {
		return false
	}

	// 檢查最底部是否含 Scatter
	for col := 0; col < len(gameSymbol); col++ {
		var row = len(gameSymbol[col]) - 1
		if gameSymbol[col][row] == SS {
			return false
		}
	}
	return true
}

// IsFGPerformance 檢查免費遊戲表演條件
func (slot *SlotProb) IsFGPerformance(gameSymbol GameSymbol) bool {
	// 檢查 Wild 及 H5 數量
	var wildCount, H5Count = 0, 0
	for col := 0; col < len(gameSymbol); col++ {
		for row := 0; row < len(gameSymbol[col]); row++ {
			if gameSymbol[col][row] == WW {
				wildCount++
			} else if gameSymbol[col][row] == H5 {
				H5Count++
			}
		}
		// 前三輪 Wild 數量超過則不表演
		if col <= 2 && wildCount > FREE_GAME_CONDITION_WILD_COUNT {
			return false
		}
	}
	// H5 數量不足則不表演
	if H5Count < FREE_GAME_CONDITION_H5_COUNT {
		return false
	}
	// 沒有 H5 得分則不表演
	if wildCount == 0 || H5Count == 0 {
		return false
	}
	return true
}

// ReplaceScatter 替換 Scatter 表演 (隨機替換一個 Scatter)
func (slot *SlotProb) ReplaceScatter(lastSymbol GameSymbol) GameSymbol {
	// 複製盤面，並取得 Scatter 位置列表
	var gameSymbol = make(GameSymbol, len(lastSymbol))
	var indices [][2]int
	for col := 0; col < len(lastSymbol); col++ {
		gameSymbol[col] = make(ReelSymbol, len(lastSymbol[col]))
		copy(gameSymbol[col], lastSymbol[col])
		for row := 0; row < len(lastSymbol[col]); row++ {
			if lastSymbol[col][row] == SS {
				indices = append(indices, [2]int{col, row})
			}
		}
	}

	// 檢查 Scatter 數量
	if len(indices) < WIN_SCATTER_COUNT {
		fmt.Printf("[ERROR] ReplaceScatter: Scatter count is invalid.\n")
		return nil
	}

	// 隨機決定取代位置
	var pos = rand.IntN(len(indices))
	var col = indices[pos][0]
	var row = indices[pos][1]
	// 排除前一輪出現的獎圖
	var excludeCol = col - 1
	if excludeCol < 0 {
		excludeCol = col + 1
	}
	// 亂數產生用來覆蓋的獎圖
	var symbol = slot.RandCoverSymbol(gameSymbol[excludeCol], -1)
	// fmt.Printf("ReplaceScatter: [%d, %d] %v -> %v\n", col, row, gameSymbol[col][row], symbol)
	gameSymbol[col][row] = symbol
	return gameSymbol
}

// ReplaceWild 替換 Wild 表演 (出現 Wild 的整輪替換)
func (slot *SlotProb) ReplaceWild(lastSymbol GameSymbol, freeGameType int) GameSymbol {
	// 複製盤面，並取得 Wild 的 col 列表
	var gameSymbol = make(GameSymbol, len(lastSymbol))
	var colList []int
	for col := 0; col < len(lastSymbol); col++ {
		gameSymbol[col] = make(ReelSymbol, len(lastSymbol[col]))
		copy(gameSymbol[col], lastSymbol[col])
		for row := 0; row < len(lastSymbol[col]); row++ {
			if lastSymbol[col][row] == WW {
				colList = append(colList, col)
				break
			}
		}
	}

	// 檢查 Wild 數量
	if len(colList) <= 0 {
		fmt.Printf("[ERROR] ReplaceWild: No Wild.\n")
		return nil
	}

	// 出現 Wild 的整輪替換
	for _, col := range colList {
		// 排除前一輪出現的獎圖
		var excludeCol = col - 1
		if excludeCol < 0 {
			excludeCol = col + 1
		}
		// 整輪替換
		for row := 0; row < len(gameSymbol[col]); row++ {
			gameSymbol[col][row] = slot.RandCoverSymbol(gameSymbol[excludeCol], freeGameType)
			// fmt.Printf("ReplaceWild: [%d, %d] %v -> %v\n", col, row, lastSymbol[col][row], gameSymbol[col][row])
		}
	}
	return gameSymbol
}

// ReplaceH5 替換 H5 表演 (替換全部的 H5 獎圖)
func (slot *SlotProb) ReplaceH5(lastSymbol GameSymbol, freeGameType int) GameSymbol {
	// 複製盤面，並取得 H5 位置列表
	var gameSymbol = make(GameSymbol, len(lastSymbol))
	var indices [][2]int
	for col := 0; col < len(lastSymbol); col++ {
		gameSymbol[col] = make(ReelSymbol, len(lastSymbol[col]))
		copy(gameSymbol[col], lastSymbol[col])
		for row := 0; row < len(lastSymbol[col]); row++ {
			if lastSymbol[col][row] == H5 {
				indices = append(indices, [2]int{col, row})
			}
		}
	}

	// 檢查 H5 數量
	if len(indices) <= 0 {
		fmt.Printf("[ERROR] ReplaceH5: No H5.\n")
		return nil
	}

	// 替換全部的 H5 獎圖
	for _, index := range indices {
		var col = index[0]
		var row = index[1]
		// 排除前一輪出現的獎圖
		var excludeCol = col - 1
		if excludeCol < 0 {
			excludeCol = col + 1
		}
		// 替換獎圖
		gameSymbol[col][row] = slot.RandCoverSymbol(gameSymbol[excludeCol], freeGameType)
		// fmt.Printf("ReplaceH5: [%d, %d] %v -> %v\n", col, row, lastSymbol[col][row], gameSymbol[col][row])
	}
	return gameSymbol
}

// ReplaceOthers 替換 Wild 以外的獎圖表演
func (slot *SlotProb) ReplaceOthers(lastSymbol GameSymbol, freeGameType int) GameSymbol {
	// 保留 Wild，其他獎圖全部替換
	var gameSymbol = make(GameSymbol, len(lastSymbol))
	// 先處理第一輪，排除第二輪出現的獎圖
	gameSymbol[0] = make(ReelSymbol, len(lastSymbol[0]))
	for row := 0; row < len(lastSymbol[0]); row++ {
		if lastSymbol[0][row] == WW {
			gameSymbol[0][row] = WW
		} else {
			gameSymbol[0][row] = slot.RandCoverSymbol(lastSymbol[1], freeGameType)
			// fmt.Printf("ReplaceOthers: [0, %d] %v -> %v\n", row, lastSymbol[0][row], gameSymbol[0][row])
		}
	}
	// 再處理剩餘的轉輪
	for col := 1; col < len(lastSymbol); col++ {
		gameSymbol[col] = make(ReelSymbol, len(lastSymbol[col]))
		// 排除前一輪出現的獎圖
		var excludeCol = col - 1
		var excludeSymbol = make(ReelSymbol, len(gameSymbol[excludeCol]))
		copy(excludeSymbol, gameSymbol[excludeCol])
		// 第三輪以後，排除前兩輪出現的獎圖 (以防前兩輪的 Wild 形成連線)
		if excludeCol > 0 {
			excludeSymbol = append(excludeSymbol, gameSymbol[excludeCol-1]...)
		}
		for row := 0; row < len(lastSymbol[col]); row++ {
			if lastSymbol[col][row] == WW {
				gameSymbol[col][row] = WW
			} else {
				gameSymbol[col][row] = slot.RandCoverSymbol(excludeSymbol, freeGameType)
				// fmt.Printf("ReplaceOthers: [%d, %d] %v -> %v\n", col, row, lastSymbol[col][row], gameSymbol[col][row])
			}
		}
	}
	return gameSymbol
}

// RandCoverSymbol 亂數產生用來覆蓋的獎圖
func (slot *SlotProb) RandCoverSymbol(excludeSymbol ReelSymbol, freeGameType int) Symbol {
	// FREE_GAME_02 直接回傳空獎圖
	if freeGameType == FREE_GAME_02 {
		return NN
	}

	// 取得可用的獎圖
	var symbolList []Symbol
	for _, symbol := range CoverSymbolList {
		// 排除獎圖
		var isExclude = false
		for _, exclude := range excludeSymbol {
			if symbol == exclude {
				isExclude = true
				break
			}
		}
		if !isExclude {
			symbolList = append(symbolList, symbol)
		}
	}
	if len(symbolList) <= 0 {
		fmt.Printf("[ERROR] RandCoverSymbol: No available symbol.\n")
		return NN
	}
	// fmt.Printf("RandCoverSymbol: SymbolList = %v\n", symbolList)

	// 隨機決定獎圖
	var dice = rand.IntN(len(symbolList))
	return symbolList[dice]
}

// SetRTPRoulette 建立 Roulette 資料
func (slot *SlotProb) SetRTPRoulette() {
	for rtp := 0; rtp < RTP_TOTAL; rtp++ {
		// 主遊戲轉輪群組 Roulette Map
		var rouletteMap = make(map[int]Utils.Roulette[int])
		for buyType, groupWT := range MGReelGroupWT[rtp] {
			var roulette = Utils.NewRouletteFromList(groupWT)
			// roulette.Dump()
			rouletteMap[buyType] = *roulette
		}
		slot.MGReelGroupRouletteMaps[rtp] = rouletteMap

		// 免費遊戲轉輪群組 Roulette Map
		rouletteMap = make(map[int]Utils.Roulette[int])
		for buyType, groupWT := range FGReelGroupWT[rtp] {
			var roulette = Utils.NewRouletteFromList(groupWT)
			// roulette.Dump()
			rouletteMap[buyType] = *roulette
		}
		slot.FGReelGroupRouletteMaps[rtp] = rouletteMap

		// 主遊戲覆蓋 Scatter 數量 Roulette Map
		rouletteMap = make(map[int]Utils.Roulette[int])
		for buyType, coverScatterCountWT := range MGCoverScatterCountWT[rtp] {
			var roulette = Utils.NewRouletteFromMap(coverScatterCountWT)
			// roulette.Dump()
			rouletteMap[buyType] = *roulette
		}
		slot.MGCoverScatterRouletteMaps[rtp] = rouletteMap

		// 主遊戲 H5 獎圖分數 Roulette
		var h5ScoreRoulette = Utils.NewRouletteFromMap(MGH5ScoreWT[rtp])
		// h5ScoreRoulette.Dump()
		slot.MGH5ScoreRoulette[rtp] = *h5ScoreRoulette

		// 免費遊戲 H5 獎圖分數 Roulette List
		var h5ScoreRouletteList []Utils.Roulette[int]
		for _, scoreWT := range FGH5ScoreWT[rtp] {
			var roulette = Utils.NewRouletteFromMap(scoreWT)
			// roulette.Dump()
			h5ScoreRouletteList = append(h5ScoreRouletteList, *roulette)
		}
		slot.FGH5ScoreRouletteList[rtp] = h5ScoreRouletteList
	}
}

// ShowGameSymbol 顯示獎圖盤面
func ShowGameSymbol(gameSymbol GameSymbol) {
	if gameSymbol == nil {
		fmt.Println("ShowGameSymbol: GameSymbol is nil.")
		return
	}
	for row := 0; row < SLOT_ROW; row++ {
		var lineSymbol = make([]Symbol, SLOT_COL)
		for col := 0; col < SLOT_COL; col++ {
			lineSymbol[col] = gameSymbol[col][row]
		}
		fmt.Println(lineSymbol)
	}
}
