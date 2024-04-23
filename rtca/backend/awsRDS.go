package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
    "fmt"
    "os"
    "github.com/joho/godotenv"
    "time"
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
    log.Println("Successfully closed connection to AWS RDS PostgreSQL")
}

func (rds *AWS_RDS) addUser(user_id int, username string) {
	_, err := rds.pg.Exec(`
        INSERT INTO Users (user_id, username) VALUES ($1, $2)`, 
        user_id, username,
    )

	if err != nil {
		log.Fatal("Error adding user:", err)
	}

    log.Println("Successfully added User")
}

func (rds *AWS_RDS) addMessage(
    sender_id int,
    contentRaw string, contentText string, contentMorse string,
) (int, time.Time) {

    var message_id int
    var timestamp time.Time
	err := rds.pg.QueryRow(`
        INSERT INTO Messages 
        (sender_id, contentRaw, contentText, contentMorse) 
        VALUES ($1, $2, $3, $4) RETURNING message_id, timestamp`, 
        sender_id, contentRaw, contentText, contentMorse,
    ).Scan(&message_id, &timestamp)

	if err != nil {
		log.Fatal("Failed inserting to Messages table:", err)
	}

    log.Println("Successfully inserted to Message table")
	return message_id, timestamp
}

func (rds *AWS_RDS) createUsersTable(){
	_, err := rds.pg.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
            user_id INT PRIMARY KEY NOT NULL,
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
            message_id SERIAL PRIMARY KEY NOT NULL,
            sender_id INT NOT NULL,
            contentRaw TEXT NOT NULL,
            contentText TEXT NOT NULL,
            contentMorse TEXT NOT NULL,
            timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (sender_id) REFERENCES Users(user_id)
        );
        
    `)
    if err != nil {
        log.Fatal("Error creating Messages table:", err)
    }
    log.Println("Successfully created Messages table")
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