package main

import (
	"net/http"
)

func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
	}
}

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
		}
	})
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/student", StudentHandler())
	HandleWithMiddleware := RequestMethodGet(mux)

	http.ListenAndServe("localhost:8080", HandleWithMiddleware)
}
