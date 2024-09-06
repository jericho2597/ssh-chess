package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	config "ssh-server/config"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", config.DbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// Verify the connection
	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database")

	// Create a simple table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS test_table (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create test table: %v", err)
	}

	return nil
}

// RunSampleQuery executes a simple query to test the database connection
func RunSampleQuery() error {
	// Insert a sample record
	_, err := db.Exec("INSERT INTO test_table (name) VALUES (?)", "Sample Name")
	if err != nil {
		return fmt.Errorf("failed to insert sample data: %v", err)
	}

	// Query the inserted record
	var name string
	err = db.QueryRow("SELECT name FROM test_table ORDER BY id DESC LIMIT 1").Scan(&name)
	if err != nil {
		return fmt.Errorf("failed to query sample data: %v", err)
	}

	log.Printf("Successfully queried sample data: %s", name)
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
