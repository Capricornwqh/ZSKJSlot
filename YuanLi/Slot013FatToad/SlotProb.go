package Slot013FatToad

import (
	"Force/GameServer/Common"
	"Force/GameServer/Utils"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"slices"
	"sort"
)

type WildStruct struct {
	WildSymbol     Symbol // Wild 獎圖
	WildCoordinate [2]int // Wild 獎圖座標
}

type SlotProb struct {
	MGReelGroupRouletteMaps    [RTP_TOTAL]map[int]Utils.Roulette[int] // 主遊戲轉輪群組 Roulette Map
	FGReelGroupRouletteMaps    [RTP_TOTAL]map[int]Utils.Roulette[int] // 免費遊戲轉輪群組 Roulette Map
	MGCoverScatterRouletteMaps [RTP_TOTAL]map[int]Utils.Roulette[int] // 主遊戲金蟾铜钱 Roulette Map
}

// Init 初始化
func (slot *SlotProb) Init() {
	// 建立 Roulette 資料
	slot.SetRTPRoulette()
}

// SetRTPRoulette 建立 Roulette 資料
func (slot *SlotProb) SetRTPRoulette() {
	for rtp := range RTP_TOTAL {
		// 主遊戲轉輪群組 Roulette Map
		rouletteMap := make(map[int]Utils.Roulette[int])
		for buyType, groupWT := range MGReelGroupWT[rtp] {
			roulette := Utils.NewRouletteFromList(groupWT)
			rouletteMap[buyType] = *roulette
		}
		slot.MGReelGroupRouletteMaps[rtp] = rouletteMap
	}
}

// MainGame
func (slot *SlotProb) Run(rtp int, lineBet int, result *SlotResult, debugCmdList []DebugCmd) int {
	result.GameMode = Common.GAME_MODE_NORMAL
	buyType := result.BuyType
	result.TotalBet = uint64(float64(lineBet*PAYLINE_TOTAL) * BetRatio[buyType])
	maxWin := uint64(lineBet * PAYLINE_TOTAL * MAX_ODDS)
	result.WWMultiplier = 0
	tumbleResult := &TumbleResult{}
	result.MGGroupIndex, result.Code, result.WWWild, result.SSWild = slot.RandMGSymbol(rtp, buyType, lineBet, tumbleResult, nil)
	// 有發生錯誤則直接結束
	if result.Code != Common.ERROR_CODE_OK {
		return result.GameMode
	}

	// 計算該盤面贏分
	tumbleResult.Win, tumbleResult.LineSymbol, tumbleResult.LineCount, tumbleResult.LineWin =
		slot.CalculateMainWin(lineBet, tumbleResult.TumbleSymbol)

	// 紀錄結果
	result.MainWin += tumbleResult.Win
	result.MGTumbleList = append(result.MGTumbleList, tumbleResult)

	// 如果赢分=0 且没有金蟾，触发天降横财
	if buyType == Common.BUY_NONE && result.MainWin <= 0 && len(result.WWWild) <= 0 &&
		rand.IntN(100) < FEATURE_WILD_MAIN_GAME {
		result.MGFeatureCount++
		tumbleResult1 := &TumbleResult{}
		tumbleSymbol := make(GameSymbol, len(tumbleResult.TumbleSymbol))
		for i, v := range tumbleResult.TumbleSymbol {
			tumbleSymbol[i] = slices.Clone(v)
		}

		count := Utils.RandChoiceByWeight(MGFeatureSSList, MGFeatureSSWT)

		// 随机选择一条中奖线
		lineIndex := rand.IntN(len(LineIndexArray))
		selectedLine := LineIndexArray[lineIndex]

		// 收集可放置铜钱的位置（按照中奖线的行位置，排除第1列和有百搭的列）
		availablePositions := make([][2]int, 0)
		for col := 1; col < SLOT_COL; col++ {
			// 检查这一列是否有百搭
			hasWild := false
			for r := range SLOT_ROW {
				if isWild(tumbleSymbol[col][r]) {
					hasWild = true
					break
				}
			}

			// 如果这一列没有百搭，则可以放置铜钱
			if !hasWild {
				row := selectedLine[col]
				availablePositions = append(availablePositions, [2]int{col, row})
			}
		}

		// 选择不超过count个位置来放置铜钱
		coinCount := min(int(count), len(availablePositions))

		// 在选定的位置上放置铜钱
		for i := range coinCount {
			position := availablePositions[i]
			col, row := position[0], position[1]

			// 放置铜钱
			tumbleSymbol[col][row] = slot.getMainWildMulti(result.BuyType, result.MGGroupIndex)
			result.SSWild = append(result.SSWild, &WildStruct{
				WildSymbol:     tumbleSymbol[col][row],
				WildCoordinate: [2]int{col, row},
			})
		}

		// 計算該盤面贏分
		tumbleResult1.Win, tumbleResult1.LineSymbol, tumbleResult1.LineCount, tumbleResult1.LineWin =
			slot.CalculateMainWin(lineBet, tumbleSymbol)
		tumbleResult1.TumbleSymbol = tumbleSymbol
		tumbleResult1.MGPerformanceType = PERFORMANCE_FEATURE
		result.MGTumbleList = append(result.MGTumbleList, tumbleResult1)
		result.MainWin += tumbleResult1.Win
		result.MGFeatureWin += tumbleResult1.Win
	}

	// 判斷是否進免費遊戲 (需未達最大贏分)
	if len(result.WWWild) > 0 && len(result.SSWild) > 0 {
		result.SSCount = 0
		result.WWLevel = 1
		result.FGSpinCount = 5
		slot.WWLevelUp(result)
	} else {
		result.WWWild = result.WWWild[:0]
		result.SSWild = result.SSWild[:0]
	}

	// 達到最大贏分則結束
	if result.MainWin >= maxWin {
		return result.GameMode
	}

	// 儲存結果
	result.TotalWin = result.MainWin
	result.FreeWin = 0
	tmpSpinCount := 0
	switch result.BuyType {
	case Common.BUY_NONE:
		result.FGIndex = Utils.RandChoiceByWeight(FGIndexList, FGIndexWT)
	case Common.BUY_FREE_SPINS:
		result.FGIndex = Utils.RandChoiceByWeight(FGIndexList, FGBuyFreeIndexWT)
	case Common.BUY_SUPER_FREE_SPINS:
		result.FGIndex = Utils.RandChoiceByWeight(FGIndexList, FGBuySuperIndexWT)
	default:
	}
	tmpSpeicalSpin := rand.IntN(5) + 1 // 特殊 Spin

	// 免費遊戲相關處理(需未達最大贏分)
	for result.FGSpinCount > 0 {
		result.GameMode = Common.GAME_MODE_FREE
		result.FGSpinCount--
		tmpSpinCount++
		if tmpSpinCount == tmpSpeicalSpin {
			result.SpecialSpin = true
		} else {
			result.SpecialSpin = false
		}
		tmpFreeWin := slot.RunFreeGame(rtp, lineBet, tmpSpinCount, result, debugCmdList)
		if result.TotalWin+tmpFreeWin >= maxWin {
			tmpFreeWin = maxWin - result.TotalWin
		}
		result.FreeWin += tmpFreeWin
		result.TotalWin += tmpFreeWin
		if result.TotalWin >= maxWin {
			// fmt.Printf("TotalWin %d >= maxWin %d, break\n", result.TotalWin, maxWin)
			break
		}

		if result.FGSpinCount == 0 {
			// 如果升级铜钱小于等于2，概率触发天降横财，强制升级
			requiredSSCount := 0
			// 计算升到下一级所需的铜钱数量
			if result.SSCount < WWLEVEL_SSWILD_1 {
				requiredSSCount = WWLEVEL_SSWILD_1 - result.SSCount
			} else if result.SSCount < WWLEVEL_SSWILD_2 {
				requiredSSCount = WWLEVEL_SSWILD_2 - result.SSCount
			} else if result.SSCount < WWLEVEL_SSWILD_3 {
				requiredSSCount = WWLEVEL_SSWILD_3 - result.SSCount
			} else if result.SSCount < WWLEVEL_SSWILD_4 {
				requiredSSCount = WWLEVEL_SSWILD_4 - result.SSCount
			} else if result.SSCount < WWLEVEL_SSWILD_5 {
				requiredSSCount = WWLEVEL_SSWILD_5 - result.SSCount
			}

			// 天降横财
			tmpValue := rand.IntN(100)
			endValue := -1
			switch result.BuyType {
			case Common.BUY_NONE:
				switch result.FGIndex {
				case FREE_INDEX_1:
					if result.WWLevel == 1 {
						endValue = FEATURE_BUY_NONE_LEVEL1_1
					} else if result.WWLevel == 2 {
						endValue = FEATURE_BUY_NONE_LEVEL2_1
					}
				case FREE_INDEX_2:
					if result.WWLevel == 1 {
						endValue = FEATURE_BUY_NONE_LEVEL1_2
					} else if result.WWLevel == 2 {
						endValue = FEATURE_BUY_NONE_LEVEL2_2
					}
				default:
				}
			case Common.BUY_FREE_SPINS:
				switch result.FGIndex {
				case FREE_INDEX_1:
					if result.WWLevel == 1 {
						endValue = FEATURE_BUY_FREE_LEVEL1_1
					} else if result.WWLevel == 2 {
						endValue = FEATURE_BUY_FREE_LEVEL2_1
					} else if result.WWLevel == 3 {
						endValue = FEATURE_BUY_FREE_LEVEL3_1
					}
				case FREE_INDEX_2:
					if result.WWLevel == 1 {
						endValue = FEATURE_BUY_FREE_LEVEL1_2
					} else if result.WWLevel == 2 {
						endValue = FEATURE_BUY_FREE_LEVEL2_2
					}
				default:
				}
			case Common.BUY_SUPER_FREE_SPINS:
				switch result.FGIndex {
				case FREE_INDEX_1:
					if result.WWLevel == 2 {
						endValue = FEATURE_BUY_SUPER_LEVEL2_1
					} else if result.WWLevel == 3 {
						endValue = FEATURE_BUY_SUPER_LEVEL3_1
					} else if result.WWLevel == 4 {
						endValue = FEATURE_BUY_SUPER_LEVEL4_1
					}
				case FREE_INDEX_2:
					if result.WWLevel == 2 {
						endValue = FEATURE_BUY_SUPER_LEVEL2_2
					} else if result.WWLevel == 3 {
						endValue = FEATURE_BUY_SUPER_LEVEL3_2
					} else if result.WWLevel == 4 {
						endValue = FEATURE_BUY_SUPER_LEVEL4_2
					}
				default:
				}
			default:
			}

			if requiredSSCount <= 2 && endValue > 0 && tmpValue < endValue {
				result.FGFeatureCount++
				tumbleSymbol := make(GameSymbol, len(result.FGSpinList[len(result.FGSpinList)-1].TumbleSymbol))
				for i, v := range result.FGSpinList[len(result.FGSpinList)-1].TumbleSymbol {
					tumbleSymbol[i] = slices.Clone(v)
				}

				// 记录每列可用的位置
				availableCols := make([]int, 0, SLOT_COL)
				availablePositions := make(map[int][][2]int)

				for col := range SLOT_COL {
					hasAvailable := false
					for row := range SLOT_ROW {
						if !isWild(tumbleSymbol[col][row]) {
							hasAvailable = true
							break
						}
					}

					if hasAvailable {
						availableCols = append(availableCols, col)
						availablePositions[col] = make([][2]int, 0)

						for row := range SLOT_ROW {
							if !isWild(tumbleSymbol[col][row]) {
								availablePositions[col] = append(availablePositions[col], [2]int{col, row})
							}
						}
					}
				}

				// 随机打乱可用列的顺序
				rand.Shuffle(len(availableCols), func(i, j int) {
					availableCols[i], availableCols[j] = availableCols[j], availableCols[i]
				})

				// 选择列并放置铜钱
				coinCount := min(requiredSSCount, len(availableCols))

				for i := range coinCount {
					col := availableCols[i]

					if len(availablePositions[col]) > 0 {
						// 随机选择一个非金蟾位置
						randomIndex := rand.IntN(len(availablePositions[col]))
						position := availablePositions[col][randomIndex]

						// 放置铜钱
						tumbleSymbol[position[0]][position[1]] = slot.getFreeWildMulti(result.BuyType, result.FGIndex)
						result.SSWild = append(result.SSWild, &WildStruct{
							WildSymbol:     tumbleSymbol[position[0]][position[1]],
							WildCoordinate: position,
						})
					}
				}

				result.FGSpinList = append(result.FGSpinList, &SpinResult{
					TumbleSymbol:    tumbleSymbol,
					PerformanceType: PERFORMANCE_FEATURE,
					SSCount:         result.SSCount,
				})

				slot.WWLevelUp(result)
			}
		}
	}

	result.FGSpinCount = tmpSpinCount
	return result.GameMode
}

// 获取主游戏铜钱倍数
func (slot *SlotProb) getMainWildMulti(buyType, index int) Symbol {
	switch buyType {
	case Common.BUY_NONE:
		if index == MAIN_GAME_01 && rand.IntN(100) < SSWILD_MAIN_GAME_1 {
			return Utils.RandChoiceByWeight(MGSSMultiList, MGSSMultiWT)
		} else if index == MAIN_GAME_02 && rand.IntN(100) < SSWILD_MAIN_GAME_2 {
			return Utils.RandChoiceByWeight(MGSSMultiList, MGSSMultiWT)
		}
	case Common.BUY_FREE_SPINS:
		if rand.IntN(100) < SSWILD_BUY_FREE {
			return Utils.RandChoiceByWeight(MGSSMultiList, MGBuySSMultiWT)
		}
	case Common.BUY_SUPER_FREE_SPINS:
		if rand.IntN(100) < SSWILD_BUY_SUPER {
			return Utils.RandChoiceByWeight(MGSSMultiList, MGBuySSMultiWT)
		}
	default:
	}

	return SS
}

// 获取免费游戏铜钱倍数
func (slot *SlotProb) getFreeWildMulti(buyType, index int) Symbol {
	switch buyType {
	case Common.BUY_NONE:
		if index == FREE_INDEX_2 && rand.IntN(100) < SSWILD_FREE_GAME_2 {
			return Utils.RandChoiceByWeight(FGSSMultiList, FGSSMultiWT)
		}
	case Common.BUY_FREE_SPINS:
		if index == FREE_INDEX_2 && rand.IntN(100) < SSWILD_BUY_FREE_2 {
			return Utils.RandChoiceByWeight(FGSSMultiList, FGSSMultiWT)
		}
	case Common.BUY_SUPER_FREE_SPINS:
		if index == FREE_INDEX_2 && rand.IntN(100) < SSWILD_BUY_SUPER_2 {
			return Utils.RandChoiceByWeight(FGSSMultiList, FGSSMultiWT)
		}
	default:
	}

	return SS
}

// getWWSize 根据金蟾等级获取其大小
func (slot *SlotProb) getWWSize(level int) int {
	switch level {
	case 2:
		return 2 // 2×2大小
	case 3:
		return 3 // 3×3大小
	case 4:
		return 4 // 4×4大小
	case 5:
		return 5 // 5×5大小
	case 6:
		return 6 // 6×6大小
	default:
		return 1
	}
}

// 辅助函数：求绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 计算金蟾到铜钱的最小距离，同时返回金蟾上与铜钱距离最近的点
func minDistanceAndPoint(toad []*WildStruct, coin *WildStruct) (int, *WildStruct) {
	minDist := math.MaxInt32
	closestPoint := &WildStruct{WildSymbol: 0, WildCoordinate: [2]int{-1, -1}} // 初始化为无效点
	for _, p := range toad {
		d := abs(p.WildCoordinate[0]-coin.WildCoordinate[0]) + abs(p.WildCoordinate[1]-coin.WildCoordinate[1])
		if d < minDist {
			minDist = d
			closestPoint = p
		}
	}
	return minDist, closestPoint
}

// 获取点相对于参考点的方位区域
// 返回值: 0=右边, 1=左边, 2=上边, 3=下边
func getDirectionAreaFromPoint(point *WildStruct, refPoint *WildStruct) int {
	if point.WildCoordinate[0] > refPoint.WildCoordinate[0] {
		return 0 // 右
	}
	if point.WildCoordinate[0] < refPoint.WildCoordinate[0] {
		return 1 // 左
	}
	if point.WildCoordinate[1] < refPoint.WildCoordinate[1] {
		return 2 // 上
	}
	if point.WildCoordinate[1] > refPoint.WildCoordinate[1] {
		return 3 // 下
	}

	// 点与参考点重叠
	return -1
}

// 计算点在方位区域内的优先级（值越小优先级越高）
func getPositionPriority(refPoint *WildStruct, point *WildStruct) int {
	// 计算到参考点的曼哈顿距离
	distToRef := abs(point.WildCoordinate[0]-refPoint.WildCoordinate[0]) + abs(point.WildCoordinate[1]-refPoint.WildCoordinate[1])
	return distToRef
}

// 找到距离金蟾最近的铜钱
func nextTarget(toad []*WildStruct, coins []*WildStruct) *WildStruct {
	type candidate struct {
		point            *WildStruct
		distance         int         // 铜钱到金蟾的最小距离
		closestToadPoint *WildStruct // 金蟾上最近的点
		directionArea    int         // 方位区域: 0=右, 1=左, 2=上, 3=下
		positionPriority int         // 在方位内的优先级
	}

	var cands []candidate
	minDist := math.MaxInt32

	// 1. 找出距离最近的铜钱
	for _, coin := range coins {
		d, closestPoint := minDistanceAndPoint(toad, coin)
		if d < minDist {
			minDist = d
			area := getDirectionAreaFromPoint(coin, closestPoint)
			priority := getPositionPriority(closestPoint, coin)
			cands = []candidate{{coin, d, closestPoint, area, priority}}
		} else if d == minDist {
			area := getDirectionAreaFromPoint(coin, closestPoint)
			priority := getPositionPriority(closestPoint, coin)
			cands = append(cands, candidate{coin, d, closestPoint, area, priority})
		}
	}

	// 2. 按方位区域优先级排序: 右(0) > 左(1) > 上(2) > 下(3)
	// 3. 在同一方位内，按照位置优先级排序
	// 4. 如果优先级相同，先取X值大的点，再取Y值小的点
	sort.Slice(cands, func(i, j int) bool {
		if cands[i].directionArea != cands[j].directionArea {
			return cands[i].directionArea < cands[j].directionArea
		}
		if cands[i].positionPriority != cands[j].positionPriority {
			return cands[i].positionPriority < cands[j].positionPriority
		}
		if cands[i].point.WildCoordinate[0] != cands[j].point.WildCoordinate[0] {
			return cands[i].point.WildCoordinate[0] > cands[j].point.WildCoordinate[0]
		}
		return cands[i].point.WildCoordinate[1] < cands[j].point.WildCoordinate[1]
	})

	return cands[0].point
}

// 获取金蟾的尺寸和左上角坐标
func getToadSizeAndTopLeft(toad []*WildStruct) (int, *WildStruct) {
	if len(toad) == 0 {
		return 0, &WildStruct{WildSymbol: 0, WildCoordinate: [2]int{0, 0}}
	}

	// 找出最小的x和y坐标作为左上角
	minX, minY := toad[0].WildCoordinate[0], toad[0].WildCoordinate[1]
	maxX, maxY := toad[0].WildCoordinate[0], toad[0].WildCoordinate[1]

	for _, p := range toad {
		if p.WildCoordinate[0] < minX {
			minX = p.WildCoordinate[0]
		}
		if p.WildCoordinate[1] < minY {
			minY = p.WildCoordinate[1]
		}
		if p.WildCoordinate[0] > maxX {
			maxX = p.WildCoordinate[0]
		}
		if p.WildCoordinate[1] > maxY {
			maxY = p.WildCoordinate[1]
		}
	}

	// 计算尺寸
	size := maxX - minX + 1
	// 确保是正方形
	if maxY-minY+1 > size {
		size = maxY - minY + 1
	}

	return size, &WildStruct{WildSymbol: 0, WildCoordinate: [2]int{minX, minY}}
}

// 移动金蟾到目标铜钱位置
func moveTo(toad []*WildStruct, target *WildStruct) []*WildStruct {
	size, topLeft := getToadSizeAndTopLeft(toad)
	minDist := math.MaxInt32
	var newTopLeft WildStruct

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			tlx, tly := target.WildCoordinate[0]-i, target.WildCoordinate[1]-j
			if tlx < 0 || tly < 0 || tlx+size > 6 || tly+size > 6 {
				continue
			}
			d := abs(topLeft.WildCoordinate[0]-tlx) + abs(topLeft.WildCoordinate[1]-tly)
			if d < minDist {
				minDist = d
				newTopLeft = WildStruct{WildSymbol: 0, WildCoordinate: [2]int{tlx, tly}}
			}
		}
	}

	// 计算金蟾需要移动的距离
	dx := newTopLeft.WildCoordinate[0] - topLeft.WildCoordinate[0]
	dy := newTopLeft.WildCoordinate[1] - topLeft.WildCoordinate[1]

	// 更新金蟾所有点的位置
	newToad := make([]*WildStruct, len(toad))
	for i, p := range toad {
		newToad[i] = &WildStruct{
			WildSymbol:     p.WildSymbol,
			WildCoordinate: [2]int{p.WildCoordinate[0] + dx, p.WildCoordinate[1] + dy},
		}
	}
	return newToad
}

// 是否有铜钱被金蟾覆盖
func isCovered(toad []*WildStruct, coin *WildStruct) bool {
	for _, p := range toad {
		if p.WildCoordinate[0] == coin.WildCoordinate[0] && p.WildCoordinate[1] == coin.WildCoordinate[1] {
			return true
		}
	}
	return false
}

// 金蟾变大算法
func (slot *SlotProb) expandFrog(wwStuct []*WildStruct, level int) []*WildStruct {
	// 如果金蟾坐标列表为空，直接返回
	if len(wwStuct) == 0 {
		return nil
	}

	// 找出金蟾的四个边界顶点
	minX, minY := wwStuct[0].WildCoordinate[0], wwStuct[0].WildCoordinate[1]
	maxX, maxY := minX, minY

	// 遍历所有金蟾坐标，找出最小和最大的X、Y值
	for _, wild := range wwStuct {
		x, y := wild.WildCoordinate[0], wild.WildCoordinate[1]
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	// 获取当前金蟾大小和新的金蟾大小
	currentWidth := maxX - minX + 1
	currentHeight := maxY - minY + 1
	newSize := slot.getWWSize(level)

	// 计算需要扩展的大小
	expandWidth := newSize - currentWidth
	expandHeight := newSize - currentHeight

	// 计算四个方向的可用空间
	rightSpace := 5 - maxX  // 右边可用空间
	leftSpace := minX       // 左边可用空间
	topSpace := minY        // 上边可用空间
	bottomSpace := 5 - maxY // 下边可用空间

	// 定义方向及其空间大小，用于排序
	type Direction struct {
		name     string
		space    int
		priority int // 优先级，数字越小优先级越高
	}

	directions := []Direction{
		{"right", rightSpace, 0},   // 优先级：右 0
		{"top", topSpace, 1},       // 优先级：上 1
		{"left", leftSpace, 2},     // 优先级：左 2
		{"bottom", bottomSpace, 3}, // 优先级：下 3
	}

	// 按可用空间从大到小排序，相同空间时按优先级排序
	slices.SortFunc(directions, func(a, b Direction) int {
		if a.space != b.space {
			return b.space - a.space // 从大到小排序
		}
		return a.priority - b.priority // 按优先级从高到低
	})

	// 计算新的左上角坐标
	newX := minX
	newY := minY

	// 处理水平扩展
	horizontalExpanded := 0
	for _, dir := range directions {
		if horizontalExpanded >= expandWidth {
			break
		}

		if dir.name == "right" {
			// 向右扩展
			expand := min(expandWidth-horizontalExpanded, dir.space)
			horizontalExpanded += expand
		} else if dir.name == "left" {
			// 向左扩展
			expand := min(expandWidth-horizontalExpanded, dir.space)
			if expand > 0 {
				newX = minX - expand
				horizontalExpanded += expand
			}
		}
	}

	// 处理垂直扩展
	verticalExpanded := 0
	for _, dir := range directions {
		if verticalExpanded >= expandHeight {
			break
		}

		if dir.name == "top" {
			// 向上扩展
			expand := min(expandHeight-verticalExpanded, dir.space)
			if expand > 0 {
				newY = minY - expand
				verticalExpanded += expand
			}
		} else if dir.name == "bottom" {
			// 向下扩展
			expand := min(expandHeight-verticalExpanded, dir.space)
			verticalExpanded += expand
		}
	}

	// 确保新的位置不会超出边界
	if newX < 0 {
		newX = 0
	}
	if newY < 0 {
		newY = 0
	}
	if newX+newSize > 6 {
		newX = 6 - newSize
	}
	if newY+newSize > 6 {
		newY = 6 - newSize
	}

	// 清空原有的金蟾坐标列表
	wwStuct = wwStuct[:0]

	// 生成新的金蟾坐标
	for i := range newSize {
		for j := range newSize {
			wwStuct = append(wwStuct, &WildStruct{
				WildSymbol:     WW,
				WildCoordinate: [2]int{newX + i, newY + j},
			})
		}
	}

	return wwStuct
}

// 金蟾升级
func (slot *SlotProb) WWLevelUp(result *SlotResult) {
	/*
	   吃夠一定數量的銅錢時，三足金蟾會變大，並增加Free Spins
	   ．Lv1 -> Lv2：累積 5 銅錢，+3 Free Spins，三足金蟾變成 2*2 大小
	   ．Lv2 -> Lv3：累積 4 銅錢，+3 Free Spins，三足金蟾變成 3*3 大小
	   ．Lv3 -> Lv4：累積 4 銅錢，+2 Free Spins，三足金蟾變成 4*4 大小
	   ．Lv4 -> Lv5：累積 3 銅錢，+2 Free Spins，三足金蟾變成 5*5 大小
	   ．Lv5 -> Lv6：累積 3 銅錢，+1 Free Spins，三足金蟾變成 6*6 大小
	*/

	// 如果没有铜钱或结果为空，则不处理
	if len(result.SSWild) <= 0 || result == nil || len(result.WWWild) <= 0 {
		return
	}

	// 获取当前盘面用于表演
	var tumbleSymbol GameSymbol
	if result.GameMode == Common.GAME_MODE_FREE {
		sourceSymbol := result.FGSpinList[len(result.FGSpinList)-1].TumbleSymbol
		tumbleSymbol = make(GameSymbol, len(sourceSymbol))
		for i, v := range sourceSymbol {
			tumbleSymbol[i] = slices.Clone(v)
		}
	} else {
		sourceSymbol := result.MGTumbleList[len(result.MGTumbleList)-1].TumbleSymbol
		tumbleSymbol = make(GameSymbol, len(sourceSymbol))
		for i, v := range sourceSymbol {
			tumbleSymbol[i] = slices.Clone(v)
		}
	}
	if len(tumbleSymbol) <= 0 {
		return
	}

	// 记录当前金蟾等级
	oldLevel := result.WWLevel

	// 逐个吃掉所有铜钱
	for len(result.SSWild) > 0 {
		// 从剩余铜钱中找出离金蟾最近的一个
		nextCoin := nextTarget(result.WWWild, result.SSWild)

		// 清除旧的金蟾位置
		for _, v := range result.WWWild {
			tumbleSymbol[v.WildCoordinate[0]][v.WildCoordinate[1]] = NN
		}

		// 移动金蟾到目标铜钱位置
		toadWilds := moveTo(result.WWWild, nextCoin)

		// 更新金蟾位置
		result.WWWild = result.WWWild[:0]
		for _, wild := range toadWilds {
			result.WWWild = append(result.WWWild, &WildStruct{
				WildSymbol:     wild.WildSymbol,
				WildCoordinate: wild.WildCoordinate,
			})
			tumbleSymbol[wild.WildCoordinate[0]][wild.WildCoordinate[1]] = WW
		}

		// 处理当前目标铜钱，增加计数
		switch nextCoin.WildSymbol {
		case SS:
			result.SSCount += 1
		case SS2:
			result.SSCount += 1
			result.WWMultiplier += 2
		case SS3:
			result.SSCount += 1
			result.WWMultiplier += 3
		case SS5:
			result.SSCount += 1
			result.WWMultiplier += 5
		}

		// 从铜钱列表中移除已吃掉的铜钱
		for i := 0; i < len(result.SSWild); i++ {
			if result.SSWild[i].WildCoordinate[0] == nextCoin.WildCoordinate[0] &&
				result.SSWild[i].WildCoordinate[1] == nextCoin.WildCoordinate[1] {
				result.SSWild = append(result.SSWild[:i], result.SSWild[i+1:]...)
				break
			}
		}

		// 添加表演效果
		tmpSpinResult := &SpinResult{
			TumbleSymbol:    make(GameSymbol, len(tumbleSymbol)),
			PerformanceType: PERFORMANCE_EAT,
			SSCount:         result.SSCount,
		}
		for i, v := range tumbleSymbol {
			tmpSpinResult.TumbleSymbol[i] = slices.Clone(v)
		}
		result.FGSpinList = append(result.FGSpinList, tmpSpinResult)
	}

	// 清空铜钱列表，因为所有铜钱都被吃掉了
	result.SSWild = result.SSWild[:0]

	// 吃完所有铜钱后，检查是否需要升级
	if result.SSCount >= WWLEVEL_SSWILD_5 {
		result.WWLevel = 6
	} else if result.SSCount >= WWLEVEL_SSWILD_4 {
		result.WWLevel = 5
	} else if result.SSCount >= WWLEVEL_SSWILD_3 {
		result.WWLevel = 4
	} else if result.SSCount >= WWLEVEL_SSWILD_2 {
		result.WWLevel = 3
	} else if result.SSCount >= WWLEVEL_SSWILD_1 {
		result.WWLevel = 2
	} else if result.SSCount >= 0 {
		result.WWLevel = 1
	}

	// 如果金蟾升级了，使用变大算法更新其大小和坐标
	if result.WWLevel > oldLevel && len(result.WWWild) > 0 {
		// 使用变大算法计算新的基准位置
		result.WWWild = slot.expandFrog(result.WWWild, result.WWLevel)

		// 在盘面上显示新的金蟾
		for _, v := range result.WWWild {
			tumbleSymbol[v.WildCoordinate[0]][v.WildCoordinate[1]] = WW
		}

		// 增加免费游戏次数
		result.FGSpinCount += FGInitSpinCount[result.WWLevel] - FGInitSpinCount[oldLevel]

		// 添加表演效果
		result.FGSpinList = append(result.FGSpinList, &SpinResult{
			TumbleSymbol:    tumbleSymbol,
			PerformanceType: PERFORMANCE_LEVELUP,
			SSCount:         result.SSCount,
		})
	}
}

// FreeGame
func (slot *SlotProb) RunFreeGame(rtp int, lineBet int, spinCount int, result *SlotResult, debugCmdList []DebugCmd) uint64 {
	buyType := result.BuyType
	spinResult := &SpinResult{
		SSCount: result.SSCount,
	}

	var debugCmd *DebugCmd
	if len(debugCmdList) > spinCount {
		debugCmd = &debugCmdList[spinCount]
	}
	_, result.Code = slot.RandFGSymbol(rtp, buyType, lineBet, spinResult, result, debugCmd)
	// 有發生錯誤則直接結束
	if result.Code != Common.ERROR_CODE_OK {
		return 0
	}

	// 計算該盤面贏分
	slot.CalculateFreeWin(lineBet, spinResult, result)

	// 紀錄結果
	result.FGSpinList = append(result.FGSpinList, spinResult)

	// 判斷是否進免費遊戲 (需未達最大贏分)
	if len(result.SSWild) > 0 {
		slot.WWLevelUp(result)
	}

	spinResult.Multiplier = result.WWMultiplier
	return spinResult.SpinWin
}

// RandMGSymbol 亂數產生主遊戲獎圖盤面
func (slot *SlotProb) RandMGSymbol(rtp int, buyType int, lineBet int, tumbleResult *TumbleResult, debugCmd *DebugCmd) (int, int, []*WildStruct, []*WildStruct) {
	roulette := slot.MGReelGroupRouletteMaps[rtp][buyType]
	groupIdx, ok := roulette.Spin()
	if !ok {
		fmt.Printf("[ERROR] RandMGSymbol: RTP = %d, BuyType = %d, MGReelGroupRoulette spin failed.\n", rtp, buyType)
		return -1, Common.ERROR_CODE_ROULETTE_SPIN_FAILED, nil, nil
	}
	reelGroup := MGReelGroup[rtp][groupIdx]

	tmpWWWild := make([]*WildStruct, 0)
	tmpSSWild := make([]*WildStruct, 0)

	// 產出盤面
	tumbleResult.TumbleSymbol = make(GameSymbol, SLOT_COL)
	tumbleResult.ReelIndex = make([]int, SLOT_COL)
	for col := range SLOT_COL {
		// 產出獎圖
		reel := reelGroup[col]
		reelLength := len(reel)
		dice := rand.IntN(reelLength)
		columnSymbol := make(ReelSymbol, SLOT_ROW)
		for row := range SLOT_ROW {
			idx := dice + row
			if idx < reelLength {
				columnSymbol[row] = reel[idx]
			} else {
				columnSymbol[row] = reel[idx-reelLength]
			}
		}
		tumbleResult.TumbleSymbol[col] = columnSymbol
		tumbleResult.ReelIndex[col] = dice
	}

	if buyType == Common.BUY_FREE_SPINS {
		// 随机打乱列2-6的顺序
		rand.Shuffle(SLOT_COL-1, func(i, j int) {
			col1 := i + 1
			col2 := j + 1
			tumbleResult.TumbleSymbol[col1], tumbleResult.TumbleSymbol[col2] = tumbleResult.TumbleSymbol[col2], tumbleResult.TumbleSymbol[col1]
			tumbleResult.ReelIndex[col1], tumbleResult.ReelIndex[col2] = tumbleResult.ReelIndex[col2], tumbleResult.ReelIndex[col1]
		})
	}

	for col := range SLOT_COL {
		for row := range SLOT_ROW {
			if tumbleResult.TumbleSymbol[col][row] == WW {
				tmpWWWild = append(tmpWWWild, &WildStruct{WildSymbol: WW, WildCoordinate: [2]int{col, row}})
			} else if tumbleResult.TumbleSymbol[col][row] == SS {
				tmpSSWild = append(tmpSSWild, &WildStruct{WildSymbol: SS, WildCoordinate: [2]int{col, row}})
			}
		}
	}

	// 处理铜钱倍数
	if len(tmpSSWild) > 0 {
		if buyType == Common.BUY_SUPER_FREE_SPINS {
			// Buy Super 的时候，从铜钱wild里面随机选一个变成倍数
			randIndex := rand.IntN(len(tmpSSWild))
			tmpSSWild[randIndex].WildSymbol = Utils.RandChoiceByWeight(MGSSMultiList, MGSSMultiWT)
			tumbleResult.TumbleSymbol[tmpSSWild[randIndex].WildCoordinate[0]][tmpSSWild[randIndex].WildCoordinate[1]] = tmpSSWild[randIndex].WildSymbol

			for i := range tmpSSWild {
				if i != randIndex {
					tmpSSWild[i].WildSymbol = slot.getMainWildMulti(buyType, groupIdx)
					tumbleResult.TumbleSymbol[tmpSSWild[i].WildCoordinate[0]][tmpSSWild[i].WildCoordinate[1]] = tmpSSWild[i].WildSymbol
				}
			}
		} else {
			for _, ssWild := range tmpSSWild {
				ssWild.WildSymbol = slot.getMainWildMulti(buyType, groupIdx)
				tumbleResult.TumbleSymbol[ssWild.WildCoordinate[0]][ssWild.WildCoordinate[1]] = ssWild.WildSymbol
			}
		}
	}

	return groupIdx, Common.ERROR_CODE_OK, tmpWWWild, tmpSSWild
}

// RandFGSymbol  亂數產生免費遊戲獎圖盤面
func (slot *SlotProb) RandFGSymbol(rtp int, buyType int, lineBet int, spinResult *SpinResult, result *SlotResult, debugCmd *DebugCmd) (int, int) {
	groupIdx := Utils.RandChoiceByWeight(FGGroupIndexList, FGLevelGroupIndexWT[result.BuyType][result.FGIndex][result.WWLevel])
	if result.SpecialSpin {
		groupIdx = FREE_GAME_04
	}
	// 處理測試指令
	var debugReelIndex []int
	if debugCmd != nil {
		// 取得轉輪群組 index，並檢查是否合法
		if len(debugCmd.DebugData) > DEBUG_INDEX_GROUP_INDEX && 0 <= debugCmd.DebugData[DEBUG_INDEX_GROUP_INDEX] && debugCmd.DebugData[DEBUG_INDEX_GROUP_INDEX] < len(MGReelGroup[rtp]) {
			groupIdx = debugCmd.DebugData[DEBUG_INDEX_GROUP_INDEX]
		}
		if len(debugCmd.DebugData) > DEBUG_INDEX_REEL_INDEX_06 {
			debugReelIndex = debugCmd.DebugData[DEBUG_INDEX_REEL_INDEX_01 : DEBUG_INDEX_REEL_INDEX_06+1]
		}
	}

	replaceSymbol := NN
	reelGroup := FGReelGroup[rtp][groupIdx]

	if groupIdx == FREE_GAME_04 {
		// 全盘替代
		replaceSymbol = Utils.RandChoiceByWeight(MysteryList, MysteryGloballWT)
	}

	// 產出盤面
	spinResult.TumbleSymbol = make(GameSymbol, SLOT_COL)
	spinResult.ReelIndex = make([]int, SLOT_COL)
	for col := range SLOT_COL {
		// 產出獎圖
		reel := reelGroup[col]
		reelLength := len(reel)
		dice := rand.IntN(reelLength)
		// 處理測試指令中的停輪位置，並檢查是否合法
		if debugReelIndex != nil && 0 <= debugReelIndex[col] && debugReelIndex[col] < reelLength {
			dice = debugReelIndex[col]
		}

		if groupIdx == FREE_GAME_01 || groupIdx == FREE_GAME_02 {
			// 按列替代
			replaceSymbol = Utils.RandChoiceByWeight(MysteryList, MysteryColWT)
		}
		columnSymbol := make(ReelSymbol, SLOT_ROW)
		for row := range SLOT_ROW {
			idx := dice + row
			if idx < reelLength {
				columnSymbol[row] = reel[idx]
			} else {
				columnSymbol[row] = reel[idx-reelLength]
			}
			// 如果是替换符号，则替换
			if replaceSymbol != NN && columnSymbol[row] == NN {
				columnSymbol[row] = replaceSymbol
			}
		}
		spinResult.TumbleSymbol[col] = columnSymbol
		spinResult.ReelIndex[col] = dice
	}
	// 金蟾
	for _, v := range result.WWWild {
		spinResult.TumbleSymbol[v.WildCoordinate[0]][v.WildCoordinate[1]] = WW
	}
	// 铜钱
	for col := range SLOT_COL {
		for row := range SLOT_ROW {
			if spinResult.TumbleSymbol[col][row] == SS {
				spinResult.TumbleSymbol[col][row] = slot.getFreeWildMulti(result.BuyType, result.FGIndex)
				result.SSWild = append(result.SSWild, &WildStruct{WildSymbol: spinResult.TumbleSymbol[col][row], WildCoordinate: [2]int{col, row}})
			}
		}
	}

	return groupIdx, Common.ERROR_CODE_OK
}

// CalculateMainWin 計算連線獎金
func (slot *SlotProb) CalculateMainWin(lineBet int, gameSymbol GameSymbol) (uint64, []Symbol, []int, []uint64) {
	tmpWin := uint64(0)
	tmpLineSymbol := make([]Symbol, len(LineIndexArray))
	tmpLineCount := make([]int, len(LineIndexArray))
	tmpLineWin := make([]uint64, len(LineIndexArray))

	for i, lineIndex := range LineIndexArray {
		lineSymbol := make([]Symbol, len(lineIndex))
		for col, row := range lineIndex {
			lineSymbol[col] = gameSymbol[col][row]
		}
		symbol, count, multiplier := slot.CheckLine(1, lineSymbol)
		if symbol >= WW && count > 0 {
			lineWin := uint64(lineBet * multiplier)
			tmpWin += lineWin
			tmpLineSymbol[i] = symbol
			tmpLineCount[i] = count
			tmpLineWin[i] = lineWin
		} else {
			tmpLineSymbol[i] = NN
		}
		// fmt.Printf("CalculateWin: %d) %v Win: %d\n", i, lineSymbol, lineWin)
	}

	return tmpWin, tmpLineSymbol, tmpLineCount, tmpLineWin
}

// CalculateFreeWin 計算連線獎金
func (slot *SlotProb) CalculateFreeWin(lineBet int, spinResult *SpinResult, result *SlotResult) {
	spinResult.SpinWin = 0
	spinResult.LineSymbol = make([]Symbol, len(LineIndexArray))
	spinResult.LineCount = make([]int, len(LineIndexArray))
	spinResult.LineWin = make([]uint64, len(LineIndexArray))

	for i, lineIndex := range LineIndexArray {
		lineSymbol := make([]Symbol, len(lineIndex))
		for col, row := range lineIndex {
			lineSymbol[col] = spinResult.TumbleSymbol[col][row]
		}
		wwMultiplier := 1
		if result.WWMultiplier > 1 {
			wwMultiplier = result.WWMultiplier
		}
		symbol, count, multiplier := slot.CheckLine(wwMultiplier, lineSymbol)
		if symbol >= WW && count > 0 {
			lineWin := uint64(lineBet * multiplier)
			spinResult.SpinWin += lineWin
			spinResult.LineSymbol[i] = symbol
			spinResult.LineCount[i] = count
			spinResult.LineWin[i] = lineWin
		} else {
			spinResult.LineSymbol[i] = NN
		}
		// fmt.Printf("CalculateWin: %d) %v Win: %d\n", i, lineSymbol, lineWin)
	}
}

// 判断符号是否是百搭
func isWild(s Symbol) bool {
	return s == WW || s == SS || s == SS2 || s == SS3 || s == SS5
}

// CheckLine 檢查連線及獎圖
func (slot *SlotProb) CheckLine(wwMultiplier int, lineSymbol []Symbol) (symbol Symbol, count int, multiplier int) {
	if len(lineSymbol) < SLOT_COL {
		fmt.Printf("[ERROR] CheckLine: LineSymbol length is invalid.\n")
		return NN, 0, 1
	}

	symbol, count = NN, 0
	findSymbol := lineSymbol[0]
	linkCount := 0
	multiplier = 0

	// 比對獎圖
	for i, s := range lineSymbol {
		// 如果当前符号不是百搭，且前一个符号不是百搭，且不匹配，则中断连线
		if !isWild(s) && !isWild(findSymbol) && s != findSymbol {
			break
		}

		// 如果找到非百搭符号，更新findSymbol
		if !isWild(s) {
			findSymbol = s
		}

		linkCount = i + 1
		switch lineSymbol[i] {
		case SS2:
			multiplier += 2
		case SS3:
			multiplier += 3
		case SS5:
			multiplier += 5
		default:
		}
	}

	if multiplier == 0 {
		multiplier = 1
	}

	// 檢查獎圖連線個數是否中獎 (findSymbol 可能是 WW 或 SS)
	if findSymbol != NN && linkCount >= WinSymbolCount[findSymbol] {
		symbol = findSymbol
		specialSymbol := NN
		count = linkCount
		wildCount := 0

		// 检查是否有连续百搭的情况，并比较赔率
		for i := range linkCount {
			if isWild(lineSymbol[i]) {
				// 計算連續百搭數
				wildCount++
				if lineSymbol[i] == WW && specialSymbol != WW {
					specialSymbol = WW
				} else if (lineSymbol[i] == SS || lineSymbol[i] == SS2 || lineSymbol[i] == SS3 || lineSymbol[i] == SS5) && (specialSymbol != SS && specialSymbol != WW) {
					specialSymbol = SS
				}

				// 如果下一个不是百搭，就结束
				if i+1 < linkCount && !isWild(lineSymbol[i+1]) {
					// 选择赔率最高的情况
					if wildCount >= WinSymbolCount[specialSymbol] && wildCount < SLOT_COL {
						if SymbolOdds[specialSymbol][wildCount] >= SymbolOdds[symbol][count] {
							symbol = specialSymbol
							count = wildCount
						}
					} else if wildCount == SLOT_COL {
						// 全百搭
						symbol = specialSymbol
						count = SLOT_COL
					}

					break
				}

			} else {
				break
			}
		}

		// 检查是否需要乘以金蟾倍数
		for i := range linkCount {
			if lineSymbol[i] == WW && wwMultiplier > 1 {
				multiplier *= wwMultiplier
				break
			}
		}
	}

	if symbol > NN {
		multiplier *= SymbolOdds[symbol][count]
	}

	return symbol, count, multiplier
}

// ShowGameSymbol 顯示獎圖盤面
func ShowGameSymbol(file *os.File, gameSymbol GameSymbol) {
	if gameSymbol == nil {
		fmt.Println("ShowGameSymbol: GameSymbol is nil.")
		return
	}
	for row := range SLOT_ROW {
		lineSymbol := make([]Symbol, SLOT_COL)
		for col := range SLOT_COL {
			lineSymbol[col] = gameSymbol[col][row]
		}
		fmt.Fprintln(file, lineSymbol)
	}
	fmt.Fprintln(file, "")
}
