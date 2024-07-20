package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	minNum := 999999999
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

func FindMax(nums ...int) int {
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func SumMinMax(nums ...int) int {
	minNum := FindMin(nums...)
	maxNum := FindMax(nums...)
	result := minNum + maxNum

	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(333, 456, 654, 123, 111, 1000, 1500, 2000, 3000, 1250, 111))
}
