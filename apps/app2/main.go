package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
		time.Sleep(2 * time.Second)
		w.Write([]byte("Hello App2"))
	})

	log.Println("Starting server on port")
	err := http.ListenAndServe("127.0.0.1:8082", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
