package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	minNum := nums[0]
	for _, num := range nums[1:] {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

func FindMax(nums ...int) int {
	maxNum := nums[0]
	for _, num := range nums[1:] {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func SumMinMax(nums ...int) int {
	return FindMax(nums...) + FindMin(nums...)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
