package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	result := ""
	words := strings.Split(str, " ")
	for _, word := range words {
		if unicode.IsUpper(rune(word[0])) && !unicode.IsUpper(rune(word[len(word)-1])) {
			word = strings.ToLower(string(word[0])) + word[1:len(word)-1] + strings.ToUpper(string(word[len(word)-1]))
		}
		for i := len(word) - 1; i >= 0; i-- {
			result += string(word[i])
		}
		result += " "
	}
	return result[:len(result)-1]
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
}
