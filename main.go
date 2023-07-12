package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTable(db *sql.DB) {
	// Create table
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}

func insertData(db *sql.DB) {
	// Insert data
	_, err := db.Exec("INSERT INTO users (name) VALUES (?)", "John Doe")
	if err != nil {
		log.Fatal(err)
	}
}

func queryData(db *sql.DB) *sql.Rows {
	// Query data
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func main() {
	// Open a connection to SQLite database
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows := queryData(db)
	defer rows.Close()
	// Process query results
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
