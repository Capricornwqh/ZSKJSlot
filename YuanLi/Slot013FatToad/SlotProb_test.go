package Slot013FatToad

import (
	"Force/GameServer/Common"
	"encoding/csv"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"testing"
)

type args struct {
	rtp          int
	lineBet      int
	result       *SlotResult
	debugCmdList []DebugCmd
	symbol       GameSymbol
}

type argCmd struct {
	name string
	slot *SlotProb
	args args
}

func TestRunMainGame(t *testing.T) {

	// tests := []struct {
	// 	name string
	// 	slot *SlotProb
	// 	args args
	// }{
	// 	// Test cases
	// 	// {
	// 	// 	"進行遊戲: 一般投注",
	// 	// 	&SlotProb{},
	// 	// 	args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_NONE}},
	// 	// },
	// 	{
	// 		"進行遊戲: 購買免費",
	// 		&SlotProb{},
	// 		args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_FREE_SPINS}},
	// 	},
	// 	// {
	// 	// 	"進行遊戲: 購買超級",
	// 	// 	&SlotProb{},
	// 	// 	args{rtp: 0, lineBet: 1, result: &SlotResult{BuyType: Common.BUY_SUPER_FREE_SPINS}},
	// 	// },
	// }

	tests := []argCmd{}
	for range 1000 {
		tests = append(tests, argCmd{
			"進行遊戲: 購買免費",
			&SlotProb{},
			args{
				rtp:     0,
				lineBet: 1,
				result:  &SlotResult{BuyType: Common.BUY_FREE_SPINS},
				symbol: GameSymbol{
					{2, 4, 6, 3, 4, 2},
					{3, 5, 2, 4, 7, 6},
					{6, 1, 2, 4, 3, 8},
					{2, 8, 4, 5, 2, 1},
					{11, 5, 6, 2, 6, 4},
					{2, 8, 3, 4, 2, 4},
				},
				debugCmdList: []DebugCmd{
					// {DebugData: []int{0, -1, -1, -1, -1, -1, 3}},

					// {DebugData: []int{0, 89, 12, 101, 95, 103, 41}},
					// {DebugData: []int{0, 58, 70, 27, 83, 2, 42}},
					// {DebugData: []int{0, 17, 20, 28, 3, 64, 37}},
					// {DebugData: []int{0, 67, 5, 111, 3, 108, 61}},
					// {DebugData: []int{0, 74, 92, 63, 56, 80, 75}},
					// {DebugData: []int{0, 70, 61, 70, 9, 69, 43}},
					// {DebugData: []int{0, 90, 91, 39, 61, 18, 54}},
					// {DebugData: []int{0, 8, 5, 71, 7, 35, 65}},

					// {DebugData: []int{0, 100, 97, 88, 62, 20, 0}},
					// {DebugData: []int{0, 77, 31, 28, 53, 87, 79}},
					// {DebugData: []int{0, 10, 111, 22, 35, 39, 5}},
					// {DebugData: []int{0, 89, 73, 74, 74, 47, 33}},
					// {DebugData: []int{0, 83, 94, 53, 68, 53, 46}},
					// {DebugData: []int{0, 29, 66, 100, 64, 5, 77}},
					// {DebugData: []int{0, 23, 73, 99, 60, 46, 84}},
					// {DebugData: []int{0, 102, 103, 28, 63, 107, 74}},
					// {DebugData: []int{0, 7, 74, 43, 65, 58, 36}},
					// {DebugData: []int{0, 57, 58, 61, 87, 102, 84}},
					// {DebugData: []int{0, 37, 50, 41, 39, 0, 26}},
					// {DebugData: []int{0, 11, 56, 92, 44, 98, 90}},
					// {DebugData: []int{0, 9, 96, 10, 8, 61, 40}},
				},
			},
		})
	}
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatalf("无法创建日志文件: %v", err)
	}
	defer logFile.Close()

	reelIndexFile, err := os.OpenFile("reelindex.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatalf("无法创建reelindex文件: %v", err)
	}
	defer reelIndexFile.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slot.Init()
			tt.slot.Run(tt.args.rtp, tt.args.lineBet, tt.args.result, tt.args.debugCmdList)

			// 输出主要日志到log.txt文件
			for i, tumble := range tt.args.result.MGTumbleList {
				ShowGameSymbol(logFile, tumble.TumbleSymbol)
				fmt.Fprintf(logFile, "%d) GroupIdx: %d, ReelIndex: %v\n", i, tt.args.result.MGGroupIndex, tumble.ReelIndex)
				fmt.Fprintf(logFile, "%d) SSCount: %d\n", i, tumble.SSCount)
				fmt.Fprintf(logFile, "%d) MGPerformanceType: %d\n", i, tumble.MGPerformanceType)
				fmt.Fprintf(logFile, "%d) LineSymbol: %v\n", i, tumble.LineSymbol)
				fmt.Fprintf(logFile, "%d) LineCount: %v\n", i, tumble.LineCount)
				fmt.Fprintf(logFile, "%d) LineWin: %v\n", i, tumble.LineWin)
				fmt.Fprintf(logFile, "%d) Win: %v\n", i, tumble.Win)
				fmt.Fprintf(logFile, "===============main game=====================\n")
			}
			fmt.Fprintf(logFile, "MainWin: %d\n", tt.args.result.MainWin)
			fmt.Fprintf(logFile, "FGSpinCount: %v\n", tt.args.result.FGSpinCount)

			for i, spin := range tt.args.result.FGSpinList {
				ShowGameSymbol(logFile, spin.TumbleSymbol)
				fmt.Fprintf(logFile, "%d] FreeGameType: %d, ReelIndex: %v\n", i, tt.args.result.FreeGameType, spin.ReelIndex)
				fmt.Fprintf(logFile, "%d] SSCount: %d\n", i, spin.SSCount)
				fmt.Fprintf(logFile, "%d] PerformanceType: %v\n", i, spin.PerformanceType)
				fmt.Fprintf(logFile, "%d] LineSymbol: %v\n", i, spin.LineSymbol)
				fmt.Fprintf(logFile, "%d] LineCount: %v\n", i, spin.LineCount)
				fmt.Fprintf(logFile, "%d] LineWin: %v\n", i, spin.LineWin)
				fmt.Fprintf(logFile, "%d] Multiplier: %d\n", i, spin.Multiplier)
				fmt.Fprintf(logFile, "%d] Win: %d\n", i, spin.SpinWin)
				fmt.Fprintf(logFile, "%d] Stage: %d\n", i, spin.Stage)
				fmt.Fprintf(logFile, "=============free game=======================\n")
			}
			fmt.Fprintf(logFile, "FGCumWildCount: %d\n", tt.args.result.FGCumWildCount)
			fmt.Fprintf(logFile, "WWMultiplier: %d\n", tt.args.result.WWMultiplier)
			fmt.Fprintf(logFile, "FGIndex: %d\n", tt.args.result.FGIndex)
			fmt.Fprintf(logFile, "FreeWin: %d\n", tt.args.result.FreeWin)
			fmt.Fprintf(logFile, "TotalWin: %d\n", tt.args.result.TotalWin)
			fmt.Fprintf(logFile, "\n\n\n\n\n")

			// 输出reelindex信息到reelindex.txt文件
			fmt.Fprintf(reelIndexFile, "[\n")
			for _, spin := range tt.args.result.FGSpinList {
				if len(spin.ReelIndex) > 0 {
					reelIndexStr := make([]string, len(spin.ReelIndex))
					for i, v := range spin.ReelIndex {
						reelIndexStr[i] = fmt.Sprintf("%d", v)
					}
					fmt.Fprintf(reelIndexFile, "[[%s],%d],\n", strings.Join(reelIndexStr, ","), spin.SpinWin)
				}
			}
			fmt.Fprintf(reelIndexFile, "],\n")
		})
	}
}

func TestCheckLine(t *testing.T) {
	type args struct {
		lineBet int
		line    []Symbol
	}
	tests := []struct {
		name string
		slot *SlotProb
		args args
		want int
	}{
		// {
		// 	name: "測試連線",
		// 	slot: &SlotProb{},
		// 	args: args{lineBet: 1, line: []Symbol{LK, LK, LK, LK, LA, SS2}},
		// 	want: 1,
		// },
		// {
		// 	name: "測試連線",
		// 	slot: &SlotProb{},
		// 	args: args{lineBet: 1, line: []Symbol{LJ, LJ, LJ, LJ, SS, H2}},
		// 	want: 1,
		// },
		{
			name: "測試連線",
			slot: &SlotProb{},
			args: args{lineBet: 1, line: []Symbol{WW, WW, WW, WW, H2, H3}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			symbal, count, multiplier := tt.slot.CheckLine(1, tt.args.line)
			fmt.Printf("name: %+v, %v, %d, %d\n", tt.args.line, symbal, count, multiplier)
		})
	}
}

func TestReel(t *testing.T) {
	// 打开CSV文件
	file, err := os.Open("D:\\Telegram\\free_game_probabilities(3).csv")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建CSV reader
	reader := csv.NewReader(file)

	// 读取所有记录
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("读取CSV失败:", err)
		return
	}

	// 确保有足够的行
	if len(records) < 2 {
		fmt.Println("CSV文件格式不正确")
		return
	}

	// 确定列数（卷轴数量）
	numReels := len(records[0]) - 1 // 减去Symbol列
	if numReels <= 0 {
		fmt.Println("未找到卷轴数据")
		return
	}

	// 创建卷轴数组
	reels := make([][]string, numReels)
	for i := range reels {
		reels[i] = []string{}
	}

	// 处理每一行（每个符号）
	for i := 1; i < len(records); i++ {
		symbol := records[i][0]

		// 处理每个卷轴
		for j := 1; j <= numReels; j++ {
			if len(records[i]) > j {
				count, err := strconv.Atoi(records[i][j])
				if err != nil {
					fmt.Printf("无法解析数字 %s: %v\n", records[i][j], err)
					continue
				}

				// 将符号添加到卷轴中指定的次数
				for range count {
					reels[j-1] = append(reels[j-1], symbol)
				}
			}
		}
	}

	// 输出结果
	// fmt.Printf("%+v\n", reels)
	for i, reel := range reels {
		// 随机打乱卷轴
		rand.Shuffle(len(reel), func(i, j int) {
			reel[i], reel[j] = reel[j], reel[i]
		})
		fmt.Printf("Reel_%d: %v (共 %d 个符号)\n", i+1, strings.Join(reel, ","), len(reel))
	}
}

func TestFindTargetCoin(t *testing.T) {
	slot := &SlotProb{}
	result := &SlotResult{
		WWWild: []*WildStruct{
			{WildSymbol: WW, WildCoordinate: [2]int{3, 1}},
		},
		SSWild: []*WildStruct{
			{WildSymbol: SS, WildCoordinate: [2]int{0, 3}},
			{WildSymbol: SS, WildCoordinate: [2]int{1, 4}},
			{WildSymbol: SS, WildCoordinate: [2]int{2, 5}},
			{WildSymbol: SS, WildCoordinate: [2]int{3, 4}},
		},
	}
	targetCoin, wwPos := slot.findTargetCoin(result.WWWild, result.SSWild)
	for i := 0; i < len(targetCoin) && i < len(wwPos); i++ {
		fmt.Printf("SSWild: %v, WWWild: %v\n", targetCoin[i].WildCoordinate, wwPos[i].WildCoordinate)
	}
	//移动向量
	moveVector := make([][2]int, len(targetCoin))
	for i := 0; i < len(targetCoin) && i < len(wwPos); i++ {
		moveVector[i] = [2]int{targetCoin[i].WildCoordinate[0] - wwPos[i].WildCoordinate[0],
			targetCoin[i].WildCoordinate[1] - wwPos[i].WildCoordinate[1]}
		fmt.Printf("Move Vector: %v\n", moveVector[i])
	}
	// 更新金蟾
	for i := 0; i < len(targetCoin) && i < len(wwPos); i++ {
		for _, v := range result.WWWild {
			v.WildCoordinate[0] += moveVector[i][0]
			v.WildCoordinate[1] += moveVector[i][1]
			fmt.Printf("%v, ", v.WildCoordinate)
		}
		// 更新向量
		for j := i + 1; j < len(moveVector); j++ {
			moveVector[j][0] -= moveVector[i][0]
			moveVector[j][1] -= moveVector[i][1]
			fmt.Printf("Updated Move Vector[%d]: %v\n", j, moveVector[j])
		}
		fmt.Println()
	}
}

func TestExpandFrog(t *testing.T) {
	slot := &SlotProb{}
	result := &SlotResult{
		WWWild: []*WildStruct{
			{WildSymbol: WW, WildCoordinate: [2]int{4, 2}},
			{WildSymbol: WW, WildCoordinate: [2]int{4, 3}},
			{WildSymbol: WW, WildCoordinate: [2]int{5, 2}},
			{WildSymbol: WW, WildCoordinate: [2]int{5, 3}},
		},
	}
	result.WWWild = slot.expandFrog(result.WWWild, 3)
	for _, v := range result.WWWild {
		fmt.Printf("%v, ", v.WildCoordinate)
	}
	fmt.Println()

	result = &SlotResult{
		WWWild: []*WildStruct{
			{WildSymbol: WW, WildCoordinate: [2]int{1, 1}},
			{WildSymbol: WW, WildCoordinate: [2]int{1, 2}},
			{WildSymbol: WW, WildCoordinate: [2]int{1, 2}},
			{WildSymbol: WW, WildCoordinate: [2]int{2, 1}},
			{WildSymbol: WW, WildCoordinate: [2]int{2, 2}},
			{WildSymbol: WW, WildCoordinate: [2]int{2, 3}},
			{WildSymbol: WW, WildCoordinate: [2]int{3, 1}},
			{WildSymbol: WW, WildCoordinate: [2]int{3, 2}},
			{WildSymbol: WW, WildCoordinate: [2]int{3, 3}},
		},
	}
	result.WWWild = slot.expandFrog(result.WWWild, 4)
	for _, v := range result.WWWild {
		fmt.Printf("%v, ", v.WildCoordinate)
	}
	fmt.Println()
}

// TestFGReelGroup 测试免费游戏轮带盘面
func (slot *SlotProb) TestFGReelGroup(rtp, groupIdx int, reelIndices []int) {
	if reelIndices == nil || len(reelIndices) < SLOT_COL {
		fmt.Println("无效的reelIndices参数，请提供", SLOT_COL, "个索引值")
		return
	}

	// 获取对应的轮带组
	reelGroup := FGReelGroup[rtp][groupIdx]

	// 创建盘面
	gameSymbol := make(GameSymbol, SLOT_COL)

	// 对每一列生成符号
	for col := range SLOT_COL {
		// 获取当前列的轮带
		reel := reelGroup[col]
		reelLength := len(reel)
		dice := reelIndices[col] // 使用提供的reelindex

		// 创建列符号
		columnSymbol := make(ReelSymbol, SLOT_ROW)
		for row := range SLOT_ROW {
			idx := dice + row
			if idx < reelLength {
				columnSymbol[row] = reel[idx]
			} else {
				columnSymbol[row] = reel[idx-reelLength]
			}
		}
		gameSymbol[col] = columnSymbol
	}

	// 输出参数信息
	fmt.Println("RTP:", rtp, "GroupIdx:", groupIdx)
	fmt.Println("ReelIndices:", reelIndices)
	fmt.Println("---------- 盘面结果 -----------")
	ShowGameSymbol(os.Stdout, gameSymbol)

	slot.TestCalculateWin(gameSymbol, 1)
}

func TestFGReelGroupWithIndices(t *testing.T) {
	var slot SlotProb
	slot.Init()
	slot.TestFGReelGroup(0, 0, []int{27, 31, 86, 18, 14, 101})
	// slot.TestFGReelGroup(0, 0, []int{80, 99, 33, 47, 27, 52})
}

func TestResult(t *testing.T) {
	slot := &SlotProb{}
	spinResult := &SpinResult{
		TumbleSymbol: GameSymbol{
			{H1, LQ, WW, WW, H4, H4},
			{LK, H3, WW, WW, LJ, LJ},
			{H2, H2, LQ, LQ, SS, LK},
			{H4, LQ, LQ, LQ, LJ, LJ},
			{LA, SS, LK, LK, H3, H4},
			{LA, LA, H1, H1, LQ, LQ},
		},
	}
	result := &SlotResult{}
	slot.CalculateFreeWin(1, spinResult, result)
	fmt.Printf("LineSymbol: %v\n", spinResult.LineSymbol)
	fmt.Printf("LineCount: %v\n", spinResult.LineCount)
	fmt.Printf("LineWin: %v\n", spinResult.LineWin)
	fmt.Printf("Win: %v\n", spinResult.SpinWin)

	// spinResult = &SpinResult{
	// 	TumbleSymbol: GameSymbol{
	// 		{LJ, LJ, H1, H1, H1, LA},
	// 		{H1, WW, WW, WW, WW, H2},
	// 		{LK, WW, WW, WW, WW, H3},
	// 		{LA, WW, WW, WW, WW, LQ},
	// 		{H4, WW, WW, WW, WW, H2},
	// 		{LJ, LJ, H3, H4, LA, LA},
	// 	},
	// }
	// result = &SlotResult{}
	// slot.CalculateFreeWin(1, spinResult, result)
	// fmt.Printf("LineSymbol: %v\n", spinResult.LineSymbol)
	// fmt.Printf("LineCount: %v\n", spinResult.LineCount)
	// fmt.Printf("LineWin: %v\n", spinResult.LineWin)
	// fmt.Printf("Win: %v\n", spinResult.SpinWin)
}

// TestCalculateWin 测试盘面赢分计算
func (slot *SlotProb) TestCalculateWin(gameSymbol GameSymbol, lineBet int) {

	spinResult := &SpinResult{
		TumbleSymbol: gameSymbol,
	}
	result := &SlotResult{}
	// 计算赢分
	slot.CalculateFreeWin(lineBet, spinResult, result)

	// 打印计算结果
	fmt.Println("====== 计算结果 ======")
	fmt.Printf("SpinWin: %d\n", spinResult.SpinWin)
	fmt.Printf("LineSymbol: %v\n", spinResult.LineSymbol)
	fmt.Printf("LineCount: %v\n", spinResult.LineCount)
	fmt.Printf("LineWin: %v\n", spinResult.LineWin)
}
