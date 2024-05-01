package main

import (
	"log"
	"fmt"
	"os"
)

func connectionString(config string) string {
	switch config {
	case "aws_rds":
		log.Println("Creating connection string for",config)
		//DSN URL-style format
		connectionString := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("AWS_RDS_POSTGRES_MASTER_USERNAME"),
			os.Getenv("AWS_RDS_POSTGRES_MASTER_PASSWORD"),
			os.Getenv("AWS_RDS_POSTGRES_ENDPOINT"),
			os.Getenv("AWS_RDS_POSTGRES_PORT"),
			os.Getenv("AWS_RDS_POSTGRES_DATABASE_NAME"),
		)
		return connectionString

	case "dockered_postgres":
		log.Println("Creating connection string for",config)
		connectionString := fmt.Sprintf(
			"user=%s password=%s port=%s dbname=%s host=postgres sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
		)
		return connectionString
	
	default:
		log.Fatal("Selected config for connectionString doesn't exist")
		return ""
	}
}