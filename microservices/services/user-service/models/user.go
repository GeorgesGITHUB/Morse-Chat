package models

import "gorm.io/gorm"

type User struct {
    ID       string `gorm:"primaryKey"`
    Name     string
    Email    string `gorm:"unique"`
    Password string
}

var _ gorm.Model