package main

import (
	"fmt"

	"example.com/hw/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := database.InitConnection()
	defer db.Close()

	rows := database.Query(db)
	fmt.Println(rows)
}
