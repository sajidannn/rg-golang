package main

import (
	"strconv"
)

func ReverseData(arr [5]int) [5]int {
	result := [5]int{}
	n:= len(arr)
	for i := 0; i < n; i++ {
		data := strconv.Itoa(arr[n-i-1])
		revData := ""
		for j := len(data) - 1; j >= 0; j-- {
			revData += string(data[j])
		}
		revDataInt, _ := strconv.Atoi(revData)
		result[i] = revDataInt
	}
	return result
}
