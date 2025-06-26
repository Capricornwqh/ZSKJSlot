package Slot013FatToad

import (
	"Force/GameServer/Common"
	"Force/GameServer/Utils"
	"fmt"
	"sync"
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

	MGTotalFeatureWin   float64 // 主遊戲天隆横财總和
	MGTotalFeatureCount int     // 主遊戲天隆横财總和
	FGTotalFeatureWin   float64 // 免費遊戲天隆横财總和
	FGTotalFeatureCount int     // 免費遊戲天隆横财總和

	FGWWLevel      [7]int // 免費遊戲金蟾等級
	FGWWMultiplier []int  // 免費遊戲金蟾乘倍

	FGIndexCount [FREE_INDEX_2 + 1]int // 免費遊戲索引次數 (依 FreeGameType 統計)

	FGH5Win      [FREE_GAME_02 + 1][MAX_STAGE + 1]float64                             // 免費遊戲 H5 獎圖贏分 (依 FreeGameType、Stage 分別統計)
	FGWildCount  [FREE_GAME_02 + 1][MAX_WIN_SCATTER_COUNT + 1][MAX_WILD_COUNT + 1]int // 免費遊戲 CumWildCount 統計 (依 FreeGameType、ScatterCount 分別統計)
	FGStageCount [FREE_GAME_02 + 1][MAX_WIN_SCATTER_COUNT + 1][MAX_STAGE + 1]int      // 免費遊戲 Stage 統計 (依 FreeGameType、ScatterCount 分別統計)
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
	rv.MGTotalFeatureWin = 0
	rv.MGTotalFeatureCount = 0
	rv.FGTotalFeatureWin = 0
	rv.FGTotalFeatureCount = 0
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

	// 确定并发数量，根据CPU核心数量或设定合理的值
	numWorkers := 10
	if numWorkers > rv.TotalCount {
		numWorkers = rv.TotalCount
	}

	// 每个协程处理的批次大小
	batchSize := rv.TotalCount / numWorkers
	if rv.TotalCount%numWorkers != 0 {
		batchSize++
	}

	var wg sync.WaitGroup
	resultChan := make(chan *SlotResult, 1000)

	// 启动工作协程
	for i := range numWorkers {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// 计算该协程负责的范围
			start := workerID * batchSize
			end := start + batchSize
			if end > rv.TotalCount {
				end = rv.TotalCount
			}

			for j := start; j < end; j++ {
				result := &SlotResult{BuyType: buyType}

				rv.slot.Run(rtp, rv.LineBet, result, nil)
				resultChan <- result
			}
		}(i)
	}

	// 等待所有投注协程结束，关闭结果通道
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 统一处理结果
	for result := range resultChan {
		rv.TotalBet += float64(result.TotalBet)
		gameMode := Common.GAME_MODE_NORMAL
		if result.FGSpinCount > 0 {
			gameMode = Common.GAME_MODE_FREE
		}
		rv.ProcessResult(gameMode, result)
	}
}

// ProcessResult 處理投注結果
func (rv *RTPVerifier) ProcessResult(gameMode int, result *SlotResult) {
	// 主遊戲
	if result.MainWin > 0 {
		rv.MGHitCount++
		rv.TotalMainWin += float64(result.MainWin)
		if rv.MaxMainWin < result.MainWin {
			rv.MaxMainWin = result.MainWin
		}
	}
	for _, tumble := range result.MGTumbleList {
		if tumble == nil || tumble.LineSymbol == nil || tumble.LineCount == nil || tumble.LineWin == nil {
			continue
		}
		for i := range LineIndexArray {
			var symbol = tumble.LineSymbol[i]
			var count = tumble.LineCount[i]
			if symbol > NN && count > 0 {
				rv.MGSymbolWin[symbol][count] += float64(tumble.LineWin[i])
			}
		}
	}

	rv.MGTotalFeatureWin += float64(result.MGFeatureWin)
	rv.MGTotalFeatureCount += result.MGFeatureCount

	// 免費遊戲
	if gameMode&Common.GAME_MODE_FREE == Common.GAME_MODE_FREE {
		rv.FGCount[result.FreeGameType]++
		rv.FGIndexCount[result.FGIndex]++
		rv.FGTotalSpinCount += result.FGSpinCount
		rv.TotalFreeWin += float64(result.FreeWin)
		if rv.MaxFreeWin < result.FreeWin {
			rv.MaxFreeWin = result.FreeWin
		}
		for _, spin := range result.FGSpinList {
			if spin == nil || spin.LineSymbol == nil || spin.LineCount == nil || spin.LineWin == nil {
				continue
			}
			if spin.SpinWin > 0 {
				rv.FGHitCount++
			}
			for i := range LineIndexArray {
				var symbol = spin.LineSymbol[i]
				var count = spin.LineCount[i]
				if symbol > NN && count > 0 {
					rv.FGSymbolWin[symbol][count] += float64(spin.LineWin[i])
				}
			}
			// rv.FGH5Win[result.FreeGameType][spin.Stage] += float64(spin.H5Win)
		}
		rv.FGWWLevel[result.WWLevel]++
		rv.FGWWMultiplier = append(rv.FGWWMultiplier, result.WWMultiplier)
		rv.FGTotalFeatureCount += result.FGFeatureCount
		// var lastSymbol = result.MGTumbleList[respinCount].TumbleSymbol
		// var scatterCount = rv.slot.GetSymbolCount(lastSymbol, SS)
		// var cumWildCount = result.FGCumWildCount
		// if cumWildCount > MAX_WILD_COUNT {
		// 	cumWildCount = MAX_WILD_COUNT
		// }
		// rv.FGWildCount[result.FreeGameType][scatterCount][cumWildCount]++
		// rv.FGStageCount[result.FreeGameType][scatterCount][result.FGStage]++
	}
}

// Dump 傾印統計結果
func (rv *RTPVerifier) Dump(detail bool) {
	fmt.Printf("Total Count: %d\n", rv.TotalCount)

	var totalWin = rv.TotalMainWin + rv.TotalFreeWin
	var FGTotalCount = Utils.Sum(rv.FGCount[:])
	fmt.Printf("TotalBet: %12.1f, TotalWin: %12.1f, Total RTP: %s\n", rv.TotalBet, totalWin, Utils.GetPercentage(totalWin, rv.TotalBet))
	fmt.Printf("  %-8s  %16s  %10s  %10s  %10s  %10s  %10s  %10s\n", "", "Win", "RTP %", "Hit %", "Trigger %", "MaxOdds", "Avg.Round", "Avg.Respin")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("  %-8s  %16.1f  %10s  %10s", "Main:", rv.TotalMainWin, Utils.GetPercentage(rv.TotalMainWin, rv.TotalBet), Utils.GetPercentage(float64(rv.MGHitCount), float64(rv.TotalCount)))
	fmt.Printf("  %10s  %10.2f  %10s  %10.2f\n", "", float64(rv.MaxMainWin)/float64(rv.LineBet*PAYLINE_TOTAL), "", float64(rv.MGTotalFeatureCount)/float64(rv.MGHitCount))
	// fmt.Printf("  %-8s  %16.1f  %10s  %10s\n", "Feature:", rv.MGTotalFeatureWin, Utils.GetPercentage(rv.MGTotalFeatureWin, rv.TotalBet), Utils.GetPercentage(float64(rv.MGTotalFeatureCount), float64(rv.TotalCount)))
	// fmt.Printf("  %10s  %10.2f  %10s  %10.2f\n", "", "", "")
	fmt.Printf("  %-8s  %16.1f  %10s  %10s", "Free:", rv.TotalFreeWin, Utils.GetPercentage(rv.TotalFreeWin, rv.TotalBet), Utils.GetPercentage(float64(rv.FGHitCount), float64(rv.FGTotalSpinCount)))
	fmt.Printf("  %10s  %10.2f  %10.2f  %10s\n", Utils.GetPercentage(float64(FGTotalCount), float64(rv.TotalCount)), float64(rv.MaxFreeWin)/float64(rv.LineBet*PAYLINE_TOTAL), float64(rv.FGTotalSpinCount)/float64(FGTotalCount), "")
	// fmt.Printf("  %-8s  %16.1f  %10s  %10s\n", "Feature:", float64(0), "", Utils.GetPercentage(float64(rv.FGTotalFeatureCount), float64(rv.FGTotalSpinCount)))
	// fmt.Printf("  %10s  %10.2f  %10.2f  %10s\n", Utils.GetPercentage(float64(FGTotalCount), float64(rv.TotalCount)), float64(rv.MaxFreeWin)/float64(rv.LineBet*PAYLINE_TOTAL), float64(rv.FGTotalSpinCount)/float64(FGTotalCount), "")
	fmt.Printf("------------------------  %20s  ------------------------\n", "FGIndex Count")
	sumFGIndex := Utils.Sum(rv.FGIndexCount[:])
	for j := range len(rv.FGIndexCount) {
		fmt.Printf("%d:%10s    ", j+1, Utils.GetPercentage(float64(rv.FGIndexCount[j]), float64(sumFGIndex)))
	}
	fmt.Println()
	// fmt.Printf("--------------------------------------------------------------------------------\n")
	// fmt.Printf("Free Count %10.2f\n", float64(rv.TotalCount)/float64(rv.FGTotalSpinCount))
	// if !detail {
	// 	return
	// }
	fmt.Printf("--------------------------------  %12s  --------------------------------\n", "WW Wild")
	sumWWLevel := Utils.Sum(rv.FGWWLevel[:])
	for j := 1; j <= 6; j++ {
		fmt.Printf("%d:%10s    ", j, Utils.GetPercentage(float64(rv.FGWWLevel[j]), float64(sumWWLevel)))
	}
	fmt.Println()
	// fmt.Printf("=================================  %10s  =================================\n", "Main  Game")
	// fmt.Printf("  %22s  %16s  %16s  %16s  %16s\n", "x2", "x3", "x4", "x5", "x6")
	// for symbol := WW; symbol <= SS; symbol++ {
	// 	fmt.Printf("  %3s)", symbol)
	// 	for count := 2; count <= SLOT_COL; count++ {
	// 		fmt.Printf("  %16s", Utils.GetPercentage(rv.MGSymbolWin[symbol][count], rv.TotalBet))
	// 	}
	// 	fmt.Println("")
	// }
	// fmt.Printf("=================================  %10s  =================================\n", "Free  Game")
	// fmt.Printf("  %22s  %16s  %16s  %16s  %16s\n", "x2", "x3", "x4", "x5", "x6")
	// for symbol := WW; symbol <= SS; symbol++ {
	// 	fmt.Printf("  %3s)", symbol)
	// 	for count := 2; count <= SLOT_COL; count++ {
	// 		fmt.Printf("  %16s", Utils.GetPercentage(rv.FGSymbolWin[symbol][count], rv.TotalBet))
	// 	}
	// 	fmt.Println("")
	// }
	// fmt.Printf("-------------------------------- %14s --------------------------------\n", "H5 Score RTP %")
	// for i, name := range FreeGameName {
	// 	fmt.Printf("  %s\n", name)
	// 	fmt.Printf("  %22s  %16s  %16s  %16s\n", "x1", "x2", "x3", "x10")
	// 	fmt.Printf("  %4s", "")
	// 	for stage := 0; stage <= MAX_STAGE; stage++ {
	// 		fmt.Printf("  %16s", Utils.GetPercentage(rv.FGH5Win[i][stage], rv.TotalBet))
	// 	}
	// 	fmt.Println("")
	// }
	// fmt.Printf("--------------------------------  %12s  --------------------------------\n", "CumWildCount")
	// for i, name := range FreeGameName {
	// 	fmt.Printf("  %s\n", name)
	// 	fmt.Printf("  %13s  %7s  %7s  %7s  %7s  %7s  %7s  %7s  %7s  %7s  %7s  %7s  %7s\n", "0)", "1)", "2)", "3)", "4)", "5)", "6)", "7)", "8)", "9)", "10)", "11)", "12)")
	// 	for sc := 3; sc <= MAX_WIN_SCATTER_COUNT; sc++ {
	// 		fmt.Printf("  %2dSS", sc)
	// 		var sumWildCount = Utils.Sum(rv.FGWildCount[i][sc][:])
	// 		for wc := 0; wc <= MAX_WILD_COUNT; wc++ {
	// 			fmt.Printf(" %8s", Utils.GetPercentage(float64(rv.FGWildCount[i][sc][wc]), float64(sumWildCount)))
	// 		}
	// 		fmt.Println("")
	// 	}
	// }
	// fmt.Printf("-----------------------------------  %5s  ------------------------------------\n", "Stage")
	// for i, name := range FreeGameName {
	// 	fmt.Printf("  %s  (%d/%d)\n", name, rv.FGCount[i], FGTotalCount)
	// 	fmt.Printf("  %22s  %16s  %16s  %16s\n", "1)", "2)", "3)", "4)")
	// 	for sc := 3; sc <= MAX_WIN_SCATTER_COUNT; sc++ {
	// 		fmt.Printf("  %2dSS", sc)
	// 		var sumStageCount = Utils.Sum(rv.FGStageCount[i][sc][:])
	// 		for stage := 0; stage <= MAX_STAGE; stage++ {
	// 			fmt.Printf("  %16s", Utils.GetPercentage(float64(rv.FGStageCount[i][sc][stage]), float64(sumStageCount)))
	// 		}
	// 		fmt.Println("")
	// 	}
	// }
}
