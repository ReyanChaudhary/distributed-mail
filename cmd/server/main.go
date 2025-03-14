package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"distributed-mail/config"
	"distributed-mail/internal/api"
	"distributed-mail/internal/storage"
	"distributed-mail/internal/queue"
)

func main() {
	// Initialize config
	config.Init()
	storage.ConnectDB()

	err := storage.DB.AutoMigrate(&storage.Email{})
	if err != nil {
		log.Fatal("Unable to migrate the DB : ", err)
	}

	// Initialize Kafka Producer
	err = queue.InitKafkaWriter("localhost:9092", "email-queue")
		if err != nil {
			log.Fatal("Kafka initialization failed:", err)
		}
	
	// Create a new Gin router
	r := gin.Default()

	// Middleware to log API requests
	r.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Setup routes from the api package
	api.SetupRoutes(r)

	log.Println("Server running on port 8080")
	log.Fatal(r.Run(":8080"))
}
