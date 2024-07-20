package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester`
	Studies  []Study `json:"studies"`
}

type Study struct {
	StudyName   string `json:"study_name"`
	StudyCredit int    `json:"study_credit"`
	Grade       string `json:"grade"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	var report Report
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return report, err
	}

	err = json.Unmarshal(jsonData, &report)
	if err != nil {
		return report, err
	}

	return report, nil
}

func GradePoint(report Report) float64 {
	var newGrade float64
	var totalPoints float64
	var totalCredits int

	if len(report.Studies) == 0 {
		return 0.0
	}

	for _, study := range report.Studies {
		switch study.Grade {
		case "A":
			newGrade = 4.0
		case "AB":
			newGrade = 3.5
		case "B":
			newGrade = 3.0
		case "BC":
			newGrade = 2.5
		case "C":
			newGrade = 2.0
		case "CD":
			newGrade = 1.5
		case "D":
			newGrade = 1.0
		case "DE":
			newGrade = 0.5
		case "E":
			newGrade = 0.0
		}

		studyGrade := newGrade * float64(study.StudyCredit)
		totalPoints += studyGrade
		totalCredits += study.StudyCredit
	}

	result := totalPoints / float64(totalCredits)

	return result
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
