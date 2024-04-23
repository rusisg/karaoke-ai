package main

import (
	"log"
	"net/http"
)

const port = "8080"

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/api/generate", GenerateTextHandler)

	log.Printf("Starting server at :", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
