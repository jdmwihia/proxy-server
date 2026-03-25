package server

import (
	"fmt"
	"net/http"
)

func Proxy() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello! You've reached the Go server")
	})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}