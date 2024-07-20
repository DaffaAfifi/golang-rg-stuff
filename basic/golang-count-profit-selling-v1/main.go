package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	var resultMap = make(map[int]int)

	for _, cabang := range data {
		for bulanKe, bulan := range cabang {
			income := bulan[0]
			expense := bulan[1]
			profit := income - expense
			resultMap[bulanKe+1] += profit
		}
	}

	numOfBulan := 0

	for k := range resultMap {
		if k > numOfBulan {
			numOfBulan = k
		}
	}
	result := make([]int, numOfBulan)

	for k, v := range resultMap {
		result[k-1] = v
	}

	return result
}

func main() {
	var numbers = [][][2]int{
		{{1000, 500}, {500, 200}},
		{{1200, 200}, {1000, 800}},
		{{500, 100}, {700, 100}},
	}
	fmt.Println(CountProfit(numbers))
}
