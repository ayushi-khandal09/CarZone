package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println("Wait for the database start up...")
	time.Sleep(5 * time.Second)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error Connecting to the database: %v", err)
	}
	fmt.Println("Succesfully Connectd to the database")
}

func GetDB() *sql.DB {
	if db == nil {
		log.Println("Warning: Database connection is not initialized")
	}
	return db
}

func CloseDB() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		} else {
			fmt.Println("Database connection closed")
		}
	} else {
		fmt.Println("Warning: Attempted to close an uninitialized database connection")
	}
}
