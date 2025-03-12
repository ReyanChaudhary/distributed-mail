package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/reyan/distributed-mail/config"
)

func main() {
	// Initialize config
	config.Init()

	// Start Gin server
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Server is running!")
	})

	log.Println("Server running on port 8080")
	log.Fatal(r.Run(":8080"))
}
