package main

func CountProfit(data [][][2]int) []int {
	if len(data) == 0 {
		return []int{}
	}

	numMonths := len(data[0])
	result := make([]int, numMonths)

	for month := 0; month < numMonths; month++ {
		totalProfit := 0

		for _, branch := range data {
			if month < len(branch) {
				sales := branch[month][0]
				expense := branch[month][1]
				profit := sales - expense
				totalProfit += profit
			}
		}

		result[month] = totalProfit
	}

	return result
}
