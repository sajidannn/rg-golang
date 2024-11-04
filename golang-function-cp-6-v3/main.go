package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ChangeToCurrency(change int) string {
	var res strings.Builder
	moneyStr := strconv.Itoa(change)

	for i := 0; i < len(moneyStr); i++ {
		if i > 0 && (len(moneyStr)-i)%3 == 0 {
			res.WriteString(".")
		}
		res.WriteByte(moneyStr[i])
	}

	return "Rp. " + res.String()
}

func MoneyChange(money int, listPrice ...int) string {
	totalPrice := 0
	for _, price := range listPrice {
		totalPrice += price
	}
	change := money - totalPrice
	if change < 0 {
		return "Uang tidak cukup"
	}
	return ChangeToCurrency(change)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(MoneyChange(100000, 50000, 10000, 10000, 5000, 5000))
}
