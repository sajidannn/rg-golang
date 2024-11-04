package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	dayStr := fmt.Sprintf("%02d", day)

	switch month {
	case 1:
		return fmt.Sprintf("%s-January-%d", dayStr, year)
	case 2:
		return fmt.Sprintf("%s-February-%d", dayStr, year)
	case 3:
		return fmt.Sprintf("%s-March-%d", dayStr, year)
	case 4:
		return fmt.Sprintf("%s-April-%d", dayStr, year)
	case 5:
		return fmt.Sprintf("%s-May-%d", dayStr, year)
	case 6:
		return fmt.Sprintf("%s-June-%d", dayStr, year)
	case 7:
		return fmt.Sprintf("%s-July-%d", dayStr, year)
	case 8:
		return fmt.Sprintf("%s-August-%d", dayStr, year)
	case 9:
		return fmt.Sprintf("%s-September-%d", dayStr, year)
	case 10:
		return fmt.Sprintf("%s-October-%d", dayStr, year)
	case 11:
		return fmt.Sprintf("%s-November-%d", dayStr, year)
	case 12:
		return fmt.Sprintf("%s-December-%d", dayStr, year)
	}
	return ""
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
