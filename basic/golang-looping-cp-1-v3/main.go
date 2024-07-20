package main

import "fmt"

func CountingNumber(n int) float64 {
	number := 0.0

	for i := 1.0; i <= float64(n); i += 0.5 {
		number += i
	}

	return number
}

func main() {
	fmt.Println(CountingNumber(100))
}
