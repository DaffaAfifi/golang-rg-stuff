package main

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	totalPrice := 0
	totalTax := 0

	for _, product := range products {
		totalPrice += product.Price
		totalTax += product.Tax
	}

	change := amount - (totalPrice + totalTax)
	fraction := make([]int, 0)
	for change > 0 {
		if change >= 1000 {
			fraction = append(fraction, 1000)
			change -= 1000
		} else if change >= 500 {
			fraction = append(fraction, 500)
			change -= 500
		} else if change >= 200 {
			fraction = append(fraction, 200)
			change -= 200
		} else if change >= 100 {
			fraction = append(fraction, 100)
			change -= 100
		} else if change >= 50 {
			fraction = append(fraction, 50)
			change -= 50
		} else if change >= 20 {
			fraction = append(fraction, 20)
			change -= 20
		} else if change >= 10 {
			fraction = append(fraction, 10)
			change -= 10
		} else if change >= 5 {
			fraction = append(fraction, 5)
			change -= 5
		} else if change >= 1 {
			fraction = append(fraction, 1)
			change -= 1
		}
	}

	return fraction
}
