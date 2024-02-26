// main.go

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

var session *gocql.Session

type inputData struct {
	Name   string `json:"name" binding:"required"`
	ID     int    `json:"id" binding:"required"`
	Salary int    `json:"salary" binding:"required"`
}

func main() {
	// Connect to ScyllaDB cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "my_keyspace"
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Initialize Gin router
	router := gin.Default()

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Define API endpoint to insert data
	router.POST("/api/data", insertData)

	// Run the server
	router.Run(":8080")
}

func insertData(c *gin.Context) {
	// Parse JSON request body
	var data inputData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert data into ScyllaDB table
	if err := session.Query("INSERT INTO input (name, id, salary) VALUES (?, ?, ?)",
		data.Name, data.ID, data.Salary).Exec(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})
}
