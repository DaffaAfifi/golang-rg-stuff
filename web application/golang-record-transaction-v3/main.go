package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	var data = make(map[string]int)
	var final = make([]string, 0)

	for _, transaction := range transactions {
		if transaction.Type == "income" {
			data[transaction.Date] += transaction.Amount
		} else if transaction.Type == "expense" {
			data[transaction.Date] -= transaction.Amount
		}
	}

	for key, value := range data {
		if value < 0 {
			strValue := strconv.Itoa(value - value - value)
			strFinal := key + ";expense;" + strValue
			final = append(final, strFinal)
		} else {
			strValue := strconv.Itoa(value)
			strFinal := key + ";income;" + strValue
			final = append(final, strFinal)
		}
	}

	sort.Slice(final, func(i, j int) bool {
		dateI := strings.Split(final[i], ";")[0]
		dateJ := strings.Split(final[j], ";")[0]
		return dateI < dateJ
	})

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(strings.Join(final, "\n"))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "expense", 100000},
		{"01/01/2021", "expense", 1000},
		{"02/01/2021", "expense", 3424},
		{"02/01/2021", "expense", 42000},
		{"03/01/2021", "expense", 22321},
		{"04/01/2021", "expense", 223200},
		{"02/01/2021", "expense", 2300},
		{"05/01/2021", "expense", 2213},
		{"06/01/2021", "expense", 4545},
		{"07/01/2021", "expense", 55500},
		{"08/01/2021", "expense", 200000},
		{"10/01/2021", "expense", 20000},
		{"11/01/2021", "expense", 10000},
		{"12/01/2021", "expense", 55500},
		{"13/01/2021", "expense", 55500},
		{"02/01/2021", "expense", 55500},
		{"02/01/2021", "expense", 10000},
		{"14/01/2021", "expense", 20000},
		{"11/01/2021", "expense", 20000},
		{"15/01/2021", "expense", 10000},
		{"16/01/2021", "expense", 20000},
		{"02/01/2021", "expense", 55500},
		{"17/01/2021", "expense", 10000},
		{"06/01/2021", "expense", 20000},
		{"18/01/2021", "expense", 10000},
		{"03/01/2021", "expense", 20000},
		{"04/01/2021", "expense", 10000},
		{"19/01/2021", "expense", 55500},
		{"20/01/2021", "expense", 55500},
		{"21/01/2021", "expense", 10000},
		{"22/01/2021", "expense", 10000},
		{"23/01/2021", "expense", 10000},
		{"24/01/2021", "expense", 10000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
