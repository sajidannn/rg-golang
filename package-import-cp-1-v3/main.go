package main

import (
	"fmt"
	"strconv"
	"strings"

	"a21hc3NpZ25tZW50/internal"
)


func AdvanceCalculator(calculate string) float32 {
	if calculate == "" {
		return 0.0
	}

	parts := strings.Split(calculate, " ")
	base, _ := strconv.ParseFloat(parts[0], 32)
	calculator := internal.NewCalculator(float32(base))

	for i := 1; i < len(parts); i += 2 {
		operator := parts[i]
		num, _ := strconv.ParseFloat(parts[i+1], 32)

		switch operator {
		case "+":
			calculator.Add(float32(num))
		case "-":
			calculator.Subtract(float32(num))
		case "*":
			calculator.Multiply(float32(num))
		case "/":
			calculator.Divide(float32(num))
		default:
			return 0.0
		}
	}
	return calculator.Result()
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
