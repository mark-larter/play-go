package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world.")
}

func main() {
	// Set handlers.
	http.HandleFunc("/", handler)

	// Start http server.
	port := ":8080"
	fmt.Printf("Starting http server on %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
 		log.Panic(err)
	}
}
