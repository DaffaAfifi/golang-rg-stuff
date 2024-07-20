package main

import (
	"fmt"
)

func CountVowelConsonant(str string) (int, int, bool) {
	var vokal string
	var konsonan string
	var trash string
	var check bool = false

	for i := 0; i < len(str); i++ {
		if string(str[i]) == "A" || string(str[i]) == "a" || string(str[i]) == "I" || string(str[i]) == "i" || string(str[i]) == "U" || string(str[i]) == "u" || string(str[i]) == "E" || string(str[i]) == "e" || string(str[i]) == "O" || string(str[i]) == "o" {
			vokal += string(str[i])
		} else if string(str[i]) == " " || string(str[i]) == "," {
			trash += string(str[i])
		} else {
			konsonan += string(str[i])
		}
	}

	if len(vokal) == 0 || len(konsonan) == 0 {
		check = true
	}

	return len(vokal), len(konsonan), check
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
}
