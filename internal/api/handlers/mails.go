package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"distributed-mail/internal/mail"
	"distributed-mail/internal/queue"
	"distributed-mail/internal/storage"
)

// KnockKnockHandler is a simple test endpoint.
func KnockKnockHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is running on Port  8080 !",
	})
}

// SendEmailHandler - Accepts email request and queues it
func SendEmailHandler(c *gin.Context) {
	var emailRequest mail.EmailRequest
	if err := c.ShouldBindJSON(&emailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	emailID, err := queue.ProduceEmailTask(emailRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to queue email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "queued", "email_id": emailID})
}

// GetEmailStatusHandler - Fetches email sending status
func GetEmailStatusHandler(c *gin.Context) {
	emailID := c.Param("email_id")
	status, err := storage.GetEmailStatus(emailID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email_id": emailID, "status": status})
}

// RetryEmailHandler - Retries sending a failed email
func RetryEmailHandler(c *gin.Context) {
	emailID := c.Param("email_id")
	err := queue.RetryEmail(emailID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retry email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "requeued", "email_id": emailID})
}