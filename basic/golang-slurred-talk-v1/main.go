package main

import "fmt"

func SlurredTalk(words *string) {
	for i := 0; i < len(*words); i++ {
		if (*words)[i] == 'S' || (*words)[i] == 'R' || (*words)[i] == 'Z' {
			*words = (*words)[:i] + "L" + (*words)[i+1:]
		} else if (*words)[i] == 's' || (*words)[i] == 'r' || (*words)[i] == 'z' {
			*words = (*words)[:i] + "l" + (*words)[i+1:]
		}
	}
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Halo saya daffa afifi syahrony"
	SlurredTalk(&words)
	fmt.Println(words)
}
