package main

import "fmt"

func SlurredTalk(words *string) {
	runes := []rune(*words)

	targetLettersUp := map[rune]bool{
		'R': true, 'S': true, 'Z': true,
	}
	targetLettersLow := map[rune]bool{
		'r': true, 's': true, 'z': true,
	}

	for i, char := range runes {
		if targetLettersUp[char] {
			runes[i] = 'L'
		} else if targetLettersLow[char] {
			runes[i] = 'l'
		}
	}

	*words = string(runes)
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println(words)
}
