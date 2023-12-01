package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbPath = "mydatabase.db"

// TimeLog represents the structure of the time_log table
type TimeLog struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"column:timestamp"`
}

// TimeResponse represents the response structure
type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

var db *gorm.DB

func init() {
	// Initialize the database connection with gorm for SQLite3
	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	// Auto migrate the TimeLog struct to the database
	if err := db.AutoMigrate(&TimeLog{}); err != nil {
		log.Fatal("Error auto migrating tables:", err)
	}

	// Check if the database connection is successful
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting DB instance:", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
}

func getTime(w http.ResponseWriter, r *http.Request) {
	// Set the timezone to Toronto
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error loading time zone: %v", err)
		return
	}

	// Get the current time in Toronto
	currentTime := time.Now().In(loc)

	// Save the current time to the database using gorm
	timeLog := TimeLog{Timestamp: currentTime}
	if err := db.Create(&timeLog).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error inserting time into the database: %v", err)
		return
	}

	// Create a response struct
	response := TimeResponse{
		CurrentTime: currentTime.Format("2006-01-02 15:04:05"),
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the struct to JSON and write the response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error encoding JSON response: %v", err)
	}
}

func main() {
	http.HandleFunc("/time", getTime)
	port := ":7575"
	fmt.Printf("Server is running on port %s\n", port)

	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
