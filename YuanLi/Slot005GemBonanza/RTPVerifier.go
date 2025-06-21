package Slot005GemBonanza

import (
	"Force/GameServer/Common"
	"Force/GameServer/Utils"
	"fmt"
)

const MAX_RESPIN_COUNT = 2 // 重轉次數統計到 2+
const MAX_WILD_COUNT = 12  // 免費遊戲 CumWildCount 統計到 12+
const MAX_STAGE = 3        // 免費遊戲最大階段

var FreeGameName = []string{"FREE_GAME_01", "FREE_GAME_02"}

type RTPVerifier struct {
	TotalCount int // 總投注次數
	LineBet    int // 單次投注額
	slot       SlotProb

	TotalBet     float64 // 總投注額
	TotalMainWin float64 // 主遊戲總贏分
	TotalFreeWin float64 // 免費遊戲總贏分
	MaxMainWin   uint64  // 主遊戲最大贏分
	MaxFreeWin   uint64  // 免費遊戲最大贏分

	MGSymbolWin [MAX_SYMBOL][SLOT_COL + 1]float64 // 主遊戲獎圖贏分列表
	FGSymbolWin [MAX_SYMBOL][SLOT_COL + 1]float64 // 免費遊戲獎圖贏分列表

	MGRespinWin [MAX_RESPIN_COUNT + 1]float64 // 主遊戲重轉次數贏分列表

	MGHitCount       int                   // 主遊戲有贏分的次數
	FGHitCount       int                   // 免費遊戲有贏分的次數
	FGCount          [FREE_GAME_02 + 1]int // 進入免費遊戲次數
	FGTotalSpinCount int                   // 免費遊戲總場次

	MGTotalRespinCount int // 主遊戲重轉次數總和

	FGH5Win      [FREE_GAME_02 + 1][MAX_STAGE + 1]float64                             // 免費遊戲 H5 獎圖贏分 (依 FreeGameType、Stage 分別統計)
	FGWildCount  [FREE_GAME_02 + 1][MAX_WIN_SCATTER_COUNT + 1][MAX_WILD_COUNT + 1]int // 免費遊戲 CumWildCount 統計 (依 FreeGameType、ScatterCount 分別統計)
	FGStageCount [FREE_GAME_02 + 1][MAX_WIN_SCATTER_COUNT + 1][MAX_STAGE + 1]int      // 免費遊戲 Stage 統計 (依 FreeGameType、ScatterCount 分別統計)
}

type RTPResult struct {
	SlotResult
	TotalBet float64
}

// Init 初始化
func (rv *RTPVerifier) Init() {
	rv.LineBet = 1
	rv.slot.Init()
}

// Clear 清除資料
func (rv *RTPVerifier) Clear() {
	rv.TotalBet = 0
	rv.TotalMainWin = 0
	rv.TotalFreeWin = 0
	rv.MaxMainWin = 0
	rv.MaxFreeWin = 0
	rv.MGHitCount = 0
	rv.FGHitCount = 0
	rv.FGTotalSpinCount = 0
	rv.MGTotalRespinCount = 0
}

// RunAll 驗證所有 RTP
func (rv *RTPVerifier) RunAll(buyType int) {
	for i := 0; i < RTP_TOTAL; i++ {
		rv.Run(i, buyType)
	}
}

// Run 驗證 RTP
func (rv *RTPVerifier) Run(rtp int, buyType int) {
	// 檢查參數 RTP 是否正確
	if rtp < 0 || rtp > RTP_TOTAL {
		fmt.Printf("Run: rtp out of range: %d\n", rtp)
		return
	}

	rv.Clear()

	// 投注額
	var totalBet = float64(rv.LineBet * PAYLINE_TOTAL)
	// 根據購買類型增加投注額
	totalBet *= BetRatio[buyType]

	for i := 0; i < rv.TotalCount; i++ {
		var result = RTPResult{SlotResult: SlotResult{BuyType: buyType}, TotalBet: totalBet}
		rv.Bet(rtp, rv.LineBet, &result)
	}
}

// Bet 投注一次
func (rv *RTPVerifier) Bet(rtp int, lineBet int, result *RTPResult) {
	// 累計總投注額
	rv.TotalBet += result.TotalBet

	// 進行遊戲
	var gameMode = rv.slot.Run(rtp, lineBet, &result.SlotResult, nil)

	// 處理結果
	rv.ProcessResult(gameMode, result)
}

// ProcessResult 處理投注結果
func (rv *RTPVerifier) ProcessResult(gameMode int, result *RTPResult) {
	// 主遊戲
	if result.MainWin > 0 {
		rv.MGHitCount++
		rv.TotalMainWin += float64(result.MainWin)
		if rv.MaxMainWin < result.MainWin {
			rv.MaxMainWin = result.MainWin
		}
	}
	for _, tumble := range result.MGTumbleList {
		for i := 0; i < len(LineIndexArray); i++ {
			var symbol = tumble.LineSymbol[i]
			var count = tumble.LineCount[i]
			if symbol > NN && count > 0 {
				rv.MGSymbolWin[symbol][count] += float64(tumble.LineWin[i])
			}
		}
	}
	var respinCount = len(result.MGTumbleList) - 1
	rv.MGTotalRespinCount += respinCount
	if respinCount > MAX_RESPIN_COUNT {
		respinCount = MAX_RESPIN_COUNT
	}
	rv.MGRespinWin[respinCount] += float64(result.MainWin)

	// 免費遊戲
	if gameMode&Common.GAME_MODE_FREE == Common.GAME_MODE_FREE {
		rv.FGCount[result.FreeGameType]++
		rv.FGTotalSpinCount += result.FGSpinCount
		rv.TotalFreeWin += float64(result.FreeWin)
		if rv.MaxFreeWin < result.FreeWin {
			rv.MaxFreeWin = result.FreeWin
		}
		for _, spin := range result.FGSpinList {
			if spin.SpinWin > 0 {
				rv.FGHitCount++
			}
			for i := 0; i < len(LineIndexArray); i++ {
				var symbol = spin.LineSymbol[i]
				var count = spin.LineCount[i]
				if symbol > NN && count > 0 {
					rv.FGSymbolWin[symbol][count] += float64(spin.LineWin[i])
				}
			}
			rv.FGH5Win[result.FreeGameType][spin.Stage] += float64(spin.H5Win)
		}
		var lastSymbol = result.MGTumbleList[respinCount].TumbleSymbol
		var scatterCount = rv.slot.GetSymbolCount(lastSymbol, SS)
		var cumWildCount = result.FGCumWildCount
		if cumWildCount > MAX_WILD_COUNT {
			cumWildCount = MAX_WILD_COUNT
		}
		rv.FGWildCount[result.FreeGameType][scatterCount][cumWildCount]++
		rv.FGStageCount[result.FreeGameType][scatterCount][result.FGStage]++
	}
}

// Dump 傾印統計結果
func (rv *RTPVerifier) Dump(detail bool) {
	fmt.Printf("Total Count: %d\n", rv.TotalCount)

	var totalWin = rv.TotalMainWin + rv.TotalFreeWin
	var FGTotalCount = Utils.Sum(rv.FGCount[:])
	fmt.Printf("TotalBet: %.1f, TotalWin: %.1f, Total RTP: %s\n", rv.TotalBet, totalWin, Utils.GetPercentage(totalWin, rv.TotalBet))
	fmt.Printf("  %16s  %10s  %10s  %10s  %10s  %10s\n", "RTP %", "Hit %", "Trigger %", "MaxOdds", "Avg.Round", "Avg.Respin")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Main: %12s  %10s", Utils.GetPercentage(rv.TotalMainWin, rv.TotalBet), Utils.GetPercentage(float64(rv.MGHitCount), float64(rv.TotalCount)))
	fmt.Printf("  %10s  %10.2f  %10s  %10.2f\n", "", float64(rv.MaxMainWin)/float64(rv.LineBet*PAYLINE_TOTAL), "", float64(rv.MGTotalRespinCount)/float64(rv.MGHitCount))
	fmt.Printf("Free: %12s  %10s", Utils.GetPercentage(rv.TotalFreeWin, rv.TotalBet), Utils.GetPercentage(float64(rv.FGHitCount), float64(rv.FGTotalSpinCount)))
	fmt.Printf("  %10s  %10.2f  %10.2f  %10s\n", Utils.GetPercentage(float64(FGTotalCount), float64(rv.TotalCount)), float64(rv.MaxFreeWin)/float64(rv.LineBet*PAYLINE_TOTAL), float64(rv.FGTotalSpinCount)/float64(FGTotalCount), "")

	if !detail {
		return
	}
	fmt.Printf("=================================  %10s  =================================\n", "Main  Game")
	fmt.Printf("  %22s  %16s  %16s  %16s\n", "x2", "x3", "x4", "x5")
	for symbol := WW; symbol < SS; symbol++ {
		fmt.Printf("  %3s)", symbol)
		for count := 2; count <= SLOT_COL; count++ {
			fmt.Printf("  %16s", Utils.GetPercentage(rv.MGSymbolWin[symbol][count], rv.TotalBet))
		}
		fmt.Println("")
	}
	fmt.Printf("--------------------------------  %12s  --------------------------------\n", "Respin Count")
	for rc := 0; rc < len(rv.MGRespinWin); rc++ {
		fmt.Printf("  %3d)  %16s\n", rc, Utils.GetPercentage(rv.MGRespinWin[rc], rv.TotalBet))
	}
	fmt.Printf("=================================  %10s  =================================\n", "Free  Game")
	fmt.Printf("  %22s  %16s  %16s  %16s\n", "x2", "x3", "x4", "x5")
	for symbol := WW; symbol < SS; symbol++ {
		fmt.Printf("  %3s)", symbol)
		for count := 2; count <= SLOT_COL; count++ {
			fmt.Printf("  %16s", Utils.GetPercentage(rv.FGSymbolWin[symbol][count], rv.TotalBet))
		}
		fmt.Println("")
	}
	fmt.Printf("-------------------------------- %14s --------------------------------\n", "H5 Score RTP %")
	for i, name := range FreeGameName {
		fmt.Printf("  %s\n", name)
		fmt.Printf("  %22s  %16s  %16s  %16s\n", "x1", "x2", "x3", "x10")
		fmt.Printf("  %4s", "")
		for stage := 0; stage <= MAX_STAGE; stage++ {
			fmt.Printf("  %16s", Utils.GetPercentage(rv.FGH5Win[i][stage], rv.TotalBet))
		}
		fmt.Println("")
	}
	fmt.Printf("--------------------------------  %12s  --------------------------------\n", "CumWildCount")
	for i, name := range FreeGameName {
		fmt.Printf("  %s\n", name)
		fmt.Printf("  %13s  %7s  %7s  %7s  %7s  %7s  %7s  %7s  %7s  %7s  %7s  %7s  %7s\n", "0)", "1)", "2)", "3)", "4)", "5)", "6)", "7)", "8)", "9)", "10)", "11)", "12)")
		for sc := 3; sc <= MAX_WIN_SCATTER_COUNT; sc++ {
			fmt.Printf("  %2dSS", sc)
			var sumWildCount = Utils.Sum(rv.FGWildCount[i][sc][:])
			for wc := 0; wc <= MAX_WILD_COUNT; wc++ {
				fmt.Printf(" %8s", Utils.GetPercentage(float64(rv.FGWildCount[i][sc][wc]), float64(sumWildCount)))
			}
			fmt.Println("")
		}
	}
	fmt.Printf("-----------------------------------  %5s  ------------------------------------\n", "Stage")
	for i, name := range FreeGameName {
		fmt.Printf("  %s  (%d/%d)\n", name, rv.FGCount[i], FGTotalCount)
		fmt.Printf("  %22s  %16s  %16s  %16s\n", "1)", "2)", "3)", "4)")
		for sc := 3; sc <= MAX_WIN_SCATTER_COUNT; sc++ {
			fmt.Printf("  %2dSS", sc)
			var sumStageCount = Utils.Sum(rv.FGStageCount[i][sc][:])
			for stage := 0; stage <= MAX_STAGE; stage++ {
				fmt.Printf("  %16s", Utils.GetPercentage(float64(rv.FGStageCount[i][sc][stage]), float64(sumStageCount)))
			}
			fmt.Println("")
		}
	}
}
