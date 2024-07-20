package main

import (
	"fmt"
	"strconv"
)

// input: arr = [456789, 44332, 2221, 12, 10]
// output: [1, 21, 1222, 23344, 987654]

func ReverseData(arr [5]int) [5]int {
	result := make([]int, 0)
	var resArr [5]int

	for i := len(arr) - 1; i >= 0; i-- {
		res1 := strconv.Itoa(arr[i])
		var reverse string
		for j := len(res1) - 1; j >= 0; j-- {
			reverse += string(res1[j])
		}
		num, _ := strconv.Atoi(reverse)
		result = append(result, num)
	}

	for i := 0; i < 5; i++ {
		resArr[i] = result[i]
	}

	return resArr
}

func main() {
	var arr = [5]int{123, 456, 11, 1, 2}
	fmt.Println(ReverseData(arr))
}
