package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	numStr := strconv.Itoa(numbers) 
	maxSum := 0                    
	maxPair := 0                  

	for i := 0; i < len(numStr)-1; i++ {
		pair := numStr[i:i+2]
		sum := int(pair[0]-'0') + int(pair[1]-'0')

		if sum > maxSum {
			maxSum = sum
			maxPair, _ = strconv.Atoi(pair) // menyimpan pasangan sebagai angka
		}
	}

	return maxPair
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
