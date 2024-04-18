package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
    "fmt"
    "os"
    "github.com/joho/godotenv"
)
// will be a struct with member bar + functions, like messageType.go

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

    // Create Users table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
            user_id SERIAL PRIMARY KEY,
            username VARCHAR(255) NOT NULL
        )
    `)
    if err != nil {
        log.Fatal("Error creating Users table:", err)
    }
    fmt.Println("Users table created successfully")

    // Create Messages table
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Messages (
            message_id SERIAL PRIMARY KEY,
            sender_id INT NOT NULL,
            contentRaw TEXT,
            contentText TEXT,
            contentMorse TEXT,
            timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (sender_id) REFERENCES Users(user_id)
        )
    `)
    if err != nil {
        log.Fatal("Error creating Messages table:", err)
    }
    fmt.Println("Messages table created successfully")

    userID := addUser(db, "Geovani")
    log.Println("Sucessfully","added user")
    log.Println("user_id is",userID)

    msgID := addMessage(db, userID, "some raw", "some text", "some Morse")
    log.Println("Sucessfully","added message")
    log.Println("msg_id is",msgID)

    msgID = addMessage(db, userID, "some raw 2", "some text 2", "some Morse 2")
    log.Println("Sucessfully","added message")
    log.Println("msg_id is",msgID)

    userID = addUser(db, "Jeff")
    log.Println("Sucessfully","added user")
    log.Println("user_id is",userID)

    msgID = addMessage(db, userID, "some raw", "some text", "some Morse")
    log.Println("Sucessfully","added message")
    log.Println("msg_id is",msgID)

    msgID = addMessage(db, userID, "some raw 2", "some text 2", "some Morse 2")
    log.Println("Sucessfully","added message")
    log.Println("msg_id is",msgID)


}

func addUser(db *sql.DB, username string) int {
	var userID int
	err := db.QueryRow(`
        INSERT INTO Users (username) VALUES ($1) RETURNING user_id`, 
        username,
    ).Scan(&userID)

	if err != nil {
		log.Fatal("Error adding user:", err)
	}
	return userID
}

func addMessage(db *sql.DB, senderID int, contentRaw string, contentText string, contentMorse string) int {
	var messageID int
	err := db.QueryRow(`
        INSERT INTO Messages (sender_id, contentRaw, contentText, contentMorse) VALUES ($1, $2, $3, $4) RETURNING message_id`, 
        senderID, contentRaw, contentText, contentMorse,
    ).Scan(&messageID)
	if err != nil {
		log.Fatal("Error adding message:", err)
	}
	return messageID
}