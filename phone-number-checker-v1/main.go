package main

import (
	"fmt"
)

func checkProvider(provider string, result *string){
	switch provider {
	case "811", "812", "813", "814", "815":
		*result = "Telkomsel"
	case "816", "817", "818", "819":
		*result = "Indosat"
	case "821", "822", "823":
		*result = "XL"
	case "827", "828", "829":
		*result = "Tri"	
	case "852", "853":
		*result = "AS"
	case "881", "882", "883", "884", "885", "886", "887", "888":
		*result = "Smartfren"
	default:
		*result = "invalid"
	}
}

func PhoneNumberChecker(number string, result *string) {
	if number[0:2] == "08" && len(number) >= 10 {
		provider := number[1:4]
		checkProvider(provider, result)
	} else if number[0:3] == "628" && len(number) >= 11 {
		provider := number[2:5]
		checkProvider(provider, result)
	} else {
		*result = "invalid"
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "628523456789"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
