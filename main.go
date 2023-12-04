package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type TimeResponse struct {
	TorontoTime string `json:"toronto_time"`
}

type AllTimesResponse struct {
	AllTimes []string `json:"all_times"`
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/all-times", allTimesHandler)
	http.ListenAndServe(":8585", nil)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	torontoTime := getCurrentTorontoTime()
	saveTimeToDatabase(torontoTime)

	response := TimeResponse{TorontoTime: torontoTime.Format(time.RFC3339)}
	json.NewEncoder(w).Encode(response)
}

func allTimesHandler(w http.ResponseWriter, r *http.Request) {
	allTimes := getAllLoggedTimesFromDatabase()

	response := AllTimesResponse{AllTimes: allTimes}
	json.NewEncoder(w).Encode(response)
}

func getCurrentTorontoTime() time.Time {
	loc, _ := time.LoadLocation("America/Toronto")
	return time.Now().In(loc)
}

func saveTimeToDatabase(time time.Time) {
	db, err := sql.Open("mysql", "root:narveer@/torontotime")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO time_table (time) VALUES (?)", time)
	if err != nil {
		panic(err)
	}
}

func getAllLoggedTimesFromDatabase() []string {
	db, err := sql.Open("mysql", "root:narveer@/torontotime")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT time FROM time_table")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var allTimes []string
	for rows.Next() {
		var timeString string
		if err := rows.Scan(&timeString); err != nil {
			panic(err)
		}
		allTimes = append(allTimes, timeString)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return allTimes
}
