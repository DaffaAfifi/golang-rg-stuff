package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {
	result := make(map[string]float32)

	for _, d := range data {
		tokens := strings.Split(d, ":")

		firstName := tokens[0]
		lastName := tokens[1]
		price, _ := strconv.ParseFloat(tokens[2], 32)
		city := tokens[3]

		if city == "BDG" {
			if !(day == "rabu" || day == "kamis" || day == "sabtu") {
				continue
			}
		}
		if city == "BKS" {
			if !(day == "selasa" || day == "kamis" || day == "jumat") {
				continue
			}
		}
		if city == "DPK" {
			if !(day == "senin" || day == "selasa") {
				continue
			}
		}
		if city == "JKT" {
			if day == "minggu" {
				continue
			}
		}

		switch day {
		case "senin", "rabu", "jumat":
			price = price * 1.1
		case "selasa", "kamis", "sabtu":
			price = price * 1.05
		}

		result[firstName+"-"+lastName] = float32(price)
	}
	return result
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
