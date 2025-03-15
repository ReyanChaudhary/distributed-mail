package api

import (

	"github.com/gin-gonic/gin"
	"distributed-mail/internal/api/handlers"
	"distributed-mail/internal/security"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/knockknock", handlers.KnockKnockHandler)
	// // Email APIs
	// router.POST("/send", handlers.SendEmailHandler)
	// router.GET("/status/:email_id", handlers.GetEmailStatusHandler)
	// router.POST("/queue/retry/:email_id", handlers.RetryEmailHandler)

		// Apply authentication middleware to protected routes
		authGroup := router.Group("/")
		authGroup.Use(security.AuthMiddleware())
		{
			authGroup.POST("/send", handlers.SendEmailHandler)
			authGroup.GET("/status/:email_id", handlers.GetEmailStatusHandler)
			authGroup.POST("/queue/retry/:email_id", handlers.RetryEmailHandler)
		}
}

