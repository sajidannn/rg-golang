package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}


func ChangeToStandartTime(time interface{}) string {
	hour := 0
	minute := 0

	switch val := time.(type) {
	case string:
		parts := strings.Split(val, ":")
		if len(parts) != 2 {
			return "Invalid input"
		}
		h, err1 := strconv.Atoi(parts[0])
		m, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return "Invalid input"
		}
		hour = h
		minute = m
	case []int:
		if len(val) != 2 {
			return "Invalid input"
		}
		hour = val[0]
		minute = val[1]
	case map[string]int:
		h, ok1 := val["hour"]
		m, ok2 := val["minute"]
		if !ok1 || !ok2 {
			return "Invalid input"
		}
		hour = h
		minute = m
	case Time:
		hour = val.Hour
		minute = val.Minute
	default:
		return "Invalid input"
	}

	period := "AM"

	if hour > 24 || hour < 0 || minute > 59 || minute < 0 {
		return "Invalid input"
	}

	if hour >= 12 {
		period = "PM"
	}
	if hour > 12 {
		hour -= 12
	} else if hour == 0 {
		hour = 12
	}

	return fmt.Sprintf("%02d:%02d %s", hour, minute, period)
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
