package main

func ExchangeCoin(amount int) []int {
	for _, coin := range []int{1000, 500, 200, 100, 50, 20, 10, 5, 1} {
		if amount >= coin {
			return append([]int{coin}, ExchangeCoin(amount-coin)...)
		}
	}
	return []int{}
}
