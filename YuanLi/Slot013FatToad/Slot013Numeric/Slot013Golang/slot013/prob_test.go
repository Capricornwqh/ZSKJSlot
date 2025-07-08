package slot013

import (
	"fmt"
	"strings"
	"testing"
)

type argCmd struct {
	name string
	slot *SlotProb
	args args
}

type args struct {
	rtp          int
	lineBet      int
	result       *SlotResult
	debugCmdList []DebugCmd
}

func TestRunMainGame(t *testing.T) {
	tests := []argCmd{}
	for range 100 {
		tests = append(tests, argCmd{
			name: "進行遊戲: 一般投注",
			slot: &SlotProb{},
			args: args{
				rtp:     0,
				lineBet: 1,
				result:  &SlotResult{BuyType: BUY_NONE},
				debugCmdList: []DebugCmd{
					{DebugData: []int{0, 12, 76, 31, 56, 43, 23, 3, 2, 3, 5, -1, -1}},
					{DebugData: []int{-1, -1, -1, -1, -1, -1, -1, -1, 0, 89, 12, 101, 95, 103, 41, 2, 3, 5, -1, -1, -1}},
				},
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			tt.slot.Run(tt.args.rtp, tt.args.lineBet, tt.args.result, tt.args.debugCmdList)

			for i, tumble := range tt.args.result.MGTumbleList {
				ShowGameSymbol(tumble.TumbleSymbol)
				fmt.Printf("%d) GroupIdx: %d, ReelIndex: %v\n", i, tt.args.result.MGGroupIndex, tumble.ReelIndex)
				fmt.Printf("%d) SSCount: %d\n", i, tumble.SSCount)
				fmt.Printf("%d) MGPerformanceType: %d\n", i, tumble.MGPerformanceType)
				fmt.Printf("%d) LineSymbol: %v\n", i, tumble.LineSymbol)
				fmt.Printf("%d) LineCount: %v\n", i, tumble.LineCount)
				fmt.Printf("%d) LineWin: %v\n", i, tumble.LineWin)
				fmt.Printf("%d) Win: %v\n", i, tumble.Win)
				fmt.Printf("===============main game=====================\n")
			}
			fmt.Printf("MainWin: %d\n", tt.args.result.MainWin)
			fmt.Printf("FGSpinCount: %v\n", tt.args.result.FGSpinCount)

			for i, spin := range tt.args.result.FGSpinList {
				ShowGameSymbol(spin.TumbleSymbol)
				fmt.Printf("%d] FreeGameType: %d, ReelIndex: %v\n", i, tt.args.result.FreeGameType, spin.ReelIndex)
				fmt.Printf("%d] SSCount: %d\n", i, spin.SSCount)
				fmt.Printf("%d] PerformanceType: %v\n", i, spin.PerformanceType)
				fmt.Printf("%d] LineSymbol: %v\n", i, spin.LineSymbol)
				fmt.Printf("%d] LineCount: %v\n", i, spin.LineCount)
				fmt.Printf("%d] LineWin: %v\n", i, spin.LineWin)
				fmt.Printf("%d] Multiplier: %d\n", i, spin.Multiplier)
				fmt.Printf("%d] Win: %d\n", i, spin.SpinWin)
				fmt.Printf("%d] Stage: %d\n", i, spin.Stage)
				fmt.Printf("=============free game=======================\n")
			}
			fmt.Printf("FGCumWildCount: %d\n", tt.args.result.FGCumWildCount)
			fmt.Printf("WWMultiplier: %d\n", tt.args.result.WWMultiplier)
			fmt.Printf("FGIndex: %d\n", tt.args.result.FGIndex)
			fmt.Printf("FreeWin: %d\n", tt.args.result.FreeWin)
			fmt.Printf("TotalWin: %d\n", tt.args.result.TotalWin)
			fmt.Printf("\n\n\n\n\n")

			fmt.Printf("[\n")
			for _, spin := range tt.args.result.FGSpinList {
				if len(spin.ReelIndex) > 0 {
					reelIndexStr := make([]string, len(spin.ReelIndex))
					for i, v := range spin.ReelIndex {
						reelIndexStr[i] = fmt.Sprintf("%d", v)
					}
					fmt.Printf("[[%s],%d],\n", strings.Join(reelIndexStr, ","), spin.SpinWin)
				}
			}
			fmt.Printf("],\n")
		})
	}
}
