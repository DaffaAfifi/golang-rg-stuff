package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var floatHeight = float64(height)
	var bmiNumber float64
	if gender == "laki-laki" {
		bmiNumber = (floatHeight - 100.0) - ((floatHeight - 100.0) * 0.10)
	} else if gender == "perempuan" {
		bmiNumber = (floatHeight - 100.0) - ((floatHeight - 100.0) * 0.15)
	}
	return bmiNumber
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 170))
	fmt.Println(BMICalculator("perempuan", 165))
}
