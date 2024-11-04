package main

import (
	"fmt"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowel := 0
	consonant := 0
	isWordEmpty := false

	vowelLetters := map[rune]bool{
		'a': true, 'i': true, 'u': true, 'e': true, 'o': true,
		'A': true, 'I': true, 'U': true, 'E': true, 'O': true,
	}

	for _, char := range str {
		if vowelLetters[char] {
			vowel++
		} else if char != ' ' && char != ',' && char != '.' {
			consonant++
		}
	}

	if vowel == 0 || consonant == 0 {
		isWordEmpty = true
	}
	return vowel, consonant, isWordEmpty // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
