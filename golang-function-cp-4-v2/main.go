package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	result := ""
	for _, d := range data {
		if strings.Contains(d, input) {
			if result == "" {
				result = d
				continue
			}
			result = fmt.Sprintf("%s,%s", result, d)
		}
	}
	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
