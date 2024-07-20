package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()
		hari := now.Weekday()
		day := now.Day()
		month := now.Month()
		year := now.Year()
		result := fmt.Sprintf("%s, %d %s %d", hari, day, month, year)

		writer.Write([]byte(result))
	}
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
