package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "user-service/models"
)

var DB *gorm.DB

func InitDB() {
    dsn := "host=postgres user=postgres password=password dbname=users port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB.AutoMigrate(&models.User{})
}
