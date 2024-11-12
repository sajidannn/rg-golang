package main

import (
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)

	if len(data) == 0 {
		return []map[string]interface{}{}
	}

	for _, v := range data {
		name := strings.Split(v, ";")[0]
		age := strings.Split(v, ";")[1]
		address := strings.Split(v, ";")[2]
		height := strings.Split(v, ";")[3]
		isMarried := strings.Split(v, ";")[4]

		intAge, _ := strconv.Atoi(age)

		result = append(result, map[string]interface{}{
			"name": name,
			"age": intAge,
			"address": address,
		})

		if height != "" {
			floatHeight, _ := strconv.ParseFloat(height, 64)
			result[len(result)-1]["height"] = floatHeight
		}

		if isMarried != "" {
			boolIsMarried, _ := strconv.ParseBool(isMarried)
			result[len(result)-1]["isMarried"] = boolIsMarried
		}
	}

	return result
}
