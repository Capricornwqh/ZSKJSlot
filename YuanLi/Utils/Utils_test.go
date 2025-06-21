package Utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRandChoiceByWeight(t *testing.T) {
	type args struct {
		totalCount int
		sequence   []int32
		weights    []uint
	}
	tests := []struct {
		name string
		args args
	}{
		// Test cases
		{
			"根據權重隨機傳回一個元素",
			args{totalCount: 1000000, sequence: []int32{1, 2, 3, 4, 5}, weights: []uint{10, 10, 10, 10, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var countMap = make(map[int]int)
			for i := 0; i < tt.args.totalCount; i++ {
				var s = RandChoiceByWeight(tt.args.sequence, tt.args.weights)
				countMap[int(s)]++
			}
			var keys = SortedMapKeys(countMap)
			for _, key := range keys {
				fmt.Printf("%v Count: %d (%s)\n", key, countMap[key], GetPercentage(float64(countMap[key]), float64(tt.args.totalCount)))
			}
		})
	}
}

func TestRandChoiceByCumWeight(t *testing.T) {
	type args struct {
		totalCount int
		sequence   []int32
		cumWeights []uint
	}
	tests := []struct {
		name string
		args args
	}{
		// Test cases
		{
			"根據累積權重隨機傳回一個元素",
			args{totalCount: 1000000, sequence: []int32{1, 2, 3, 4, 5}, cumWeights: []uint{10, 20, 30, 40, 50}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var countMap = make(map[int]int)
			for i := 0; i < tt.args.totalCount; i++ {
				var s = RandChoiceByCumWeight(tt.args.sequence, tt.args.cumWeights)
				countMap[int(s)]++
			}
			var keys = SortedMapKeys(countMap)
			for _, key := range keys {
				fmt.Printf("%v Count: %d (%s)\n", key, countMap[key], GetPercentage(float64(countMap[key]), float64(tt.args.totalCount)))
			}
		})
	}
}

func TestCheckMaxWin(t *testing.T) {
	type args struct {
		win    uint64
		maxWin uint64
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 bool
	}{
		// Test cases
		{
			"比較並回傳最大贏分 True",
			args{win: 300, maxWin: 200},
			200,
			true,
		},
		{
			"比較並回傳最大贏分 False",
			args{win: 100, maxWin: 200},
			100,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckMaxWin(tt.args.win, tt.args.maxWin)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckMaxWin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckMaxWin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Test cases
		{
			"加總陣列中的值 55",
			args{list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveMaximum(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// Test cases
		{
			"移除列表中最大值 1",
			args{list: []int{1}},
			[]int{},
		},
		{
			"移除列表中最大值 10",
			args{list: []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}},
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveMaximum(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
