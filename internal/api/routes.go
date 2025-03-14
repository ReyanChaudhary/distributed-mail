package api

import (

	"github.com/gin-gonic/gin"
	"distributed-mail/internal/api/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/knockknock", handlers.KnockKnockHandler)
	// Email APIs
	router.POST("/send", handlers.SendEmailHandler)
	router.GET("/status/:email_id", handlers.GetEmailStatusHandler)
	router.POST("/queue/retry/:email_id", handlers.RetryEmailHandler)
}

