package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB loads .env, initializes DB connection and logs errors properly
func InitDB() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("⚠️  Warning: .env file not found, relying on OS environment variables.")
	}
	//connStr := "host=localhost port=5432 user=postgres password=admin dbname=feedback_db sslmode=disable"

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("DB ping failed: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database.")
}

// HealthCheckDB checks if DB is reachable (for testing or liveness)
func HealthCheckDB() error {
	if DB == nil {
		return fmt.Errorf("DB connection is nil")
	}
	return DB.Ping()
}
