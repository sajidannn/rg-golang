package main

import "fmt"

func CountingNumber(n int) float64 {
	i := 1.0
	result := 0.0

	for i <= float64(n) {
		result = result + i
		i += 0.5
	}
	
	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(5))
}
