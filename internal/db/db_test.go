package db_test

import (
	"student-feedback/internal/db"
	"testing"
)

func TestHealthCheckDB(t *testing.T) {
	db.InitDB()

	err := db.HealthCheckDB()
	if err != nil {
		t.Errorf("Database health check failed: %v", err)
	}
}
