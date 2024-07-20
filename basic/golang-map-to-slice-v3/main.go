package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	var result [][]string

	for key, value := range mapData {
		res := []string{key, value}
		result = append(result, res)
	}
	return result // TODO: replace this
}

func main() {
	data := map[string]string{
		"hello": "world",
		"John":  "Doe",
		"Age":   "14",
	}

	fmt.Println(MapToSlice(data))
}
