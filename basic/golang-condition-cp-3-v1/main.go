package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	floatMath, floatScience, floatEnglish, floatIndonesia := float32(math), float32(science), float32(english), float32(indonesia)
	avg := ((floatMath + floatScience + floatEnglish + floatIndonesia) / 4.0)
	var status string

	if avg == 100.0 {
		status = "Sempurna"
	} else if avg >= 90 && avg < 100 {
		status = "Sangat Baik"
	} else if avg >= 80 && avg < 90 {
		status = "Baik"
	} else if avg >= 70 && avg < 80 {
		status = "Cukup"
	} else if avg >= 60 && avg < 70 {
		status = "Kurang"
	} else if avg < 60 {
		status = "Sangat kurang"
	}

	return status
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(100, 100, 100, 100))
}
