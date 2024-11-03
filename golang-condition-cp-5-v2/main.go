package main

import "fmt"

func TicketPlayground(height, age int) int {
if age < 5 {
	return -1
}

switch {
case age >= 13:
	return 100000
case age >= 12 || height > 160:
	return 60000
case age >= 10 || height > 150:
	return 40000
case age >= 8 || height > 135:
	return 25000
case age >= 5 || height > 120:
	return 15000
}

return -1}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
