package storage

import (
	"errors"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

// DB is the global database connection
var DB *gorm.DB

// EmailRecord represents the email record in the database
type EmailRecord struct {
	ID     string `gorm:"primaryKey"`
	Status string
}

// ConnectDB initializes the database connection
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

	// AutoMigrate creates the table if it does not exist
	err = DB.AutoMigrate(&EmailRecord{})

    if err != nil {
        log.Fatal("Failed to migrate database : ", err) 
    }
    
	log.Println("Database connected successfully")
}

// SaveEmail stores the email record in the database
func SaveEmail(emailID string, status string) error {
	email := EmailRecord{ID: emailID, Status: status}
	result := DB.Create(&email)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetEmailStatus retrieves the email status from the database
func GetEmailStatus(emailID string) (string, error) {
	var email EmailRecord
	result := DB.First(&email, "id = ?", emailID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", errors.New("email not found")
		}
		return "", result.Error
	}
	return email.Status, nil
}
