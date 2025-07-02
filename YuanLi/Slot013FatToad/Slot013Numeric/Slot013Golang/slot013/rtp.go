package slot013

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
)

const MAX_RESPIN_COUNT = 2 // 重轉次數統計到 2+
const MAX_WILD_COUNT = 12  // 免費遊戲 CumWildCount 統計到 12+
const MAX_STAGE = 3        // 免費遊戲最大階段

var FreeGameName = []string{"FREE_GAME_01", "FREE_GAME_02"}

type RTPVerifier struct {
	TotalCount  int // 總投注次數
	LineBet     int // 單次投注額
	slot        SlotProb
	TotalBet    float64    `json:"total_bet"`     // 單次總下注額
	Bet         float64    `json:"bet"`           // 总下注额
	Win         float64    `json:"win"`           // 总赢得金额
	BaseWin     float64    `json:"base_win"`      // 基础游戏赢得金额
	BaseHit     int        `json:"base_hit"`      // 基础游戏中奖次数
	BaseSymWin  BaseSymWin `json:"base_sym_win"`  // 各符号中奖统计
	FreeHit     int        `json:"free_hit"`      // 触发免费游戏次数
	FreeWinHit  int        `json:"free_win_hit"`  // 免费游戏中有中奖的次数
	FreeWin     float64    `json:"free_win"`      // 免费游戏总赢得金额
	FreeSpin    int        `json:"free_spin"`     // 免费游戏总旋转次数
	FeatureHit  int        `json:"feature_hit"`   // 天降横财触发次数
	MaxWin      float64    `json:"max_win"`       // 最大赢额倍数
	MaxWinTimes int        `json:"max_win_times"` // 达到限制最大赢额的次数
	Variance    []float64  `json:"variance"`      // 方差值

	// 不同赢额倍数区间的统计 [总体, 免费游戏, 基础野生增益]
	Win0        [3]int `json:"win_0"`         // 0倍赌注的次数
	Win01       [3]int `json:"win_0_1"`       // 0-1倍赌注的次数
	Win12       [3]int `json:"win_1_2"`       // 1-2倍赌注的次数
	Win23       [3]int `json:"win_2_3"`       // 2-3倍赌注的次数
	Win34       [3]int `json:"win_3_4"`       // 3-4倍赌注的次数
	Win45       [3]int `json:"win_4_5"`       // 4-5倍赌注的次数
	Win510      [3]int `json:"win_5_10"`      // 5-10倍赌注的次数
	Win1020     [3]int `json:"win_10_20"`     // 10-20倍赌注的次数
	Win2050     [3]int `json:"win_20_50"`     // 20-50倍赌注的次数
	Win50100    [3]int `json:"win_50_100"`    // 50-100倍赌注的次数
	Win100500   [3]int `json:"win_100_500"`   // 100-500倍赌注的次数
	Win5001000  [3]int `json:"win_500_1000"`  // 500-1000倍赌注的次数
	Win10005000 [3]int `json:"win_1000_5000"` // 1000-5000倍赌注的次数
	Win5000     [3]int `json:"win_5000"`      // 5000倍以上赌注的次数

	Feature010     int     `json:"feature_0_10"`     // 天降横财触发但赢额小于10倍赌注的次数
	Feature010Win  float64 `json:"feature_0_10_win"` // 天降横财触发但赢额小于10倍赌注的总赢额
	Feature1000Win float64 `json:"feature_1000_win"` // 天降横财触发且赢额大于等于500倍赌注的总赢额

	// 蟾蜍特色功能相关统计
	ToadSizeCount     [6]int  `json:"toad_size_count"`      // 蟾蜍大小统计(1-6)
	ToadMulCount      [50]int `json:"toad_mul_count"`       // 蟾蜍乘数统计(0-49)
	BaseWildBuffCount int     `json:"base_wild_buff_count"` // 基础游戏中野生增益触发次数
	WinBuffWin        float64 `json:"win_buff_win"`         // 赢额增益产生的总赢额
	NearMissCount     int     `json:"near_miss_count"`      // 主游戏金蟾出现次数
}

// BaseSymWin 各符号中奖统计结构体（按符号类型和连线数）
type BaseSymWin struct {
	H1 [6]int `json:"H1"` // H1高分符号的3连、4连、5连等中奖次数
	H2 [6]int `json:"H2"` // H2高分符号
	H3 [6]int `json:"H3"` // H3高分符号
	H4 [6]int `json:"H4"` // H4高分符号
	LA [6]int `json:"LA"` // A符号的中奖统计
	LK [6]int `json:"LK"` // K符号的中奖统计
	LQ [6]int `json:"LQ"` // Q符号的中奖统计
	LJ [6]int `json:"LJ"` // J符号的中奖统计
}

// Init 初始化
func (rv *RTPVerifier) Init() {
	rv.LineBet = 1
	rv.slot.Init()
}

// Clear 清除資料
func (rv *RTPVerifier) Clear() {

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
		rv.ProcessResult(result)
	}
}

// ProcessResult 處理投注結果
func (rv *RTPVerifier) ProcessResult(result *SlotResult) {
	rv.TotalBet = float64(result.TotalBet)
	rv.Variance = append(rv.Variance, float64(result.TotalWin)/rv.TotalBet)
	rv.Bet += float64(result.TotalBet)
	rv.Win += float64(result.TotalWin)
	rv.winCount(result.TotalWin, result.TotalBet, 0)
	if result.TotalWin > uint64(rv.MaxWin) {
		rv.MaxWin = float64(result.TotalWin)
	}
	if result.MainWin > 0 {
		rv.BaseWin += float64(result.MainWin)
		rv.BaseHit++
	}

	if result.FGSpinCount > 0 {
		rv.FreeHit++
		rv.FreeWinHit += result.FreeWinCount
		rv.FreeSpin += result.FGSpinCount
		rv.winCount(result.TotalWin, result.TotalBet, 1)
		if result.FreeWin > 0 {
			rv.FreeWin += float64(result.FreeWin)

			if float64(result.FreeWin)/rv.TotalBet < 10 {
				rv.Feature010++
				rv.Feature010Win += float64(result.FreeWin)
			} else if float64(result.FreeWin)/rv.TotalBet >= 500 {
				rv.Feature1000Win += float64(result.FreeWin)
			}
		}
	}

	if result.MGFeatureCount > 0 {
		rv.FeatureHit++
		rv.WinBuffWin += float64(result.MGFeatureWin)
		rv.BaseWildBuffCount++
		rv.winCount(result.TotalWin, result.TotalBet, 2)
	}

	if len(result.WWWild) > 0 {
		rv.NearMissCount++
		if result.WWLevel > 0 {
			rv.ToadSizeCount[result.WWLevel-1]++
		}

		rv.ToadMulCount[result.WWMultiplier]++
	}

	if result.MaxWin {
		rv.MaxWinTimes++
	}
}

// Dump 傾印統計結果
func (rv *RTPVerifier) Dump(detail bool) {
	// 打印总体RTP
	fmt.Printf("Total RTP：%.6f\n", float64(rv.Win)/float64(rv.Bet))
	fmt.Println("=============================")
	fmt.Printf("Base RTP：%.6f\n", float64(rv.BaseWin)/float64(rv.Bet))
	fmt.Printf("Base Hit Rate：%.6f\n", float64(rv.BaseHit)/float64(rv.TotalCount))

	// 处理可能的除零情况
	if rv.NearMissCount > 0 {
		fmt.Printf("Near Miss 间隔: %.6f\n", float64(rv.TotalCount)/float64(rv.NearMissCount))
	}

	// 天降横财相关统计
	if rv.FeatureHit > 0 {
		fmt.Printf("天降横财 RTP: %.6f\n", float64(rv.WinBuffWin)/float64(rv.Bet))
		fmt.Printf("天降横财的平均倍数: %.6f\n", float64(rv.WinBuffWin)/float64(rv.FeatureHit)/float64(rv.TotalBet))
		fmt.Printf("天降横财的间隔: %.6f\n", float64(rv.TotalCount)/float64(rv.FeatureHit))
		fmt.Printf("Base 天降概率: %.6f\n", float64(rv.BaseWildBuffCount)/float64(rv.TotalCount))
	}

	fmt.Println("=============================")

	// 免费游戏相关统计
	if rv.FreeHit > 0 {
		fmt.Printf("Free RTP：%.6f\n", float64(rv.FreeWin)/float64(rv.Bet))
		fmt.Printf("Free触发间隔：%.6f\n", float64(rv.TotalCount)/float64(rv.FreeHit))
		fmt.Printf("Free平均倍数：%.6f\n", float64(rv.FreeWin)/float64(rv.FreeHit)/float64(rv.TotalBet))

		if rv.FreeSpin > 0 {
			fmt.Printf("Free hit rate：%.6f\n", float64(rv.FreeWinHit)/float64(rv.FreeSpin))
		}

		fmt.Printf("Free平均次数：%.6f\n", float64(rv.FreeSpin)/float64(rv.FreeHit))
		fmt.Printf("Free win < 10 total bet: 占比/ RTP %.6f %.6f\n",
			float64(rv.Feature010)/float64(rv.FreeHit),
			float64(rv.Feature010Win)/float64(rv.Bet))
		fmt.Printf("Free win >= 500 total bet: RTP %.6f\n",
			float64(rv.Feature1000Win)/float64(rv.Bet))
	}

	fmt.Println("=============================")

	// 最大赢额统计
	fmt.Printf("Max_Win: %v\n", rv.MaxWin/rv.TotalBet)
	fmt.Printf("Max_Win Time / Test Times: %d / %d\n", rv.MaxWinTimes, rv.TotalCount)

	// 计算方差
	fmt.Printf("方差 Variance: %.4f\n", sampleStdDev(rv.Variance))

	// 使用tabwriter打印赢额分布表格
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Win Mul\ttotal\tfree\tbase天降")

	// 打印赢额分布
	rv.printWinDistribution(w, "win_0", rv.Win0)
	rv.printWinDistribution(w, "win_0_1", rv.Win01)
	rv.printWinDistribution(w, "win_1_2", rv.Win12)
	rv.printWinDistribution(w, "win_2_3", rv.Win23)
	rv.printWinDistribution(w, "win_3_4", rv.Win34)
	rv.printWinDistribution(w, "win_4_5", rv.Win45)
	rv.printWinDistribution(w, "win_5_10", rv.Win510)
	rv.printWinDistribution(w, "win_10_20", rv.Win1020)
	rv.printWinDistribution(w, "win_20_50", rv.Win2050)
	rv.printWinDistribution(w, "win_50_100", rv.Win50100)
	rv.printWinDistribution(w, "win_100_500", rv.Win100500)
	rv.printWinDistribution(w, "win_500_1000", rv.Win5001000)
	rv.printWinDistribution(w, "win_1000_5000", rv.Win10005000)
	rv.printWinDistribution(w, "win_5000", rv.Win5000)

	w.Flush()

	fmt.Println("=============================")

	// 打印蟾蜍大小统计
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "toad size\tpercent")

	toadSizeSum := 0
	for _, size := range rv.ToadSizeCount {
		toadSizeSum += size
	}

	for i := 0; i < 6; i++ {
		percent := 0.0
		if toadSizeSum > 0 {
			percent = float64(rv.ToadSizeCount[i]) / float64(toadSizeSum)
		}
		fmt.Fprintf(w, "%d\t%.5f\n", i+1, percent)
	}
	w.Flush()

	fmt.Println("=============================")

	// 打印蟾蜍乘数统计
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "toad Mul\tpercent")

	toadMulSum := 0
	for _, mul := range rv.ToadMulCount {
		toadMulSum += mul
	}

	for k := 0; k < len(rv.ToadMulCount); k++ {
		if rv.ToadMulCount[k] != 0 {
			percent := float64(rv.ToadMulCount[k]) / float64(toadMulSum)
			fmt.Fprintf(w, "%d\t%.5f\n", k, percent)
		}
	}
	w.Flush()
}

// 辅助函数：打印赢额分布行
func (rv *RTPVerifier) printWinDistribution(w *tabwriter.Writer, key string, values [3]int) {
	total := 0.0
	free := 0.0
	baseWild := 0.0

	if rv.TotalCount > 0 {
		total = float64(values[0]) / float64(rv.TotalCount)
	}

	if rv.FreeHit > 0 {
		free = float64(values[1]) / float64(rv.FreeHit)
	}

	if rv.BaseWildBuffCount > 0 {
		baseWild = float64(values[2]) / float64(rv.BaseWildBuffCount)
	}

	fmt.Fprintf(w, "%s\t%.4f\t%.4f\t%.4f\n", key, total, free, baseWild)
}

func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// 计算标准差
func sampleStdDev(data []float64) float64 {
	if len(data) < 2 {
		return 0 // 需要至少两个数据点
	}

	m := mean(data)
	sumSquares := 0.0
	for _, value := range data {
		diff := value - m
		sumSquares += diff * diff
	}

	return math.Sqrt(sumSquares / float64(len(data)-1))
}

// 统计赢额倍数区间
func (rv *RTPVerifier) winCount(spinTotalWin, totalBet uint64, idx int) {
	winMultiple := float64(spinTotalWin) / float64(totalBet)

	if winMultiple == 0 {
		rv.Win0[idx]++
	} else if 0 < winMultiple && winMultiple <= 1 {
		rv.Win01[idx]++
	} else if 1 < winMultiple && winMultiple <= 2 {
		rv.Win12[idx]++
	} else if 2 < winMultiple && winMultiple <= 3 {
		rv.Win23[idx]++
	} else if 3 < winMultiple && winMultiple <= 4 {
		rv.Win34[idx]++
	} else if 4 < winMultiple && winMultiple <= 5 {
		rv.Win45[idx]++
	} else if 5 < winMultiple && winMultiple <= 10 {
		rv.Win510[idx]++
	} else if 10 < winMultiple && winMultiple <= 20 {
		rv.Win1020[idx]++
	} else if 20 < winMultiple && winMultiple <= 50 {
		rv.Win2050[idx]++
	} else if 50 < winMultiple && winMultiple <= 100 {
		rv.Win50100[idx]++
	} else if 100 < winMultiple && winMultiple <= 500 {
		rv.Win100500[idx]++
	} else if 500 < winMultiple && winMultiple <= 1000 {
		rv.Win5001000[idx]++
	} else if 1000 < winMultiple && winMultiple <= 5000 {
		rv.Win10005000[idx]++
	} else if 5000 < winMultiple {
		rv.Win5000[idx]++
	}
}
