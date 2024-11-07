package main

import (
	"fmt"
	"strings"
)

// TODO: answer here

func ChangeOutput (data []string) map[string][]string{
	result := make(map[string][]string)

	var (
		header string
		position string
		value string
	)

	for _, item := range data {
		header = strings.Split(item, "-")[0]
		position = strings.Split(item, "-")[2]

		if position == "first" {
			value = strings.Split(item, "-")[3]
			if header == "phone" {
				result[header] = append(result[header], value)
			}
			continue
		}
		value += " " + strings.Split(item, "-")[3]
		result[header] = append(result[header], value)
	}

	return result
}


// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}

	res := ChangeOutput(data)

	fmt.Println(res)
}
