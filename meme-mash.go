package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to MemeMash.")
}

func main() {
	fmt.Println("Initializing...")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
