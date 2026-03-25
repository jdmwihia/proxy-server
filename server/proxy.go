package server

import (
	"fmt"
	"net/http"
	"log"
)

func landing(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! You've reached the Go server")
}

func Proxy() {
	http.HandleFunc("/", landing)

	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
