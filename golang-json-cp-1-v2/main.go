package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Study struct {
	StudyName string `json:"study_name"`
	StudyCredit int `json:"study_credit"`
	Grade string `json:"grade"`
}

type Report struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
	Semester int `json:"semester"`
	Studies []Study `json:"studies"`
}


// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	var report Report

	file, err := os.ReadFile(filename)
	if err != nil {
		return Report{}, fmt.Errorf("failed to open file: %w", err)
	}

	err = json.Unmarshal([]byte(file), &report)
	if err != nil {
		return Report{}, err
	}

	return report, nil
}

func GradePoint(report Report) float64 {
	if len(report.Studies) == 0 {
		return 0.0
	}
	
	IPs := 0.0
	totalCredit := 0
	for _, study := range report.Studies {
		switch study.Grade {
		case "A":
			IPs += 4.0 * float64(study.StudyCredit)
		case "AB":
			IPs += 3.5 * float64(study.StudyCredit)
		case "B":
			IPs += 3.0 * float64(study.StudyCredit)
		case "BC":
			IPs += 2.5 * float64(study.StudyCredit)
		case "C":
			IPs += 2.0 * float64(study.StudyCredit)
		case "CD":
			IPs += 1.5 * float64(study.StudyCredit)
		case "D":
			IPs += 1.0 * float64(study.StudyCredit)
		case "DE":
			IPs += 0.5 * float64(study.StudyCredit)
		case "E":
			IPs += 0.0 * float64(study.StudyCredit)
		}
		totalCredit += study.StudyCredit
	}
	return IPs/float64(totalCredit)
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
