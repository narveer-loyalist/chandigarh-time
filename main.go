package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
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
	//saveTimeToDatabase(torontoTime)

	response := TimeResponse{TorontoTime: torontoTime.Format(time.RFC3339)}
	json.NewEncoder(w).Encode(response)
}

func getCurrentTorontoTime() time.Time {
	loc, _ := time.LoadLocation("America/Toronto")
	return time.Now().In(loc)
}

func saveTimeToDatabase(time time.Time) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS time_table (id INTEGER PRIMARY KEY AUTOINCREMENT, time DATETIME)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO time_table (time) VALUES (?)", time)
	if err != nil {
		panic(err)
	}
}
