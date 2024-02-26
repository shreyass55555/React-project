package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

// Define a struct for your data (marks and names)
type Student struct {
	ID    int    `json:"id"`
	Marks int    `json:"marks"`
	Name  string `json:"name"`
}

func main() {
	r := mux.NewRouter()

	// Define your API endpoint
	r.HandleFunc("/api/marks", GetMarks).Methods("GET")

	// Add CORS middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

// Handler function to return marks data
func GetMarks(w http.ResponseWriter, r *http.Request) {
	// Connect to Cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1") // Change to your Cassandra cluster IP address
	cluster.Keyspace = "my_keyspace"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Error connecting to Cassandra:", err)
	}
	defer session.Close()

	// Query marks data from Cassandra
	var students []Student
	iter := session.Query("SELECT id, marks, name FROM marks").Iter()
	var id, marksVal int
	var name string
	for iter.Scan(&id, &marksVal, &name) {
		students = append(students, Student{ID: id, Marks: marksVal, Name: name})
	}
	if err := iter.Close(); err != nil {
		log.Fatal("Error querying marks data:", err)
	}

	// Return marks data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
