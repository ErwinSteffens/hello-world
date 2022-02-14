package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	log.Println("Hello world running at :8080")

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
