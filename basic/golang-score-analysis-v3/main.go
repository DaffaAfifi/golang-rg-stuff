package main

import "fmt"

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {
	score := s.Grades

	if len(score) == 0 {
		return 0, 0, 0
	}

	// avg
	var total int
	for i := 0; i < len(score); i++ {
		total += score[i]
	}
	avg := float64(total) / float64(len(score))

	// min
	min := 101
	for i := 0; i < len(score); i++ {
		if score[i] < min {
			min = score[i]
		}
	}

	// max
	max := 0
	for i := 0; i < len(score); i++ {
		if score[i] > max {
			max = score[i]
		}
	}

	return avg, min, max
}

// gunakan untuk melakukan debugging
func main() {
	s := School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 92, 80, 70, 60},
	}
	avg, min, max := Analysis(s)
	s.AddGrade(0, 10, 20)
	avg, min, max = Analysis(s)

	fmt.Println(avg, min, max)
}
