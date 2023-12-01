package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

type TimeResponse struct {
	TorontoTime string `json:"toronto_time"`
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8080", nil)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	torontoTime := getCurrentTorontoTime()
	saveTimeToDatabase(torontoTime)

	response := TimeResponse{TorontoTime: torontoTime.Format(time.RFC3339)}
	json.NewEncoder(w).Encode(response)
}

func getCurrentTorontoTime() time.Time {
	loc, _ := time.LoadLocation("America/Toronto")
	return time.Now().In(loc)
}

func saveTimeToDatabase(time time.Time) {
	// Replace the placeholders with your SQL Server details
	server := "NARVEER-SAHARAN\\SQLEXPRESS"
	port := 1433
	user := "sa"
	password := "narveer"
	database := "chandigarhTime"

	// Build the connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	// Open a connection to the database
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// _, err = db.Exec("INSERT INTO time_table (time) VALUES (?)", time)
	// if err != nil {
	// 	panic(err)
	// }
	// Check the connection
	_, err = db.Exec("INSERT INTO time_table (time) VALUES (?)", time)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the SQL Server database.")
}
