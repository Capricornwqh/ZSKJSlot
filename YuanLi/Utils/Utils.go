package Utils

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"reflect"
	"sort"
)

// GetPercentage 取得百分比顯示字串
func GetPercentage(numerator, denominator float64) string {
	if denominator == 0 {
		return "NaN"
	}
	return fmt.Sprintf("%.4f%%", numerator*100/denominator)
}

// SortedMapKeys 取得 map 排序過後的 keys
func SortedMapKeys(m interface{}) (keyList []int) {
	keys := reflect.ValueOf(m).MapKeys()

	for _, key := range keys {
		keyList = append(keyList, key.Interface().(int))
	}
	sort.Ints(keyList)
	return
}

// RandChoiceByWeight 根據權重隨機傳回一個元素
func RandChoiceByWeight[T any](sequence []T, weights []uint) T {
	if len(sequence) != len(weights) || len(sequence) == 0 {
		panic("Invalid length.")
	}

	var sum = 0
	for i := 0; i < len(weights); i++ {
		sum += int(weights[i])
	}

	var dice = rand.IntN(sum)
	for i := 0; i < len(weights); i++ {
		dice -= int(weights[i])
		if dice < 0 {
			return sequence[i]
		}
	}
	panic("Rand failed.")
}

// RandChoiceByCumWeight 根據累積權重隨機傳回一個元素
func RandChoiceByCumWeight[T any](sequence []T, cumWeights []uint) T {
	if len(sequence) != len(cumWeights) {
		panic("Invalid length.")
	}

	var dice = rand.UintN(cumWeights[len(cumWeights)-1])
	for i := 0; i < len(cumWeights); i++ {
		if dice < cumWeights[i] {
			return sequence[i]
		}
	}
	panic("Rand failed.")
}

// CheckMaxWin 比較並回傳最大贏分限制
func CheckMaxWin[T cmp.Ordered](win T, maxWin T) (T, bool) {
	if win > maxWin {
		return maxWin, true
	}
	return win, false
}

// Sum 加總陣列中的值
func Sum[T cmp.Ordered](list []T) T {
	var sum T
	for _, v := range list {
		sum += v
	}
	return sum
}

// RemoveMaximum 移除列表中最大值
func RemoveMaximum[T cmp.Ordered](list []T) []T {
	if len(list) <= 1 {
		return []T{}
	}
	var maxValue T
	var index = -1
	for i, value := range list {
		if value > maxValue {
			maxValue = value
			index = i
		}
	}
	if index < 0 {
		return list
	}
	// fmt.Printf("RemoveMaximum: [%d] %v\n", index, maxValue)
	return append(list[:index], list[index+1:]...)
}
