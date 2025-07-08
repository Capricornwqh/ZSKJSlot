package service_cargod

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// 统计数据结构体
type CountData struct {
	CountBetList      []int64
	CountMainWinList  []int64
	CountEventWinList []int64
	CountJackpotWin   int64
	EventTriggerList  []int
	// 新增：按levelIndex统计的数据
	LevelBetList      []int64 // 每个level的总投注
	LevelMainWinList  []int64 // 每个level的主游戏奖金
	LevelEventWinList []int64 // 每个level的特殊玩法奖金
	LevelJackpotWin   []int64 // 每个level的累积奖池奖金
	LevelHitCount     []int   // 每个level的命中次数
}

// 测试车神游戏的RTP
func TestCarRTP(t *testing.T) {
	unitTimes := 1000000000
	numProcesses := 10

	startTime := time.Now()

	// 创建channel用于收集结果
	resultChan := make(chan CountData, numProcesses)
	var wg sync.WaitGroup

	// 启动goroutines
	for i := 1; i <= numProcesses; i++ {
		wg.Add(1)
		go runTestWithChannel(unitTimes, i, resultChan, &wg)
	}

	// 等待所有goroutines完成
	wg.Wait()
	close(resultChan)

	// 统计结果
	var allBet, allMainWin, allEventWin, allJackpotWin int64
	levelBetList := make([]int64, 36) // 36个level
	levelMainWinList := make([]int64, 36)
	levelEventWinList := make([]int64, 36)
	levelJackpotWin := make([]int64, 36)
	levelHitCount := make([]int, 36)

	// 从channel读取所有结果
	for countData := range resultChan {
		for _, bet := range countData.CountBetList {
			allBet += bet
		}
		for _, win := range countData.CountMainWinList {
			allMainWin += win
		}
		for _, win := range countData.CountEventWinList {
			allEventWin += win
		}
		allJackpotWin += countData.CountJackpotWin

		// 累加按level统计的数据
		for i := range 36 {
			levelBetList[i] += countData.LevelBetList[i]
			levelMainWinList[i] += countData.LevelMainWinList[i]
			levelEventWinList[i] += countData.LevelEventWinList[i]
			levelJackpotWin[i] += countData.LevelJackpotWin[i]
			levelHitCount[i] += countData.LevelHitCount[i]
		}
	}

	elapsed := time.Since(startTime)
	hours := int(elapsed.Hours())
	minutes := int(elapsed.Minutes()) % 60
	seconds := int(elapsed.Seconds()) % 60
	p := message.NewPrinter(language.Chinese)
	p.Printf("Test %d次 耗时: %d时%d分%d秒\n", unitTimes*numProcesses, hours, minutes, seconds)

	// 总体RTP
	fmt.Printf("=== 总体RTP ===\n")
	fmt.Printf("Total RTP: %.6f\n", float64(allMainWin+allEventWin+allJackpotWin)/float64(allBet))
	fmt.Printf("Main RTP: %.6f\n", float64(allMainWin)/float64(allBet))
	fmt.Printf("Event RTP: %.6f\n", float64(allEventWin)/float64(allBet))
	fmt.Printf("Jackpot RTP: %.6f\n", float64(allJackpotWin)/float64(allBet))

	// 按Level统计RTP
	fmt.Printf("\n=== 按Level统计RTP ===\n")
	fmt.Printf("Level | 命中次数 | 投注金额 | Main RTP | Event RTP | Jackpot RTP | Total RTP\n")
	fmt.Printf("------|----------|----------|-----------|-------------|-------------|------\n")

	for i := range 36 {
		if levelBetList[i] > 0 {
			mainRTP := float64(levelMainWinList[i]) / float64(levelBetList[i])
			eventRTP := float64(levelEventWinList[i]) / float64(levelBetList[i])
			jackpotRTP := float64(levelJackpotWin[i]) / float64(levelBetList[i])
			totalRTP := mainRTP + eventRTP + jackpotRTP

			p.Printf("%5d | %8d | %9d | %9.6f | %11.6f | %11.6f | %.6f\n",
				i, levelHitCount[i], levelBetList[i], mainRTP, eventRTP, jackpotRTP, totalRTP)
		}
	}
}

// 新的运行测试函数，使用channel返回结果
func runTestWithChannel(unitTimes int, processID int, resultChan chan<- CountData, wg *sync.WaitGroup) {
	defer wg.Done()

	// 统计数据
	countBetList := make([]int64, 12)
	countMainWinList := make([]int64, 12)
	countEventWinList := make([]int64, 12)
	var countJackpotWin int64 = 0
	eventTriggerList := make([]int, 23)

	// 新增：按level统计的数据
	levelBetList := make([]int64, 36) // 36个level
	levelMainWinList := make([]int64, 36)
	levelEventWinList := make([]int64, 36)
	levelJackpotWin := make([]int64, 36)
	levelHitCount := make([]int, 36)

	carGodService := NewCarGodService()
	betAmount := int64(100)
	carGodService.Car.BetTable[rand.Intn(12)] = betAmount

	for i := range unitTimes {
		// 进度打印
		if i%(unitTimes/20) == 0 && processID == 1 {
			fmt.Printf("Testing: %d%%\n", i*100/unitTimes)
		}

		result := carGodService.MainGame()

		// 统计投注
		for k, v := range carGodService.Car.BetTable {
			if v != 0 {
				countBetList[k] += v
			}
		}

		// 统计结果
		hitCar := result.HitCar
		hitEventID := result.EventID
		levelIndex := result.LevelIndex

		countMainWinList[hitCar] += result.MainGameWin
		countEventWinList[hitCar] += result.EventWin
		countJackpotWin += result.JackpotWin

		// 按level统计
		levelHitCount[levelIndex]++
		levelBetList[levelIndex] += betAmount
		levelMainWinList[levelIndex] += result.MainGameWin
		levelEventWinList[levelIndex] += result.EventWin
		levelJackpotWin[levelIndex] += result.JackpotWin

		if hitEventID >= 0 {
			eventTriggerList[hitEventID]++
		}
	}

	// 通过channel发送结果
	countData := CountData{
		CountBetList:      countBetList,
		CountMainWinList:  countMainWinList,
		CountEventWinList: countEventWinList,
		CountJackpotWin:   countJackpotWin,
		EventTriggerList:  eventTriggerList,
		LevelBetList:      levelBetList,
		LevelMainWinList:  levelMainWinList,
		LevelEventWinList: levelEventWinList,
		LevelJackpotWin:   levelJackpotWin,
		LevelHitCount:     levelHitCount,
	}

	resultChan <- countData
}
