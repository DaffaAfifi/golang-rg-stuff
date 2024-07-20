package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	tokens := strings.Split(calculate, " ")
	if len(tokens) == 0 {
		return 0
	}

	number, err := strconv.ParseFloat(tokens[0], 32)
	if err != nil {
		return 0
	}

	calc := internal.NewCalculator(float32(number))

	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		operand, _ := strconv.ParseFloat(tokens[i+1], 32)
		if err != nil {
			return 0
		}

		switch operator {
		case "+":
			calc.Add(float32(operand))
		case "-":
			calc.Subtract(float32(operand))
		case "*":
			calc.Multiply(float32(operand))
		case "/":
			calc.Divide(float32(operand))
		default:
			return 0
		}
	}

	return calc.Result()
}

func main() {
	res := AdvanceCalculator("100 / 10 / 10 / 10")

	fmt.Println(res)
}
