package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type Point struct {
	x, y int // x: 列，y: 行
}

type Toad struct {
	size int
	pos  Point // 左上角坐标
}

// 计算金蟾覆盖的区域
func (t Toad) Area() []Point {
	points := []Point{}
	for i := 0; i < t.size; i++ {
		for j := 0; j < t.size; j++ {
			points = append(points, Point{t.pos.x + i, t.pos.y + j})
		}
	}
	return points
}

// 计算金蟾到铜钱的最小距离
func minDistance(toad Toad, coin Point) int {
	minDist := math.MaxInt32
	for _, p := range toad.Area() {
		d := abs(p.x-coin.x) + abs(p.y-coin.y)
		if d < minDist {
			minDist = d
		}
	}
	return minDist
}

// 获取点相对于金蟾的方位区域
// 返回值: 0=右边, 1=左边, 2=上边, 3=下边
func getDirectionArea(toad Toad, point Point) int {
	// 金蟾的范围
	toadMinX := toad.pos.x
	toadMaxX := toad.pos.x + toad.size - 1
	toadMinY := toad.pos.y
	toadMaxY := toad.pos.y + toad.size - 1

	// 右边区域
	if point.x > toadMaxX {
		return 0
	}
	// 左边区域
	if point.x < toadMinX {
		return 1
	}
	// 上边区域
	if point.y < toadMinY {
		return 2
	}
	// 下边区域
	if point.y > toadMaxY {
		return 3
	}

	// 金蟾覆盖区域内
	return -1
}

// 计算点在方位区域内的优先级（值越小优先级越高）
func getPositionPriority(toad Toad, point Point, area int) int {
	// 计算金蟾中心点
	centerX := toad.pos.x + (toad.size-1)/2
	centerY := toad.pos.y + (toad.size-1)/2

	// 计算到金蟾中心的曼哈顿距离
	distToCenter := abs(point.x-centerX) + abs(point.y-centerY)

	return distToCenter
}

// 找到距离金蟾最近的铜钱
func nextTarget(toad Toad, coins []Point) Point {
	type candidate struct {
		point            Point
		distance         int // 铜钱到金蟾的最小距离
		directionArea    int // 方位区域: 0=右, 1=左, 2=上, 3=下
		positionPriority int // 在方位内的优先级
	}

	var cands []candidate
	minDist := math.MaxInt32

	// 1. 找出距离最近的铜钱
	for _, coin := range coins {
		d := minDistance(toad, coin)
		if d < minDist {
			minDist = d
			area := getDirectionArea(toad, coin)
			priority := getPositionPriority(toad, coin, area)
			cands = []candidate{{coin, d, area, priority}}
		} else if d == minDist {
			area := getDirectionArea(toad, coin)
			priority := getPositionPriority(toad, coin, area)
			cands = append(cands, candidate{coin, d, area, priority})
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
		if cands[i].point.x != cands[j].point.x {
			return cands[i].point.x > cands[j].point.x
		}
		return cands[i].point.y < cands[j].point.y
	})

	return cands[0].point
}

// 移动金蟾到目标铜钱位置
func moveTo(toad *Toad, target Point) {
	minDist := math.MaxInt32
	var newTopLeft Point
	for i := 0; i < toad.size; i++ {
		for j := 0; j < toad.size; j++ {
			tlx, tly := target.x-i, target.y-j
			if tlx < 0 || tly < 0 || tlx+toad.size > 6 || tly+toad.size > 6 {
				continue
			}
			d := abs(toad.pos.x-tlx) + abs(toad.pos.y-tly)
			if d < minDist {
				minDist = d
				newTopLeft = Point{tlx, tly}
			}
		}
	}
	toad.pos = newTopLeft
}

// 是否有铜钱被金蟾覆盖
func isCovered(toad Toad, coin Point) bool {
	for _, p := range toad.Area() {
		if p == coin {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func printGrid(toad Toad, coins []Point, step int, caseNum int) {
	grid := [6][6]rune{}
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for _, c := range coins {
		grid[c.y][c.x] = '$'
	}
	for _, p := range toad.Area() {
		grid[p.y][p.x] = 'T'
	}
	fmt.Printf("\033[2J\033[H")
	fmt.Printf("测试用例 %d, 步骤 %d:\n", caseNum, step)
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Printf("%c ", grid[i][j])
		}
		fmt.Println()
	}
	time.Sleep(1000 * time.Millisecond)
}

// 运行单个测试用例的模拟
func runSimulation(toad Toad, coins []Point, caseNum int) []Point {
	step := 0
	printGrid(toad, coins, step, caseNum)
	sortCoins := []Point{}

	for len(coins) > 0 {
		step++
		target := nextTarget(toad, coins)
		fmt.Printf("\n金蟾当前位置: %+v，移动目标铜钱: %+v\n", toad.pos, target)
		moveTo(&toad, target)

		newCoins := []Point{}
		sortCoins = append(sortCoins, target)
		for _, c := range coins {
			if !isCovered(toad, c) {
				newCoins = append(newCoins, c)
			} else {
				if c == target {
					continue
				}
				fmt.Printf("吃掉铜钱: %+v\n", c)
				sortCoins = append(sortCoins, c)
			}
		}
		coins = newCoins
		printGrid(toad, coins, step, caseNum)
	}

	return sortCoins
}

func main() {
	// 定义多组测试用例
	testCases := []struct {
		toad  Toad
		coins []Point
	}{
		// {
		// 	toad:  Toad{size: 2, pos: Point{0, 0}},
		// 	coins: []Point{{0, 4}, {1, 4}, {2, 4}, {3, 3}, {4, 2}, {5, 0}},
		// },
		// {
		// 	toad:  Toad{size: 2, pos: Point{2, 2}},
		// 	coins: []Point{{4, 4}, {2, 5}, {0, 5}, {1, 1}, {3, 0}, {5, 0}},
		// },
		// {
		// 	toad:  Toad{size: 2, pos: Point{3, 3}},
		// 	coins: []Point{{1, 3}, {2, 2}, {3, 1}, {4, 1}, {5, 2}},
		// },
		{
			toad:  Toad{size: 3, pos: Point{0, 0}},
			coins: []Point{{0, 5}, {1, 5}, {2, 5}, {3, 4}, {4, 3}, {5, 0}},
		},
		{
			toad:  Toad{size: 3, pos: Point{2, 2}},
			coins: []Point{{0, 2}, {1, 5}, {2, 0}, {3, 5}, {4, 5}, {5, 1}},
		},
		{
			toad:  Toad{size: 3, pos: Point{1, 1}},
			coins: []Point{{0, 1}, {1, 0}, {2, 0}, {3, 4}, {4, 3}},
		},
	}

	// 存储每个测试用例的结果
	results := make([][]Point, len(testCases))

	// 运行所有测试用例
	for i, tc := range testCases {
		fmt.Printf("\n开始运行测试用例 %d\n", i+1)
		// 创建副本以防修改原始数据
		toad := tc.toad
		coins := make([]Point, len(tc.coins))
		copy(coins, tc.coins)

		results[i] = runSimulation(toad, coins, i+1)
	}

	// 打印所有测试用例的结果摘要
	fmt.Println("\n\n=======================================")
	fmt.Println("所有测试用例运行完毕，结果摘要如下：")
	fmt.Println("=======================================")

	for i, result := range results {
		fmt.Printf("\n测试用例 %d 铜钱吃掉顺序：\n", i+1)
		for j, coin := range result {
			fmt.Printf("  %d. 位置(%d, %d)\n", j+1, coin.x, coin.y)
		}
	}
}
