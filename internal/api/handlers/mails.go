package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// KnockKnockHandler is a simple test endpoint.
func KnockKnockHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is running on Port  8080 !",
	})
}
