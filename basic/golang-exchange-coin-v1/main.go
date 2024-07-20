package main

import "fmt"

func ExchangeCoin(amount int) []int {
	fraction := make([]int, 0)

	for amount > 0 {
		if amount >= 1000 {
			fraction = append(fraction, 1000)
			amount -= 1000
		} else if amount >= 500 {
			fraction = append(fraction, 500)
			amount -= 500
		} else if amount >= 200 {
			fraction = append(fraction, 200)
			amount -= 200
		} else if amount >= 100 {
			fraction = append(fraction, 100)
			amount -= 100
		} else if amount >= 50 {
			fraction = append(fraction, 50)
			amount -= 50
		} else if amount >= 20 {
			fraction = append(fraction, 20)
			amount -= 20
		} else if amount >= 10 {
			fraction = append(fraction, 10)
			amount -= 10
		} else if amount >= 5 {
			fraction = append(fraction, 5)
			amount -= 5
		} else if amount >= 1 {
			fraction = append(fraction, 1)
			amount -= 1
		}
	}

	return fraction
}

func main() {
	fmt.Println(ExchangeCoin(0))
}
