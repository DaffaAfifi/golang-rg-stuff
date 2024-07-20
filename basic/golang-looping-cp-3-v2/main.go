package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	var unreadLetters string

	for i := 0; i <= len(text)-1; i++ {
		if string(text[i]) == "R" || string(text[i]) == "r" || string(text[i]) == "S" || string(text[i]) == "s" || string(text[i]) == "T" || string(text[i]) == "t" || string(text[i]) == "Z" || string(text[i]) == "z" {
			unreadLetters += string(text[i])
		}
	}

	amountUnreadLetters := len(unreadLetters)

	return amountUnreadLetters

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Zebra Zig Zag"))
}
