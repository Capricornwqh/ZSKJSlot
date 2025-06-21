package Utils

import (
	"fmt"
	"math/rand/v2"
)

// Roulette 輪盤
type Roulette[T comparable] struct {
	sum         uint
	weightList  []uint
	elementList []T
}

// Add 新增元素及其比重至輪盤中
func (r *Roulette[T]) Add(elem T, ratio uint) bool {
	if ratio == 0 {
		return false
	}
	if r.Exists(elem) {
		fmt.Printf("Element %v already exists.\n", elem)
		return false
	}
	r.sum += ratio
	r.weightList = append(r.weightList, r.sum)
	r.elementList = append(r.elementList, elem)
	return true
}

func (r *Roulette[T]) Size() int {
	return len(r.elementList)
}

func (r *Roulette[T]) IsEmpty() bool {
	return len(r.elementList) == 0
}

func (r *Roulette[T]) Clear() {
	r.sum = 0
	r.weightList = make([]uint, 0)
	r.elementList = make([]T, 0)
}

func (r *Roulette[T]) Exists(elem T) bool {
	for _, el := range r.elementList {
		if el == elem {
			return true
		}
	}
	return false
}

// Spin 隨機骰出一個輪盤中的元素
func (r *Roulette[T]) Spin() (T, bool) {
	if r.sum == 0 || len(r.elementList) == 0 {
		return *new(T), false
	}
	var dice = rand.UintN(r.sum)
	for index, ratio := range r.weightList {
		if dice < ratio {
			elem := r.elementList[index]
			// fmt.Printf("[Spin] Dice = %d, Ratio = %v, Element = %v.\n", dice, ratio, elem)
			return elem, true
		}
	}
	return *new(T), false
}

func (r *Roulette[T]) Dump() {
	fmt.Println("-- Dump Roulette --")
	for i, v := range r.weightList {
		if i >= 1 {
			fmt.Printf("Ratio = %v, Element = %v\n", v-r.weightList[i-1], r.elementList[i])
		} else {
			fmt.Printf("Ratio = %v, Element = %v\n", v, r.elementList[i])
		}
	}
	fmt.Printf("RatioSum = %v\n", r.sum)
}

func NewRoulette[T comparable]() *Roulette[T] {
	return &Roulette[T]{}
}

func NewRouletteFromList(list []uint) *Roulette[int] {
	var r = NewRoulette[int]()
	for i, v := range list {
		r.Add(i, v)
	}
	return r
}

func NewRouletteFromMap[T comparable](m map[T]uint) *Roulette[T] {
	var r = NewRoulette[T]()
	for k, v := range m {
		r.Add(k, v)
	}
	return r
}

func NewRouletteFromTwoList[T comparable](elemList []T, weightList []uint) *Roulette[T] {
	if len(elemList) != len(weightList) || len(elemList) == 0 {
		panic("Element list is invalid.")
	}
	var r = NewRoulette[T]()
	for i, elem := range elemList {
		r.Add(elem, weightList[i])
	}
	return r
}
