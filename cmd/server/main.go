package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/reyan/distributed-mail/config"
	"github.com/reyan/distributed-mail/internal/api"
)

func main() {
	// Initialize config
	config.Init()

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
