package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func filterDay(location, day string) bool {	
	switch location {
	case "JKT":
		if day != "minggu" {
			return true
		}
	case "BDG":
		if day == "rabu" || day == "kamis" || day == "sabtu" {
			return true
		}
	case "BKS":
		if day == "selasa" || day == "kamis" || day == "jumat" {
			return true
		}
	case "DPK":
		if day == "senin" || day == "selasa" {
			return true
		}
	}
	return false
}

func filterPrice(day string) float32 {
	switch day {
	case "senin", "rabu", "jumat":
		return 0.1
	case "selasa", "kamis", "sabtu":
		return 0.05
	}
	return 0
}

func DeliveryOrder(data []string, day string) map[string]float32 {
	deliveryData := make(map[string]float32)

	for _, v := range data {
		fullName := strings.Split(v, ":")[0] + "-" + strings.Split(v, ":")[1]
		location := strings.Split(v, ":")[3]
		price, _ := strconv.Atoi(strings.Split(v, ":")[2])

		isValid := filterDay(location, day)

		if isValid {
			adminFee := filterPrice(day)
			totalPrice := float32(price) + (float32(price) * adminFee)
			deliveryData[fullName] = totalPrice
		}
	}

	return deliveryData
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
