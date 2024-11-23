package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	sort.Slice(transactions, func(i, j int) bool {
		dateI, _ := time.Parse("02/01/2006", transactions[i].Date)
		dateJ, _ := time.Parse("02/01/2006", transactions[j].Date)
		return dateI.Before(dateJ)
	})

	currentDate := transactions[0].Date
	totalAmount := 0
	trans := []string{}

	for _, transaction := range transactions {
		if transaction.Date == currentDate {
			if transaction.Type == "income" {
				totalAmount += transaction.Amount
			} else {
				totalAmount -= transaction.Amount
			}
		} else {
			status := "income"
			if totalAmount < 0 {
				status = "expense"
				totalAmount = int(math.Abs(float64(totalAmount)))
			}
			data := fmt.Sprintf("%s;%s;%d", currentDate, status, totalAmount)
			trans = append(trans, data)

			currentDate = transaction.Date
			totalAmount = 0
			if transaction.Type == "income" {
				totalAmount += transaction.Amount
			} else {
				totalAmount -= transaction.Amount
			}
		}
	}

	status := "income"
	if totalAmount < 0 {
		status = "expense"
		totalAmount = int(math.Abs(float64(totalAmount)))
	}
	data := fmt.Sprintf("%s;%s;%d", currentDate, status, totalAmount)
	trans = append(trans, data)

	output := strings.Join(trans, "\n")
	err := os.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}



func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"02/01/2021", "expense", 10000},
		{"02/01/2021", "expense", 10000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
