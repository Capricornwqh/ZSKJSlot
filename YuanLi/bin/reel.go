package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 打开reel.txt文件
	file, err := os.Open("./reel.txt")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	// 创建扫描器
	scanner := bufio.NewScanner(file)

	// 用于保存原始数据的二维数组
	var data [][]string

	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()
		// 跳过注释行
		if strings.HasPrefix(line, "//") {
			continue
		}

		// 按制表符分割每行
		symbols := strings.Split(line, "\t")
		data = append(data, symbols)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件出错:", err)
		return
	}

	// 获取最大列数
	maxCols := 0
	for _, row := range data {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	// 按列重组数据（转置矩阵）
	columns := make([][]string, maxCols)
	for _, row := range data {
		for colIdx, symbol := range row {
			if symbol != "" { // 跳过空字符串
				columns[colIdx] = append(columns[colIdx], symbol)
			}
		}
	}

	// // 输出格式化的slice
	// fmt.Println("// 按列存储的轮盘数据")
	// fmt.Println("var reelData = [][]string{")
	// for _, col := range columns {
	// 	fmt.Print("\t{")
	// 	for i, symbol := range col {
	// 		fmt.Printf("\"%s\"", symbol)
	// 		if i < len(col)-1 {
	// 			fmt.Print(", ")
	// 		}
	// 	}
	// 	fmt.Println("},")
	// }
	// fmt.Println("}")

	// 也可以输出更简洁的格式，无引号
	fmt.Println("\n// 简洁格式：")
	for i, col := range columns {
		fmt.Print("{")
		for j, symbol := range col {
			fmt.Print(symbol)
			if j < len(col)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("}")
		if i < len(columns)-1 {
			fmt.Print(",\n")
		}
	}
}
