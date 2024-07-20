package main

import (
	"fmt"
)

func FindShortestName(names string) string {
	curName := ""
	minName := "ddddddddddddddddddddddddddddddddddddd"

	for _, c := range names {
		if string(c) == ";" || string(c) == " " || string(c) == "," {
			if len(curName) < len(minName) {
				minName = curName
			} else if len(curName) == len(minName) {
				if curName < minName {
					minName = curName
				}
			}
			curName = ""
		} else {
			curName += string(c)
		}
	}

	if curName != "" {
		if len(curName) < len(minName) {
			minName = curName
		} else if len(curName) == len(minName) {
			if curName < minName {
				minName = curName
			}
		}
	}

	return minName
}

// gunakan untuk melakukan debug
func main() {
	// fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tiz;Tio")) // "Tia"
}
