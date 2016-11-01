package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

// Todo is a representation of each item that
// we are going to maintain in our service
type Todo struct {
	ID        uint64
	Task      string
	Completed bool
}

// In-memory collection of Todos
var todos []Todo

// If we store the Todos in a database, this will be the LastInsertID
var todoCounter uint64

// The single lock is used to synchronize access to
// the above array as well as the counter variables.
var lock sync.RWMutex

var hmacSampleSecret = []byte("Top secret signing key")

// A middleware for validating JWT tokens
func authCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.StandardClaims); ok && token.Valid {
			log.Println(claims)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Authenticate and return JWT
	})

	http.Handle("/todos", authCheck(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Return all the Todo items
			lock.RLock()
			defer lock.RUnlock()

			j, err := json.Marshal(todos)
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(j)

		} else if r.Method == http.MethodPost {
			// Create a new Todo item
			decoder := json.NewDecoder(r.Body)
			var t Todo
			err := decoder.Decode(&t)
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			if len(t.Task) < 1 {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			t.Completed = false

			lock.Lock()
			t.ID = todoCounter
			todoCounter++
			todos = append(todos, t)
			lock.Unlock()
			w.WriteHeader(http.StatusCreated)
		} else {
			http.NotFound(w, r)
		}
	})))

	http.Handle("/todo", authCheck(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			// Update the particular Todo item
		} else if r.Method == http.MethodDelete {
			// Delete the particular Todo item
		} else {
			http.NotFound(w, r)
		}
	})))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
