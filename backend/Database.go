package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
    "time"
    "fmt"
)

type Database struct {
    postgres *sql.DB
}

const (
    driverName = "postgres"
)

func (db *Database) OpenConnection(){
    // Open a connection to the database
    log.Println("DEBUG::",connectionString("dockered_postgres"))
    maxAttempts := 10
	initialDelay := 1 * time.Second
	maxDelay := 30 * time.Second
    err := exponentialBackoff(maxAttempts, initialDelay, maxDelay, func() error {
        var err error
        db.postgres, err = sql.Open(
            driverName, connectionString("dockered_postgres"),
        )
        if err != nil {
            log.Println("Failed opening Database:",err)
            return err
        }
        log.Println("Successfully connected to Database")
        
        // Ping the database to ensure the connection is successful
        err = db.postgres.Ping()
        if err != nil {
            log.Println("Failed pinging Database:",err)
            return err
        }
        log.Println("Successfully pinged the Database")
        return nil
    })

    if err != nil {
        log.Fatal("Failed to connect to PostgreSQL:", err)
    }

}

func exponentialBackoff(maxAttempts int, initialDelay time.Duration, maxDelay time.Duration, operation func() error) error {
    var (
        err          error
        retryAttempt int
        delay        time.Duration
    )

    for retryAttempt = 0; retryAttempt < maxAttempts; retryAttempt++ {
        if err = operation(); err == nil {
            return nil
        }

        // Calculate the next delay using exponential backoff strategy
        delay = initialDelay * time.Duration(1<<uint(retryAttempt))
        if delay > maxDelay {
            delay = maxDelay
        }

        log.Printf("Attempt %d failed. Retrying in %s...\n", retryAttempt+1, delay)
        time.Sleep(delay)
    }

    return fmt.Errorf("exceeded maximum number of attempts (%d)", maxAttempts)
}

// Defer call closeConnection() after openConnection()
func (db *Database) CloseConnection(){
    log.Println("Closing connection to Database")
    db.postgres.Close()
    log.Println("Successfully closed connection to Database")
}

func (db *Database) PostUser(user_id int, username string) {
	_, err := db.postgres.Exec(`
        INSERT INTO Users (user_id, username) VALUES ($1, $2)`, 
        user_id, username,
    )

	if err != nil {
		log.Fatal("Failed adding user:", err)
	}

    log.Println("Successfully added User")
}

func (db *Database) PostMessage(
    sender_id int,
    contentRaw string, contentText string, contentMorse string,
) (int, time.Time) {

    var message_id int
    var timestamp time.Time
	err := db.postgres.QueryRow(`
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

func (db *Database) CreateTables(){
    db.createUsersTable()
    db.createMessagesTable()
}

func (db *Database) createUsersTable(){
	_, err := db.postgres.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
            user_id INT PRIMARY KEY NOT NULL,
            username VARCHAR(255) NOT NULL
        )
    `)
    if err != nil {
        log.Fatal("Failed creating Users table:", err)
    }
    log.Println("Successfully created Users table")
}

func (db *Database) createMessagesTable(){
    _, err := db.postgres.Exec(`
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
        log.Fatal("Failed creating Messages table:", err)
    }
    log.Println("Successfully created Messages table")
}

// For testing
func (db *Database) DeleteTables(){
    _, err := db.postgres.Exec(`
        DROP TABLE public.messages;
        DROP TABLE public.users;
    `)
    if err != nil {
        log.Fatal("Failed to delete table:", err)
    }
    log.Println("Successfully deleted all tables")
}