package main

import (
	"fmt"
	"strconv"
)

func ChangeToCurrency(change int) string {
	//Don't worry about this implementation
	//Just use this function :)

	var res = ""

	moneyStr := strconv.Itoa(change)

	for i := len(moneyStr) - 1; i >= 0; i-- {
		if i == len(moneyStr)-1 {
			res = string(moneyStr[i]) + res
		} else if (len(moneyStr)-i)%3 == 0 {
			res = "." + string(moneyStr[i]) + res
		} else {
			res = string(moneyStr[i]) + res
		}
	}

	if res[0] == '.' {
		res = res[1:]
	}

	return "Rp. " + res
}

func MoneyChange(money int, listPrice ...int) string {
	var totalPrice int
	var change int

	for _, price := range listPrice {
		totalPrice += price
	}

	if money >= totalPrice {
		change = money - totalPrice
	} else {
		return "Uang tidak cukup"
	}

	changeCurrency := ChangeToCurrency(change)

	return changeCurrency
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(MoneyChange(150000, 50000, 30000, 30000, 30000, 20000, 50000))
}
