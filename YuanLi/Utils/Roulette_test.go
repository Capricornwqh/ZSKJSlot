package Utils

import (
	"fmt"
	"testing"
)

func TestRouletteAdd(t *testing.T) {
	itemMap := map[int]uint{1: 10, 2: 20, 3: 30, 4: 40}
	roulette := NewRoulette[int]()
	for k, v := range itemMap {
		roulette.Add(k, v)
	}
	roulette.Dump()
}

func TestRouletteNewFromList(t *testing.T) {
	itemList := []uint{10, 20, 30, 40, 50}
	roulette := NewRouletteFromList(itemList)
	roulette.Dump()
}

func TestRouletteNewFromMap(t *testing.T) {
	itemMap := map[int32]uint{1: 10, 2: 20, 3: 30, 4: 40}
	roulette := NewRouletteFromMap[int32](itemMap)
	roulette.Dump()
}

func TestNewRouletteFromTwoList(t *testing.T) {
	elemList := []float64{-4, 0.5, 1, 1.5, 2}
	weightList := []uint{10, 20, 30, 40, 50}
	roulette := NewRouletteFromTwoList(elemList, weightList)
	roulette.Dump()
}

func TestRouletteSpin(t *testing.T) {
	const MaxSpinCount = 1000000000
	itemMap := map[int]uint{1: 10, 2: 20, 3: 30, 4: 40}
	roulette := NewRoulette[int]()
	for k, v := range itemMap {
		roulette.Add(k, v)
	}

	countMap := map[int]int{}
	for i := 0; i < MaxSpinCount; i++ {
		symbol, ok := roulette.Spin()
		if ok {
			countMap[symbol]++
		}
	}

	fmt.Printf("Spin count: %d\n", MaxSpinCount)
	for k, v := range countMap {
		fmt.Printf("Element = %d, Count = %d, Ratio = %.3f%%\n", k, v, float64(v)/float64(MaxSpinCount)*100)
	}
}
