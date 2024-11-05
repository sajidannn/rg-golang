package main

func SchedulableDays(date1 []int, date2 []int) []int {
	result := make([]int, 0) 
	for _, d1 := range date1 {
		for _, d2 := range date2 {
			if d1 == d2 {
				result = append(result, d1)
			}
		}
	}
return result
}
