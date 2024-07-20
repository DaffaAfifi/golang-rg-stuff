package main

import "fmt"

// input: date1 = [1, 2, 3, 4], date2 = [3, 4, 5]
// output: [3, 4]

func SchedulableDays(date1 []int, date2 []int) []int {
	var imam = make([]int, len(date1))
	var permana = make([]int, len(date2))
	var result = make([]int, 0)
	copy(imam, date1)
	copy(permana, date2)

	for i := 0; i < len(imam); i++ {
		temp1 := imam[i]
		for i := 0; i < len(permana); i++ {
			temp2 := permana[i]
			if temp1 == temp2 {
				result = append(result, temp2)
			}
		}
	}

	return result
}

func main() {
	var date1 = []int{1, 2, 3, 4}
	var date2 = []int{3, 4, 5}
	fmt.Println(SchedulableDays(date1, date2))
}
