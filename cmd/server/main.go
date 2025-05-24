package main

import (
    "log"
    "student-feedback/internal/db"
    "student-feedback/internal/handlers"

    "github.com/gin-gonic/gin"
)

func main() {
    db.InitDB() // Initialize database connection

    r := gin.Default()
    handlers.RegisterRoutes(r)

    log.Println("Server running on http://localhost:8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
