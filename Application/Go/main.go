package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type UserStats struct {
	ProductID int `json:"product_id"`

	TimeTaken int64 `json:"time_taken"`
}

var session *gocql.Session

func init() {
	// Connect to ScyllaDB
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "my_keyspace"
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Error connecting to ScyllaDB:", err)
	}
	log.Println("Connected to ScyllaDB successfully")
}

func getUserStatsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the query parameter
	userID := mux.Vars(r)["userId"]
	var productID int
	var timeTaken int64

	// Query user stats by user ID
	if err := session.Query("SELECT product_id, time_taken FROM user_stats WHERE user_id = ? ALLOW FILTERING", userID).Scan(&productID, &timeTaken); err != nil {
		log.Println("Error fetching user stats:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Create a UserStats object
	userStats := UserStats{
		ProductID: productID,
		TimeTaken: timeTaken,
	}

	// Marshal the UserStats object to JSON
	jsonResponse, err := json.Marshal(userStats)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonResponse)
}

func main() {
	defer session.Close()

	// Create a new router
	r := mux.NewRouter()

	// Register handler for the /userStats endpoint
	r.HandleFunc("/userStats/{userId}", getUserStatsHandler).Methods("GET")

	// Enable CORS middleware
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3001"}), // Update with your React app's URL
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// Start the HTTP server with CORS middleware
	log.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", cors(r)))
}
