package storage

import (
    "time"
)

// Email represents the email entity in the database
type Email struct {
    ID        string         `gorm:"primaryKey"`
    Recipient string
    Subject   string
    Body      string
    Status    string
    CreatedAt time.Time
    UpdatedAt time.Time
}
