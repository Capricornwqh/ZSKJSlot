package Slot005GemBonanza

import (
	"Force/GameServer/Common"
	"Force/GameServer/Utils"
	"fmt"
	"reflect"
	"testing"
)

func TestSlotProbRun(t *testing.T) {
	type args struct {
		rtp          int
		lineBet      int
		result       *SlotResult
		debugCmdList []DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"進行遊戲: 一般投注",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_NONE}},
		},
		{
			"進行遊戲: 額外投注",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_EXTRA_BET}},
		},
		{
			"進行遊戲: 購買免費",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_FREE_SPINS}},
		},
		{
			"進行遊戲: 購買超級",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_SUPER_FREE_SPINS}},
		},
		{
			"進行遊戲 測試指令: 指定主遊戲盤面 Respin 不表演",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_NONE},
				debugCmdList: []DebugCmd{
					{DebugData: []int{2, 12, 2, 30, 10, 5, -1, 0, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 0, 1}},
				}},
		},
		{
			"進行遊戲 測試指令: 覆蓋3個Scatter 表演 REPLACE_SCATTER",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_EXTRA_BET},
				debugCmdList: []DebugCmd{
					{DebugData: []int{-1, -1, -1, -1, -1, -1, 3, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 0, 1}},
				}},
		},
		{
			"進行遊戲 測試指令: 指定主遊戲盤面 指定免費遊戲盤面1面 FREE_GAME_01",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_FREE_SPINS},
				debugCmdList: []DebugCmd{
					{DebugData: []int{-1, 5, 5, 5, 5, 5, -1, FREE_GAME_01, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 0, 0}},
					{DebugData: []int{-1, 5, 5, 5, 5, 5, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 0, -1}},
				}},
		},
		{
			"進行遊戲 測試指令: 指定主遊戲盤面 指定免費遊戲盤面1面 FREE_GAME_02",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_FREE_SPINS},
				debugCmdList: []DebugCmd{
					{DebugData: []int{-1, 5, 5, 5, 5, 5, -1, FREE_GAME_02, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 0, 0}},
					{DebugData: []int{-1, 5, 5, 5, 5, 5, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 1, -1}},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var gameMode = tt.slot.Run(tt.args.rtp, tt.args.lineBet, tt.args.result, tt.args.debugCmdList)
			fmt.Printf("Game Mode: %d\n", gameMode)
			for i, tumble := range tt.args.result.MGTumbleList {
				ShowGameSymbol(tumble.TumbleSymbol)
				fmt.Printf("%d) GroupIdx: %d, ReelIndex: %v\n", i, tt.args.result.MGGroupIndex, tumble.ReelIndex)
				fmt.Printf("%d) H5CountList: %v, H5ScoreArray: %v\n", i, tt.slot.GetH5CountList(tumble.TumbleSymbol), tumble.H5ScoreArray)
				fmt.Printf("%d) LineSymbol: %v\n", i, tumble.LineSymbol)
				fmt.Printf("%d) LineCount: %v\n", i, tumble.LineCount)
				fmt.Printf("%d) LineWin: %v\n", i, tumble.LineWin)
				fmt.Printf("%d) Win: %v\n", i, tumble.Win)
			}
			if tt.args.result.MGPerformanceType > NONE {
				fmt.Printf("PerformanceType: %v\n", tt.args.result.MGPerformanceType)
				ShowGameSymbol(tt.args.result.MGPerformanceSymbol)
			}
			fmt.Printf("MainWin: %d\n", tt.args.result.MainWin)

			fmt.Printf("FGSpinCount: %v\n", tt.args.result.FGSpinCount)
			for i, spin := range tt.args.result.FGSpinList {
				ShowGameSymbol(spin.TumbleSymbol)
				fmt.Printf("%d] FreeGameType: %d, ReelIndex: %v\n", i, tt.args.result.FreeGameType, spin.ReelIndex)
				fmt.Printf("%d] H5CountList: %v, H5ScoreArray: %v\n", i, tt.slot.GetH5CountList(spin.TumbleSymbol), spin.H5ScoreArray)
				fmt.Printf("%d] LineSymbol: %v\n", i, spin.LineSymbol)
				fmt.Printf("%d] LineCount: %v\n", i, spin.LineCount)
				fmt.Printf("%d] LineWin: %v\n", i, spin.LineWin)
				fmt.Printf("%d] H5Win: %d, SpinWin: %d\n", i, spin.H5Win, spin.SpinWin)
				if spin.PerformanceType > NONE {
					fmt.Printf("%d] PerformanceType: %v\n", i, spin.PerformanceType)
					ShowGameSymbol(spin.PerformanceSymbol)
				}
				fmt.Printf("%d] CumWildCount: %d\n", i, spin.CumWildCount)
				fmt.Printf("%d] Multiplier: %d\n", i, spin.Multiplier)
				fmt.Printf("%d] Stage: %d\n", i, spin.Stage)
			}
			fmt.Printf("FGCumWildCount: %d\n", tt.args.result.FGCumWildCount)
			fmt.Printf("FGMultiplier: %d\n", tt.args.result.FGMultiplier)
			fmt.Printf("FGStage: %d\n", tt.args.result.FGStage)
			fmt.Printf("FreeWin: %d\n", tt.args.result.FreeWin)
			fmt.Printf("TotalWin: %d\n", tt.args.result.TotalWin)
		})
	}
}

func TestSlotProbRunFreeGame(t *testing.T) {
	type args struct {
		rtp          int
		lineBet      int
		scatterCount int
		result       *SlotResult
		debugCmdList []DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"進行免費遊戲: 一般投注 3SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 3, result: &SlotResult{BuyType: Common.BUY_NONE}},
		},
		{
			"進行免費遊戲: 一般投注 4SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 4, result: &SlotResult{BuyType: Common.BUY_NONE}},
		},
		{
			"進行免費遊戲: 一般投注 5SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 5, result: &SlotResult{BuyType: Common.BUY_NONE}},
		},
		{
			"進行免費遊戲: 額外投注 3SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 3, result: &SlotResult{BuyType: Common.BUY_EXTRA_BET}},
		},
		{
			"進行免費遊戲: 額外投注 4SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 4, result: &SlotResult{BuyType: Common.BUY_EXTRA_BET}},
		},
		{
			"進行免費遊戲: 額外投注 5SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 5, result: &SlotResult{BuyType: Common.BUY_EXTRA_BET}},
		},
		{
			"進行免費遊戲: 購買免費 3SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 3, result: &SlotResult{BuyType: Common.BUY_FREE_SPINS}},
		},
		{
			"進行免費遊戲: 購買免費 4SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 4, result: &SlotResult{BuyType: Common.BUY_FREE_SPINS}},
		},
		{
			"進行免費遊戲: 購買免費 5SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 5, result: &SlotResult{BuyType: Common.BUY_FREE_SPINS}},
		},
		{
			"進行免費遊戲: 購買超級 3SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 3, result: &SlotResult{BuyType: Common.BUY_SUPER_FREE_SPINS}},
		},
		{
			"進行免費遊戲: 購買超級 4SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 4, result: &SlotResult{BuyType: Common.BUY_SUPER_FREE_SPINS}},
		},
		{
			"進行免費遊戲: 購買超級 5SS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 5, result: &SlotResult{BuyType: Common.BUY_SUPER_FREE_SPINS}},
		},
		{
			"進行免費遊戲 測試指令: 指定免費遊戲盤面2面 FREE_GAME_01 覆蓋1個Wild 指定H5獎圖分數",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 3, result: &SlotResult{BuyType: Common.BUY_NONE},
				debugCmdList: []DebugCmd{
					{DebugData: []int{DEBUG_INDEX_FREE_GAME_TYPE: FREE_GAME_01}},
					{DebugData: []int{-1, 0, 1, 1, 0, 0, -1, 0, -1, 3, 3, 3, 3, -1, 4, 4, 4, 4, -1, 12, 12, 12, 12, 1, -1}},
					{DebugData: []int{-1, 34, 33, 27, 18, 21, -1, 0, -1, 3, 3, 4, 4, -1, 5, 5, 6, 6, -1, 12, 12, 12, 12, 1, -1}},
				}},
		},
		{
			"進行免費遊戲 測試指令: 指定免費遊戲盤面2面 FREE_GAME_02 覆蓋2個Wild 指定H5獎圖分數",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 3, result: &SlotResult{BuyType: Common.BUY_NONE},
				debugCmdList: []DebugCmd{
					{DebugData: []int{DEBUG_INDEX_FREE_GAME_TYPE: FREE_GAME_02}},
					{DebugData: []int{-1, 0, 1, 1, 0, 0, -1, 0, -1, 3, 3, 3, 3, -1, 4, 4, 4, 4, -1, 12, 12, 12, 12, 2, -1}},
					{DebugData: []int{-1, 35, 36, 30, 22, 24, -1, 0, -1, 3, 3, 4, 4, -1, 5, 5, 6, 6, -1, 12, 12, 12, 12, 2, -1}},
				}},
		},
		{
			"進行免費遊戲 測試指令: 指定免費遊戲盤面3面 FREE_GAME_01 表演 REPLACE_WILD、REPLACE_H5、REPLACE_OTHERS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 3, result: &SlotResult{BuyType: Common.BUY_NONE},
				debugCmdList: []DebugCmd{
					{DebugData: []int{DEBUG_INDEX_FREE_GAME_TYPE: FREE_GAME_01}},
					{DebugData: []int{-1, 8, 12, 10, 5, 12, -1, 0, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 1, 1}},
					{DebugData: []int{-1, 8, 12, 10, 5, 12, -1, 0, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 1, 2}},
					{DebugData: []int{-1, 8, 12, 10, 5, 12, -1, 0, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 1, 3}},
				}},
		},
		{
			"進行免費遊戲 測試指令: 指定免費遊戲盤面3面 FREE_GAME_02 表演 REPLACE_WILD、REPLACE_H5、REPLACE_OTHERS",
			&SlotProb{},
			args{rtp: 0, lineBet: 1, scatterCount: 3, result: &SlotResult{BuyType: Common.BUY_NONE},
				debugCmdList: []DebugCmd{
					{DebugData: []int{DEBUG_INDEX_FREE_GAME_TYPE: FREE_GAME_02}},
					{DebugData: []int{-1, 1, 4, 5, 2, 3, -1, 0, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 1, 1}},
					{DebugData: []int{-1, 1, 4, 5, 2, 3, -1, 0, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 1, 2}},
					{DebugData: []int{-1, 1, 4, 5, 2, 3, -1, 0, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, -2, -1, -1, -1, -1, 1, 3}},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			tt.args.result.FGSpinCount = FGInitSpinCount[tt.args.scatterCount]
			_ = tt.slot.RunFreeGame(tt.args.rtp, tt.args.lineBet, tt.args.scatterCount, tt.args.result, tt.args.debugCmdList)
			fmt.Printf("FGSpinCount: %v\n", tt.args.result.FGSpinCount)
			for i, spin := range tt.args.result.FGSpinList {
				ShowGameSymbol(spin.TumbleSymbol)
				fmt.Printf("%d] FreeGameType: %d, ReelIndex: %v\n", i, tt.args.result.FreeGameType, spin.ReelIndex)
				fmt.Printf("%d] H5CountList: %v, H5ScoreArray: %v\n", i, tt.slot.GetH5CountList(spin.TumbleSymbol), spin.H5ScoreArray)
				fmt.Printf("%d] LineSymbol: %v\n", i, spin.LineSymbol)
				fmt.Printf("%d] LineCount: %v\n", i, spin.LineCount)
				fmt.Printf("%d] LineWin: %v\n", i, spin.LineWin)
				fmt.Printf("%d] H5Win: %d, SpinWin: %d\n", i, spin.H5Win, spin.SpinWin)
				if spin.PerformanceType > NONE {
					fmt.Printf("%d] PerformanceType: %v\n", i, spin.PerformanceType)
					ShowGameSymbol(spin.PerformanceSymbol)
				}
				fmt.Printf("%d] CumWildCount: %d\n", i, spin.CumWildCount)
				fmt.Printf("%d] Multiplier: %d\n", i, spin.Multiplier)
				fmt.Printf("%d] Stage: %d\n", i, spin.Stage)
			}
			fmt.Printf("FGCumWildCount: %d\n", tt.args.result.FGCumWildCount)
			fmt.Printf("FGMultiplier: %d\n", tt.args.result.FGMultiplier)
			fmt.Printf("FGStage: %d\n", tt.args.result.FGStage)
			fmt.Printf("FreeWin: %d\n", tt.args.result.FreeWin)
		})
	}
}

func TestSlotProbRandMGSymbol(t *testing.T) {
	type args struct {
		rtp      int
		buyType  int
		debugCmd *DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"亂數產生主遊戲獎圖: 一般投注",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_NONE},
		},
		{
			"亂數產生主遊戲獎圖: 額外投注",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_EXTRA_BET},
		},
		{
			"亂數產生主遊戲獎圖: 購買免費",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_FREE_SPINS},
		},
		{
			"亂數產生主遊戲獎圖: 購買超級",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_SUPER_FREE_SPINS},
		},
		{
			"測試指令: 指定主遊戲盤面 覆蓋5個Scatter 指定H5獎圖分數",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_NONE,
				debugCmd: &DebugCmd{DebugData: []int{1, 3, 3, 3, 3, 3, 5, FREE_GAME_02, -1, 2, 2, 2, 2, -1, 2, 2, 2, 2, -1, 2, 2, 2, 2, 0, 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var gameSymbol, groupIdx, reelIndex, scoreArray, _ = tt.slot.RandMGSymbol(tt.args.rtp, tt.args.buyType, 1, tt.args.debugCmd)
			ShowGameSymbol(gameSymbol)
			fmt.Printf("GroupIdx: %d, ReelIndex: %v\n", groupIdx, reelIndex)
			fmt.Printf("H5CountList: %v, H5ScoreArray: %v\n", tt.slot.GetH5CountList(gameSymbol), scoreArray)
		})
	}
}

func TestSlotProbRandMGSymbolWT(t *testing.T) {
	type args struct {
		totalCount int
		rtp        int
		buyType    int
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"主遊戲轉輪群組權重: 一般投注",
			&SlotProb{},
			args{totalCount: 15220000, rtp: 0, buyType: Common.BUY_NONE},
		},
		{
			"主遊戲轉輪群組權重: 額外投注",
			&SlotProb{},
			args{totalCount: 15220000, rtp: 0, buyType: Common.BUY_EXTRA_BET},
		},
		{
			"主遊戲轉輪群組權重: 購買免費",
			&SlotProb{},
			args{totalCount: 1000, rtp: 0, buyType: Common.BUY_FREE_SPINS},
		},
		{
			"主遊戲轉輪群組權重: 購買超級",
			&SlotProb{},
			args{totalCount: 1000, rtp: 0, buyType: Common.BUY_SUPER_FREE_SPINS},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var countMap = make(map[int]int)
			for i := 0; i < tt.args.totalCount; i++ {
				var _, groupIdx, _, _, _ = tt.slot.RandMGSymbol(tt.args.rtp, tt.args.buyType, 1, nil)
				countMap[groupIdx]++
			}
			var keys = Utils.SortedMapKeys(countMap)
			for _, key := range keys {
				fmt.Printf(" GroupIdx: %d, Count: %d (%s)\n", key, countMap[key], Utils.GetPercentage(float64(countMap[key]), float64(tt.args.totalCount)))
			}
		})
	}
}

func TestSlotProbRandFreeGameType(t *testing.T) {
	type args struct {
		totalCount   int
		rtp          int
		buyType      int
		debugCmdList []DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"亂數決定免費遊戲類型: 一般投注",
			&SlotProb{},
			args{totalCount: 120000, rtp: 0, buyType: Common.BUY_NONE},
		},
		{
			"亂數決定免費遊戲類型: 額外投注",
			&SlotProb{},
			args{totalCount: 120000, rtp: 0, buyType: Common.BUY_EXTRA_BET},
		},
		{
			"亂數決定免費遊戲類型: 購買免費",
			&SlotProb{},
			args{totalCount: 100000, rtp: 0, buyType: Common.BUY_FREE_SPINS},
		},
		{
			"亂數決定免費遊戲類型: 購買超級",
			&SlotProb{},
			args{totalCount: 1000, rtp: 0, buyType: Common.BUY_SUPER_FREE_SPINS},
		},
		{
			"測試指令: 指定免費遊戲類型 購買免費 FREE_GAME_01",
			&SlotProb{},
			args{totalCount: 100, rtp: 0, buyType: Common.BUY_FREE_SPINS,
				debugCmdList: []DebugCmd{{DebugData: []int{DEBUG_INDEX_FREE_GAME_TYPE: FREE_GAME_01}}}},
		},
		{
			"測試指令: 指定免費遊戲類型 購買免費 FREE_GAME_02",
			&SlotProb{},
			args{totalCount: 100, rtp: 0, buyType: Common.BUY_FREE_SPINS,
				debugCmdList: []DebugCmd{{DebugData: []int{DEBUG_INDEX_FREE_GAME_TYPE: FREE_GAME_02}}}},
		},
		{
			"測試指令: 指定免費遊戲類型 購買超級 FREE_GAME_01 結果: 不生效",
			&SlotProb{},
			args{totalCount: 100, rtp: 0, buyType: Common.BUY_SUPER_FREE_SPINS,
				debugCmdList: []DebugCmd{{DebugData: []int{DEBUG_INDEX_FREE_GAME_TYPE: FREE_GAME_01}}}},
		},
		{
			"測試指令: 指定免費遊戲類型 購買超級 FREE_GAME_02",
			&SlotProb{},
			args{totalCount: 100, rtp: 0, buyType: Common.BUY_SUPER_FREE_SPINS,
				debugCmdList: []DebugCmd{{DebugData: []int{DEBUG_INDEX_FREE_GAME_TYPE: FREE_GAME_02}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var countMap = make(map[int]int)
			for i := 0; i < tt.args.totalCount; i++ {
				var freeGameType = tt.slot.RandFreeGameType(tt.args.rtp, tt.args.buyType, tt.args.debugCmdList)
				countMap[freeGameType]++
			}
			var keys = Utils.SortedMapKeys(countMap)
			for _, key := range keys {
				fmt.Printf(" FreeGameType: %d, Count: %d (%s)\n", key, countMap[key], Utils.GetPercentage(float64(countMap[key]), float64(tt.args.totalCount)))
			}
		})
	}
}

func TestSlotProbCreateWildTableIdList(t *testing.T) {
	type args struct {
		rtp          int
		freeGameType int
		scatterCount int
		stage        int
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"生成 WildTableId 列表: FREE_GAME_01 3SS Stage1",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 3, stage: 0},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 3SS Stage2",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 3, stage: 1},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 3SS Stage3",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 3, stage: 2},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 3SS Stage4",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 3, stage: 3},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 4SS Stage1",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 4, stage: 0},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 4SS Stage2",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 4, stage: 1},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 4SS Stage3",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 4, stage: 2},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 4SS Stage4",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 4, stage: 3},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 5SS Stage1",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 5, stage: 0},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 5SS Stage2",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 5, stage: 1},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 5SS Stage3",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 5, stage: 2},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_01 5SS Stage4",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, scatterCount: 5, stage: 3},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 3SS Stage1",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 3, stage: 0},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 3SS Stage2",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 3, stage: 1},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 3SS Stage3",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 3, stage: 2},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 3SS Stage4",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 3, stage: 3},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 4SS Stage1",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 4, stage: 0},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 4SS Stage2",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 4, stage: 1},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 4SS Stage3",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 4, stage: 2},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 4SS Stage4",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 4, stage: 3},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 5SS Stage1",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 5, stage: 0},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 5SS Stage2",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 5, stage: 1},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 5SS Stage3",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 5, stage: 2},
		},
		{
			"生成 WildTableId 列表: FREE_GAME_02 5SS Stage4",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, scatterCount: 5, stage: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var idList = tt.slot.CreateWildTableIdList(tt.args.rtp, tt.args.freeGameType, tt.args.scatterCount, tt.args.stage)
			fmt.Printf("CreateWildTableIdList: shuffle idList = %v\n", idList)
		})
	}
}

func TestSlotProbRandFGSymbol(t *testing.T) {
	type args struct {
		rtp          int
		freeGameType int
		wildTableId  int
		debugCmd     *DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"亂數產生免費遊戲獎圖: FREE_GAME_01 WILD_TABLE_01",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_01},
		},
		{
			"亂數產生免費遊戲獎圖: FREE_GAME_01 WILD_TABLE_02",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_02},
		},
		{
			"亂數產生免費遊戲獎圖: FREE_GAME_01 WILD_TABLE_03",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_03},
		},
		{
			"亂數產生免費遊戲獎圖: FREE_GAME_02 WILD_TABLE_01",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_01},
		},
		{
			"亂數產生免費遊戲獎圖: FREE_GAME_02 WILD_TABLE_02",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_02},
		},
		{
			"亂數產生免費遊戲獎圖: FREE_GAME_02 WILD_TABLE_03",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_03},
		},
		{
			"測試指令: 指定免費遊戲盤面 FREE_GAME_01 覆蓋1個Wild 指定H5獎圖分數",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_01,
				debugCmd: &DebugCmd{DebugData: []int{-1, 1, 1, 1, 1, 1, -1, 0, -1, 3, 3, 3, 3, -1, 4, 4, 4, 4, -1, 12, 12, 12, 12, 1, -1}}},
		},
		{
			"測試指令: 指定免費遊戲盤面 FREE_GAME_02 覆蓋2個Wild 指定H5獎圖分數",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_01,
				debugCmd: &DebugCmd{DebugData: []int{-1, 1, 1, 1, 1, 1, -1, 0, -1, 3, 3, 3, 3, -1, 4, 4, 4, 4, -1, 12, 12, 12, 12, 2, -1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var reelGroup = FGReelGroup[tt.args.rtp][tt.args.freeGameType]
			var gameSymbol, reelIndex, scoreArray = tt.slot.RandFGSymbol(reelGroup, tt.args.rtp, tt.args.freeGameType, 1, tt.args.wildTableId, tt.args.debugCmd)
			ShowGameSymbol(gameSymbol)
			fmt.Printf("FreeGameType: %d, ReelIndex: %v\n", tt.args.freeGameType, reelIndex)
			fmt.Printf("H5CountList: %v, H5ScoreArray: %v\n", tt.slot.GetH5CountList(gameSymbol), scoreArray)
		})
	}
}

func TestSlotProbRespinSymbol(t *testing.T) {
	type args struct {
		rtp      int
		buyType  int
		debugCmd *DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"重轉盤面: 一般投注",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_NONE},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var gameSymbol, groupIdx, reelIndex, scoreArray, _ = tt.slot.RandMGSymbol(tt.args.rtp, tt.args.buyType, 1, tt.args.debugCmd)
			ShowGameSymbol(gameSymbol)
			fmt.Printf("GroupIdx: %d, ReelIndex: %v\n", groupIdx, reelIndex)
			fmt.Printf("H5CountList: %v, H5ScoreArray: %v\n", tt.slot.GetH5CountList(gameSymbol), scoreArray)
			fmt.Printf("-------------------------------  %14s  -------------------------------\n", "Re-spin Symbol")
			var newGameSymbol, newReelIndex, newScoreArray = tt.slot.RespinSymbol(gameSymbol, tt.args.rtp, groupIdx, 1, reelIndex, scoreArray)
			ShowGameSymbol(newGameSymbol)
			fmt.Printf("GroupIdx: %d, NewReelIndex: %v\n", groupIdx, newReelIndex)
			fmt.Printf("NewH5CountList: %v, NewH5ScoreArray: %v\n", tt.slot.GetH5CountList(newGameSymbol), newScoreArray)
		})
	}
}

func TestSlotProbRandH5ScoreList(t *testing.T) {
	type args struct {
		rtp          int
		freeGameType int
		count        int
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"主遊戲 H5 獎圖分數",
			&SlotProb{},
			args{rtp: 0, freeGameType: -1, count: 122350000},
		},
		{
			"免費遊戲 H5 獎圖分數 FREE_GAME_01",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, count: 122350000},
		},
		{
			"免費遊戲 H5 獎圖分數 FREE_GAME_02",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, count: 62070000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var countMap = make(map[int]int)
			var scoreList, _ = tt.slot.RandH5ScoreList(tt.args.rtp, tt.args.freeGameType, 1, tt.args.count, nil)
			for _, score := range scoreList {
				countMap[score]++
			}
			var keys = Utils.SortedMapKeys(countMap)
			for _, key := range keys {
				fmt.Printf("[%d] Count: %d (%s)\n", key, countMap[key], Utils.GetPercentage(float64(countMap[key]), float64(tt.args.count)))
			}
		})
	}
}

func TestSlotProbRandWildCount(t *testing.T) {
	type args struct {
		totalCount  int
		rtp         int
		wildTableId int
		H5Count     int
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"亂數決定覆蓋 Wild 數量: WILD_TABLE_01 H5Count 1",
			&SlotProb{},
			args{totalCount: 10000000, rtp: 0, wildTableId: WILD_TABLE_01, H5Count: 1},
		},
		{
			"亂數決定覆蓋 Wild 數量: WILD_TABLE_01 H5Count 2",
			&SlotProb{},
			args{totalCount: 10000000, rtp: 0, wildTableId: WILD_TABLE_01, H5Count: 2},
		},
		{
			"亂數決定覆蓋 Wild 數量: WILD_TABLE_01 H5Count 3",
			&SlotProb{},
			args{totalCount: 10000000, rtp: 0, wildTableId: WILD_TABLE_01, H5Count: 3},
		},
		{
			"亂數決定覆蓋 Wild 數量: WILD_TABLE_02 H5Count 1",
			&SlotProb{},
			args{totalCount: 10000000, rtp: 0, wildTableId: WILD_TABLE_02, H5Count: 1},
		},
		{
			"亂數決定覆蓋 Wild 數量: WILD_TABLE_02 H5Count 2",
			&SlotProb{},
			args{totalCount: 10000000, rtp: 0, wildTableId: WILD_TABLE_02, H5Count: 2},
		},
		{
			"亂數決定覆蓋 Wild 數量: WILD_TABLE_02 H5Count 3",
			&SlotProb{},
			args{totalCount: 10000000, rtp: 0, wildTableId: WILD_TABLE_02, H5Count: 3},
		},
		{
			"亂數決定覆蓋 Wild 數量: WILD_TABLE_03 H5Count 1",
			&SlotProb{},
			args{totalCount: 10000000, rtp: 0, wildTableId: WILD_TABLE_03, H5Count: 1},
		},
		{
			"亂數決定覆蓋 Wild 數量: WILD_TABLE_03 H5Count 2",
			&SlotProb{},
			args{totalCount: 10000000, rtp: 0, wildTableId: WILD_TABLE_03, H5Count: 2},
		},
		{
			"亂數決定覆蓋 Wild 數量: WILD_TABLE_03 H5Count 3",
			&SlotProb{},
			args{totalCount: 10000000, rtp: 0, wildTableId: WILD_TABLE_03, H5Count: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var countMap = make(map[int]int)
			for i := 0; i < tt.args.totalCount; i++ {
				var wildCount = tt.slot.RandWildCount(tt.args.rtp, tt.args.wildTableId, tt.args.H5Count)
				countMap[wildCount]++
			}
			var keys = Utils.SortedMapKeys(countMap)
			for _, key := range keys {
				fmt.Printf("[%d] Count: %d (%s)\n", key, countMap[key], Utils.GetPercentage(float64(countMap[key]), float64(tt.args.totalCount)))
			}
		})
	}
}

func TestSlotProbRemoveMaxH5Score(t *testing.T) {
	type args struct {
		scoreArray [][]int
	}
	tests := []struct {
		name  string
		slot  *SlotProb
		args  args
		want  [][]int
		want1 int
	}{
		// Test cases
		{
			"移除 H5 最高分 20",
			&SlotProb{},
			args{scoreArray: [][]int{{5, 20, 10}, {5, 2}, {}, {2}, {}}},
			[][]int{{5, 10}, {5, 2}, {}, {2}, {}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.slot.RemoveMaxH5Score(tt.args.scoreArray)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlotProb.RemoveMaxH5Score() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SlotProb.RemoveMaxH5Score() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSlotProbCalculateWin(t *testing.T) {
	type args struct {
		rtp      int
		buyType  int
		debugCmd *DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"計算連線獎金: 一般投注",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_NONE},
		},
		{
			"計算連線獎金: 額外投注",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_EXTRA_BET},
		},
		{
			"計算連線獎金: 購買免費",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_FREE_SPINS},
		},
		{
			"計算連線獎金: 購買超級",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_SUPER_FREE_SPINS},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var gameSymbol, groupIdx, reelIndex, scoreArray, _ = tt.slot.RandMGSymbol(tt.args.rtp, tt.args.buyType, 1, tt.args.debugCmd)
			ShowGameSymbol(gameSymbol)
			fmt.Printf("GroupIdx: %d, ReelIndex: %v\n", groupIdx, reelIndex)
			fmt.Printf("H5CountList: %v, H5ScoreArray: %v\n", tt.slot.GetH5CountList(gameSymbol), scoreArray)
			var totalWin, lineSymbolList, lineCountList, lineWinList = tt.slot.CalculateWin(1, gameSymbol)
			fmt.Printf("TotalWin: %d\n", totalWin)
			fmt.Printf("LineSymbolList: %v\n", lineSymbolList)
			fmt.Printf("LineCountList: %v\n", lineCountList)
			fmt.Printf("LineWinList: %v\n", lineWinList)
		})
	}
}

func TestSlotProbCheckLine(t *testing.T) {
	type args struct {
		lineSymbol []Symbol
	}
	tests := []struct {
		name       string
		slot       *SlotProb
		args       args
		wantSymbol Symbol
		wantCount  int
	}{
		// Test cases
		{
			"檢查連線及獎圖: [WW, WW, WW, WW, WW]",
			&SlotProb{},
			args{lineSymbol: []Symbol{WW, WW, WW, WW, WW}},
			WW,
			5,
		},
		{
			"檢查連線及獎圖: [WW, WW, WW, WW, H4]",
			&SlotProb{},
			args{lineSymbol: []Symbol{WW, WW, WW, WW, H4}},
			H4,
			5,
		},
		{
			"檢查連線及獎圖: [WW, WW, WW, WW, H5]",
			&SlotProb{},
			args{lineSymbol: []Symbol{WW, WW, WW, WW, H5}},
			WW,
			4,
		},
		{
			"檢查連線及獎圖: [WW, WW, WW, H5, WW]",
			&SlotProb{},
			args{lineSymbol: []Symbol{WW, WW, WW, H5, WW}},
			H5,
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSymbol, gotCount := tt.slot.CheckLine(tt.args.lineSymbol)
			if !reflect.DeepEqual(gotSymbol, tt.wantSymbol) {
				t.Errorf("SlotProb.CheckLine() gotSymbol = %v, want %v", gotSymbol, tt.wantSymbol)
			}
			if gotCount != tt.wantCount {
				t.Errorf("SlotProb.CheckLine() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestSlotProbIsRespin(t *testing.T) {
	type args struct {
		gameSymbol GameSymbol
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
		want bool
	}{
		// Test cases
		{
			"檢查是否符合重轉條件 符合",
			&SlotProb{},
			args{gameSymbol: GameSymbol{{LA, LT, H1}, {LJ, LK, H5}, {LJ, SS, LK}, {H4, SS, H5}, {H3, H5, H5}}},
			true,
		},
		{
			"檢查是否符合重轉條件 不符合: 三個 Scatter",
			&SlotProb{},
			args{gameSymbol: GameSymbol{{LA, LT, H1}, {LJ, LK, H5}, {LJ, SS, LK}, {H4, SS, H5}, {H3, H5, SS}}},
			false,
		},
		{
			"檢查是否符合重轉條件 不符合: 一個 Scatter",
			&SlotProb{},
			args{gameSymbol: GameSymbol{{LA, LT, H1}, {LJ, LK, H5}, {LJ, SS, LK}, {H4, LJ, H5}, {H3, H5, H5}}},
			false,
		},
		{
			"檢查是否符合重轉條件 不符合: Scatter 在底部",
			&SlotProb{},
			args{gameSymbol: GameSymbol{{LA, LT, H1}, {LJ, LK, H5}, {LJ, SS, LK}, {H4, LJ, SS}, {H3, H5, H5}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slot.IsRespin(tt.args.gameSymbol); got != tt.want {
				t.Errorf("SlotProb.IsRespin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlotProbReplaceScatter(t *testing.T) {
	type args struct {
		rtp      int
		buyType  int
		debugCmd *DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"替換 Scatter 表演: 一般投注 Scatter 數量不足",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_NONE},
		},
		{
			"替換 Scatter 表演: 額外投注 Scatter 數量不足",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_EXTRA_BET},
		},
		{
			"替換 Scatter 表演: 購買免費",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_FREE_SPINS},
		},
		{
			"替換 Scatter 表演: 購買超級",
			&SlotProb{},
			args{rtp: 0, buyType: Common.BUY_SUPER_FREE_SPINS},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var gameSymbol, groupIdx, reelIndex, scoreArray, _ = tt.slot.RandMGSymbol(tt.args.rtp, tt.args.buyType, 1, tt.args.debugCmd)
			ShowGameSymbol(gameSymbol)
			fmt.Printf("GroupIdx: %d, ReelIndex: %v\n", groupIdx, reelIndex)
			fmt.Printf("H5CountList: %v, H5ScoreArray: %v\n", tt.slot.GetH5CountList(gameSymbol), scoreArray)
			fmt.Printf("-------------------------------  %14s  -------------------------------\n", "ReplaceScatter")
			var performanceSymbol = tt.slot.ReplaceScatter(gameSymbol)
			ShowGameSymbol(performanceSymbol)
		})
	}
}

func TestSlotProbReplaceWild(t *testing.T) {
	type args struct {
		rtp          int
		freeGameType int
		wildTableId  int
		debugCmd     *DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"替換 Wild 表演: FREE_GAME_01 WILD_TABLE_01",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_01},
		},
		{
			"替換 Wild 表演: FREE_GAME_01 WILD_TABLE_02",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_02},
		},
		{
			"替換 Wild 表演: FREE_GAME_01 WILD_TABLE_03",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_03},
		},
		{
			"替換 Wild 表演: FREE_GAME_02 WILD_TABLE_01",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_01},
		},
		{
			"替換 Wild 表演: FREE_GAME_02 WILD_TABLE_02",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_02},
		},
		{
			"替換 Wild 表演: FREE_GAME_02 WILD_TABLE_03",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_03},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var reelGroup = FGReelGroup[tt.args.rtp][tt.args.freeGameType]
			var gameSymbol, reelIndex, scoreArray = tt.slot.RandFGSymbol(reelGroup, tt.args.rtp, tt.args.freeGameType, 1, tt.args.wildTableId, tt.args.debugCmd)
			ShowGameSymbol(gameSymbol)
			fmt.Printf("FreeGameType: %d, ReelIndex: %v\n", tt.args.freeGameType, reelIndex)
			fmt.Printf("H5CountList: %v, H5ScoreArray: %v\n", tt.slot.GetH5CountList(gameSymbol), scoreArray)
			fmt.Printf("--------------------------------  %12s  --------------------------------\n", "Replace Wild")
			var performanceSymbol = tt.slot.ReplaceWild(gameSymbol, tt.args.freeGameType)
			ShowGameSymbol(performanceSymbol)
		})
	}
}

func TestSlotProbReplaceH5(t *testing.T) {
	type args struct {
		rtp          int
		freeGameType int
		wildTableId  int
		debugCmd     *DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"替換 H5 表演: FREE_GAME_01 WILD_TABLE_01",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_01},
		},
		{
			"替換 H5 表演: FREE_GAME_01 WILD_TABLE_02",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_02},
		},
		{
			"替換 H5 表演: FREE_GAME_01 WILD_TABLE_03",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_03},
		},
		{
			"替換 H5 表演: FREE_GAME_02 WILD_TABLE_01",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_01},
		},
		{
			"替換 H5 表演: FREE_GAME_02 WILD_TABLE_02",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_02},
		},
		{
			"替換 H5 表演: FREE_GAME_02 WILD_TABLE_03",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_03},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var reelGroup = FGReelGroup[tt.args.rtp][tt.args.freeGameType]
			var gameSymbol, reelIndex, scoreArray = tt.slot.RandFGSymbol(reelGroup, tt.args.rtp, tt.args.freeGameType, 1, tt.args.wildTableId, tt.args.debugCmd)
			ShowGameSymbol(gameSymbol)
			fmt.Printf("FreeGameType: %d, ReelIndex: %v\n", tt.args.freeGameType, reelIndex)
			fmt.Printf("H5CountList: %v, H5ScoreArray: %v\n", tt.slot.GetH5CountList(gameSymbol), scoreArray)
			fmt.Printf("--------------------------------  %12s  --------------------------------\n", " Replace H5 ")
			var performanceSymbol = tt.slot.ReplaceH5(gameSymbol, tt.args.freeGameType)
			ShowGameSymbol(performanceSymbol)
		})
	}
}

func TestSlotProbReplaceOthers(t *testing.T) {
	type args struct {
		rtp          int
		freeGameType int
		wildTableId  int
		debugCmd     *DebugCmd
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"替換 Wild 以外的獎圖表演: FREE_GAME_01 WILD_TABLE_01",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_01},
		},
		{
			"替換 Wild 以外的獎圖表演: FREE_GAME_01 WILD_TABLE_02",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_02},
		},
		{
			"替換 Wild 以外的獎圖表演: FREE_GAME_01 WILD_TABLE_03",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_01, wildTableId: WILD_TABLE_03},
		},
		{
			"替換 Wild 以外的獎圖表演: FREE_GAME_02 WILD_TABLE_01",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_01},
		},
		{
			"替換 Wild 以外的獎圖表演: FREE_GAME_02 WILD_TABLE_02",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_02},
		},
		{
			"替換 Wild 以外的獎圖表演: FREE_GAME_02 WILD_TABLE_03",
			&SlotProb{},
			args{rtp: 0, freeGameType: FREE_GAME_02, wildTableId: WILD_TABLE_03},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			var reelGroup = FGReelGroup[tt.args.rtp][tt.args.freeGameType]
			var gameSymbol, reelIndex, scoreArray = tt.slot.RandFGSymbol(reelGroup, tt.args.rtp, tt.args.freeGameType, 1, tt.args.wildTableId, tt.args.debugCmd)
			ShowGameSymbol(gameSymbol)
			fmt.Printf("FreeGameType: %d, ReelIndex: %v\n", tt.args.freeGameType, reelIndex)
			fmt.Printf("H5CountList: %v, H5ScoreArray: %v\n", tt.slot.GetH5CountList(gameSymbol), scoreArray)
			fmt.Printf("-------------------------------  %14s  -------------------------------\n", "Replace Others")
			var performanceSymbol = tt.slot.ReplaceOthers(gameSymbol, tt.args.freeGameType)
			ShowGameSymbol(performanceSymbol)
		})
	}
}

func TestSlotProbRandCoverSymbol(t *testing.T) {
	type args struct {
		excludeSymbol ReelSymbol
		freeGameType  int
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
	}{
		// Test cases
		{
			"亂數產生用來覆蓋的獎圖 主遊戲 排除 H2, H3, H4",
			&SlotProb{},
			args{excludeSymbol: ReelSymbol{H2, H3, H4}, freeGameType: -1},
		},
		{
			"亂數產生用來覆蓋的獎圖 FREE_GAME_01 排除 LJ, LK, LQ, LJ, LT",
			&SlotProb{},
			args{excludeSymbol: ReelSymbol{LA, LK, LQ, LJ, LT}, freeGameType: FREE_GAME_01},
		},
		{
			"亂數產生用來覆蓋的獎圖 FREE_GAME_02 回傳 NN",
			&SlotProb{},
			args{excludeSymbol: ReelSymbol{}, freeGameType: FREE_GAME_02},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var symbol = tt.slot.RandCoverSymbol(tt.args.excludeSymbol, tt.args.freeGameType)
			fmt.Printf("RandCoverSymbol: Symbol = %v\n", symbol)
		})
	}
}
