package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		w.Write([]byte("Hello App1"))
	})

	log.Println("Starting server on port")
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
