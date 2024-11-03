package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	price := VIP * 30 + regular * 20 + student * 10
	var priceAfterDiscount float32 

	if price >= 100 {
		if day % 2 == 1 {
			if VIP + regular + student >= 5 {
				priceAfterDiscount = float32(price) - float32(price) * 0.25
			} else {
				priceAfterDiscount = float32(price) - float32(price) * 0.15
			}
		}	else {
			if VIP + regular + student >= 5 {
				priceAfterDiscount = float32(price) - float32(price) * 0.20
			} else {
				priceAfterDiscount = float32(price) - float32(price) * 0.10
			}
		}
		return priceAfterDiscount
	}
	
	return float32(price)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
