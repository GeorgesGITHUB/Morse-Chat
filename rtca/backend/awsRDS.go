package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

type AWS_RDS struct {
    pg *sql.DB
}

const (
    masterUsername = "postgres"
    endpoint = "pg-db-1.c9mk0csuwqro.us-east-2.rds.amazonaws.com"
    port = "5432"
    databaseName = "postgres"
    driverName = "postgres"
)

//var users = make(map[int]string)

func (rds *AWS_RDS) openConnection(){
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file:", err)
    }

    connectionString := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=disable",
        masterUsername,
        os.Getenv("MASTER_PASSWORD"),
        endpoint,
        port,
        databaseName,
    )

    // Open a connection to the database
    rds.pg, err = sql.Open(driverName, connectionString)
    if err != nil {
        log.Println("Error opening db")
        log.Fatal(err)
    }

    // Ping the database to ensure the connection is successful
    err = rds.pg.Ping()
    if err != nil {
        log.Println("Error pinging db")
        log.Fatal(err)
    }

    log.Println("Successfully connected to AWS RDS PostgreSQL")
}

// Defer call closeConnection() after openConnection()
func (rds *AWS_RDS) closeConnection(){
    rds.pg.Close()
}

func (rds *AWS_RDS) createUsersTable(){
	_, err := rds.pg.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
            user_id SERIAL PRIMARY KEY,
            username VARCHAR(255) NOT NULL
        )
    `)
    if err != nil {
        log.Fatal("Error creating Users table:", err)
    }
    log.Println("Successfully created Users table")
}

func (rds *AWS_RDS) createMessagesTable(){
    _, err := rds.pg.Exec(`
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
    log.Println("Successfully created Messages table")
}

func (rds *AWS_RDS) addUser(username string) int {
	var userID int
	err := rds.pg.QueryRow(`
        INSERT INTO Users (username) 
        VALUES ($1) RETURNING user_id`, 
        username,
    ).Scan(&userID)

	if err != nil {
		log.Fatal("Error adding user:", err)
	}

    log.Printf("Successfully added User with user_id = %d", userID)
	return userID
}

func (rds *AWS_RDS) addMessage( senderID int, 
    contentRaw string, contentText string, contentMorse string) int {

	var message_id int
	err := rds.pg.QueryRow(`
        INSERT INTO Messages (sender_id, contentRaw, contentText, contentMorse) 
        VALUES ($1, $2, $3, $4) RETURNING message_id`, 
        senderID, contentRaw, contentText, contentMorse,
    ).Scan(&message_id)

	if err != nil {
		log.Fatal("Error adding message:", err)
	}

    log.Printf("Successfully added Message with message_id = %d", message_id)
	return message_id
}

// For testing
func (rds *AWS_RDS) deleteAllTables(){
    _, err := rds.pg.Exec(`
        DROP TABLE public.messages;
        DROP TABLE public.users;
    `)
    if err != nil {
        log.Fatal("Error creating Messages table:", err)
    }
    log.Println("Successfully deleted all tables")
}

func (rds *AWS_RDS) getAllUserMessages(user_id int) []Message{
    return nil
}

func (rds *AWS_RDS) getAllUserIDs() []int{
    return nil
}