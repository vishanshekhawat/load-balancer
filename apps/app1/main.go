package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello App1"))
	})

	log.Println("Starting server on port")
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
