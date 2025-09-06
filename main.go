package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The backend is running")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	})

	fmt.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func printer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
