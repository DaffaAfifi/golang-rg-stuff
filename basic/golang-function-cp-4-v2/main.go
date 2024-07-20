package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	result := ""
	yangPertama := true

	for _, v := range data {
		fmt.Println(v)
		apakahMengandungInput := strings.Contains(v, input)
		fmt.Println(apakahMengandungInput)

		if apakahMengandungInput {
			if yangPertama {
				result += "," + v
				yangPertama = false
			} else {
				result += "," + v
			}
		}
	}
	fmt.Println(result[:len(result)-1])
	return ""
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("motor", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}

// curItem := ""
// simItem := ""
// var items = strings.Join(data, ",")

// for _, item := range items {
// 	if string(item) == "," {
// 		if strings.Contains(curItem, input) {
// 			simItem += curItem + ","
// 		}
// 		curItem = ""
// 	} else {
// 		curItem += string(item)
// 	}
// }

// if curItem != "" {
// 	if strings.Contains(curItem, input) {
// 		simItem += curItem + ","
// 	}
// }

// finalItem := simItem[:len(simItem)-1]
// return finalItem
