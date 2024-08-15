package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", htmlHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
