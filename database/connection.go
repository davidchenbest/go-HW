package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitConnection() *sql.DB {
	// Open a connection to SQLite database
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateTable(db *sql.DB) {
	// Create table
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(db *sql.DB) {
	// Insert data
	_, err := db.Exec("INSERT INTO users (name) VALUES (?)", "John Doe")
	if err != nil {
		log.Fatal(err)
	}
}

func Query(db *sql.DB) []string {
	// Query data
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Process query results
	results := []string{}
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		result := fmt.Sprintf("%v, %s", id, name)
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}
