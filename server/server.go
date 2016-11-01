package main

import (
	"log"
	"net/http"
	"sync"
)

// Todo is a representation of each item that
// we are going to maintain in our service
type Todo struct {
	ID        uint64
	Task      string
	Completed bool
}

var todos []Todo
var lock sync.RWMutex

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Authenticate and return JWT
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Return all the Todo items
		} else if r.Method == http.MethodPost {
			// Create a new Todo item
		} else {
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			// Update the particular Todo item
		} else if r.Method == http.MethodDelete {
			// Delete the particular Todo item
		} else {
			http.NotFound(w, r)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
