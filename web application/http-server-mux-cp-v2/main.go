package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		hari := now.Weekday()
		day := now.Day()
		month := now.Month()
		year := now.Year()
		result := fmt.Sprintf("%s, %d %s %d", hari, day, month, year)

		w.Write([]byte(result))
	}
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.Write([]byte("Hello there"))
			return
		}
		result := fmt.Sprintf("Hello, %s!", name)
		w.Write([]byte(result))
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())

	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
