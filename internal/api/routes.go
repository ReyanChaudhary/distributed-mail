package api

import (

	"github.com/gin-gonic/gin"
	"github.com/reyan/distributed-mail/internal/api/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/knockknock", handlers.KnockKnockHandler)
}
