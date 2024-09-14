package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	//_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Open database connection with error handling
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=test sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close() // Close database connection when main exits

	// Set up HTTP handlers
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/create", createUser)

	// Create HTTP server instance
	srv := &http.Server{Addr: ":8080"}

	// Start server in a separate goroutine
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	log.Println("Server is running on port 8080")

	// Set up interrupt channel for graceful shutdown
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, os.Kill)

	// Wait for interrupt signal
	<-interruptChan

	// Perform graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error: %v\n", err)
		return
	}
	log.Println("Server shut down successfully")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Query database directly without goroutines
	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Ensure rows are closed after use

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Scan error: %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "User: %s\n", name)
	}

	// Check for any row-level errors
	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Rows error: %v", err), http.StatusInternalServerError)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// Get username from query parameters
	username := r.URL.Query().Get("name")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Use parameterized query to prevent SQL injection
	_, err := db.Exec("INSERT INTO users (name) VALUES ($1)", username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User %s created successfully", username)
}
