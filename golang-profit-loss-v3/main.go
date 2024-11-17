package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	var content, err = os.ReadFile(path)
	if err != nil {
			return nil, err
	}
	if len(content) == 0 {
		return []string{}, nil
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func CalculateProfitLoss(data []string) string {
	total := 0
	date := ""
	for _, line := range data {
		splittedLine := strings.Split(line, ";")
		value, _ := strconv.Atoi(splittedLine[2])
		
		if splittedLine[1] == "income" {
			total += value
		} else if splittedLine[1] == "expense" {
			total -= value
		}
		date = splittedLine[0]
	}

	if total > 0 {
		return fmt.Sprintf("%s;profit;%d", date, total)
	} 
	return fmt.Sprintf("%s;loss;%d", date, int(math.Abs(float64(total))))
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
