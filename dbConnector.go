package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "db.sqlite3"

// GetDBConnection returns a connection to the SQLite database.
func _GetDBConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	// Check if the database connection is alive
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}
	return db, nil
}

func dbInitialization() {
	db, err := _GetDBConnection()

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
}

func getAllDbItems() {
	db, err := _GetDBConnection()

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

func insertBuild() {
	db, err := _GetDBConnection()
	
	// Insert data into the table
	insertDataQuery := "INSERT INTO users (name, age) VALUES (?, ?);"
	_, err = db.Exec(insertDataQuery, "John Doe", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted successfully")
}
