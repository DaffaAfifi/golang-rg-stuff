package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	ticketAmount := (VIP + regular + student)
	totalPrice := ((VIP * 30) + (regular * 20) + (student * 10))
	floatTotalPrice := float32(totalPrice)
	var totalPriceAfterDiscount float32

	if floatTotalPrice >= 100.0 {
		if day%2 == 0 {
			if ticketAmount < 5 {
				totalPriceAfterDiscount = floatTotalPrice * 0.9
				return totalPriceAfterDiscount
			} else {
				totalPriceAfterDiscount = floatTotalPrice * 0.8
				return totalPriceAfterDiscount
			}
		} else {
			if ticketAmount < 5 {
				totalPriceAfterDiscount = floatTotalPrice * 0.85
				return totalPriceAfterDiscount
			} else {
				totalPriceAfterDiscount = floatTotalPrice * 0.75
				return totalPriceAfterDiscount
			}
		}
	} else {
		return floatTotalPrice
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(4, 0, 0, 21))
}
