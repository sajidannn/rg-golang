package main

import "fmt"

func CountingLetter(text string) int {
	total := 0
	for _, letter := range text {
		if letter == 'R' || letter == 'S' || letter == 'T' || letter == 'Z' ||  letter == 'r' || letter == 's' || letter == 't' || letter == 'z'{
			total++
		}
	}
	return total

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
