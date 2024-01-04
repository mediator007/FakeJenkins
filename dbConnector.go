package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
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
		CREATE TABLE IF NOT EXISTS builds (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			creationTime DATETIME DEFAULT CURRENT_TIMESTAMP,
			executionTime INTEGER
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
	rows, err := db.Query("SELECT id, creationTime, executionTime FROM builds;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Query results:")
	for rows.Next() {
		var id, executionTime int
		var creationTime time.Time
		err := rows.Scan(&id, &creationTime, &executionTime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, creationTime: %s, executionTime: %d\n", id, creationTime, executionTime)
	}

	// Handle errors from iterating over rows (if any)
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func insertBuild(execTime int) (int64, error){
	db, err := _GetDBConnection()
	// Insert data into the table
	insertDataQuery := "INSERT INTO builds (executionTime) VALUES (?);"
	result, err := db.Exec(insertDataQuery, execTime)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println("Build inserted successfully with id - ", lastInsertID)
	return lastInsertID, nil
}
