package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

func testRDS(){
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file:", err)
    }

    connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        "postgres", // Master username
        os.Getenv("MASTER_PASSWORD"), // Master password
        "pg-db-1.c9mk0csuwqro.us-east-2.rds.amazonaws.com", // Host Endpoint
        "5432", // Port Number
        "postgres", // Database Name
    )

    // Open a connection to the database
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        log.Println("Error opening db")
        log.Fatal(err)
    }

    defer db.Close()

    // Ping the database to ensure the connection is successful
    err = db.Ping()
    if err != nil {
        log.Println("Error pinging db")
        log.Fatal(err)
    }

    log.Println("successfully completed pgAWS.mainRoutine")
}