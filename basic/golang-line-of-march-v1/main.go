package main

import "fmt"

// input: heights : [172, 170, 150, 165, 144, 155, 159]
// output: [144, 150, 155, 159, 165, 170, 172]

func Sortheight(height []int) []int {
	result := make([]int, len(height))
	copy(result, height)

	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				temp := result[j]
				result[j] = result[j+1]
				result[j+1] = temp
			}
		}
	}

	return result
}

func main() {
	var heights = []int{172, 170, 150, 165, 144, 155, 159}
	fmt.Println(Sortheight(heights))
}
