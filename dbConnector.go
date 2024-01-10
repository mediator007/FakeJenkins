package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "db.sqlite3"

type Build struct {
	ID            int64
	ExecutionTime int64
	StartTime     time.Time
	BuildStatus   string
	JobName       string
}

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

func getAllBuilds() ([]Build, error) {
	db, _ := _GetDBConnection()
	rows, err := db.Query("SELECT * FROM builds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var builds []Build

	for rows.Next() {
		var build Build
		err := rows.Scan(&build.ID, &build.ExecutionTime, &build.StartTime, &build.BuildStatus, &build.JobName)
		if err != nil {
			return nil, err
		}
		builds = append(builds, build)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return builds, nil
}

func dbInitialization() {
	db, err := _GetDBConnection()

	// Create a table (if it doesn't exist)
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS builds (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			executionTime INTEGER,
			startTime DATETIME 
				DEFAULT CURRENT_TIMESTAMP,
			buildStatus CHAR(20) 
				DEFAULT 'INQUEUE' 
				CHECK (buildStatus IN (
					'INPROGRESS', 'INQUEUE', 
					'ABORTED', 'FAILED',
					'SUCCESSFUL'
					)
				),
			jobName CHAR(10) 
				DEFAULT 'DEFAULT' 
				CHECK (jobName IN (
					'DEFAULT', 'GENERATOR'
					)
				)
		);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'users' created successfully")
}

func insertBuild(execTime int) (int64, error) {
	db, err := _GetDBConnection()
	// Insert data into the table
	insertDataQuery := "INSERT INTO builds (executionTime) VALUES (?);"
	result, err := db.Exec(insertDataQuery, execTime)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println("Build inserted successfully with id - ", lastInsertID)
	return lastInsertID, nil
}

func updateBuildStatus(buildNumber string, status string) (int64, error) {
	db, err := _GetDBConnection()
	// Insert data into the table
	updateDataQuery := "UPDATE builds SET buildStatus = ? WHERE id = ?"
	result, err := db.Exec(updateDataQuery, status, buildNumber)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error getting RowsAffected: %v", err)
	}
	fmt.Println("Build inserted successfully with id - ", rowsAffected)
	return rowsAffected, nil
}

func getBuildByBuildNumber(buildNumber string) (Build, error) {
	var build Build
	db, err := _GetDBConnection()
	query := "SELECT * FROM builds WHERE id = ?"
	row := db.QueryRow(query, buildNumber)

	err = row.Scan(
		&build.ID,
		&build.ExecutionTime,
		&build.StartTime,
		&build.BuildStatus,
		&build.JobName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case when no rows are found
			return Build{}, fmt.Errorf("no build found")
		}
		return Build{}, err
	}

	return build, nil
}
