package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	average := (math + science + english + indonesia) / 4

	switch {
	case average == 100:
		return "Sempurna"
	case average >= 90:
		return "Sangat Baik"
	case average >= 80:
		return "Baik"
	case average >= 70:
		return "Cukup"
	case average >= 60:
		return "Kurang"
	case average < 60:
		return "Sangat kurang"
	}

	return "" 
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 50, 10, 60))
}
