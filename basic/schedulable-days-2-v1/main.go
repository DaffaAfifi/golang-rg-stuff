package main

import "fmt"

// input: villager = [[7, 12, 19, 22], [12, 19, 21, 23], [7, 12, 19], [12, 19]]
// output: [12, 19]

func SchedulableDays(villager [][]int) []int {
	available := make(map[int]int)
	for _, v := range villager {
		for _, tanggal := range v {
			available[tanggal] = available[tanggal] + 1
		}
	}

	result := []int{}
	for k, v := range available {
		if v == len(villager) {
			result = append(result, k)
		}
	}

	return result
}

func main() {
	var data = [][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}
	fmt.Println(SchedulableDays(data))
}
