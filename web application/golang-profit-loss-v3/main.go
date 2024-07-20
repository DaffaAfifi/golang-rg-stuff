package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	var data, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		emptyArr := []string{}
		return emptyArr, nil
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func CalculateProfitLoss(data []string) string {
	var total = 0
	var msg string
	transactions := make(map[string]int)

	for _, line := range data {
		parts := strings.Split(line, ";")
		amount, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			panic(err)
		}

		if parts[1] == "income" {
			transactions[parts[1]]++
			total += amount
		} else if parts[1] == "expense" {
			transactions[parts[1]]++
			total -= amount
		}
	}

	lastData := data[len(data)-1]
	date := strings.Split(lastData, ";")
	msgDate := date[0]
	lossTotal := strconv.Itoa(total - total - total)
	strTotal := strconv.Itoa(total)

	fmt.Println(lossTotal, strTotal)

	if transactions["income"] == 0 {
		msg = msgDate + ";loss;" + lossTotal
		return msg
	}

	if total > 0 {
		msg = msgDate + ";profit;" + strTotal
	} else {
		msg = msgDate + ";loss;" + strTotal
	}

	return msg
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
