package storage

import (
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
    err := godotenv.Load() // Load .env file
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := "host=" + os.Getenv("DB_HOST") +
           " user=" + os.Getenv("DB_USER") +
           " password=" + os.Getenv("DB_PASSWORD") +
           " dbname=" + os.Getenv("DB_NAME") +
           " port=" + os.Getenv("DB_PORT") +
           " sslmode=" + os.Getenv("DB_SSLMODE")

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("Database connected successfully")
}
