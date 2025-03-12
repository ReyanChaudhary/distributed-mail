package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes API routes
func SetupRoutes(router *gin.Engine) {
	router.GET("/knockknock", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running!"})
	})
}
