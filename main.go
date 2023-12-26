package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open the database, creating it if it doesn't exist
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to SQLite database")

	// Create a table (if it doesn't exist)
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER
		);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'users' created successfully")

	// Insert data into the table
	insertDataQuery := "INSERT INTO users (name, age) VALUES (?, ?);"
	_, err = db.Exec(insertDataQuery, "John Doe", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted successfully")

	// Query data from the table
	rows, err := db.Query("SELECT id, name, age FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Query results:")
	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// Handle errors from iterating over rows (if any)
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
