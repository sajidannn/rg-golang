package main

import (
	"sort"
)

func SchedulableDays(villager [][]int) []int {
		if len(villager) == 0 {
			return []int{}
		}
	
		dateMap := make(map[int]int)
		for _, date := range villager[0] {
			dateMap[date] = 1
		}
	
		for i := 1; i < len(villager); i++ {
			currentVillager := villager[i]
			tempMap := make(map[int]int)
			for _, date := range currentVillager {
				if _, exists := dateMap[date]; exists {
					tempMap[date] = dateMap[date] + 1
				}
			}
			dateMap = tempMap
		}
	
		var result []int
		for date, count := range dateMap {
			if count == len(villager) {
				result = append(result, date)
			}
		}
	
		sort.Ints(result)
	
		return result
}
