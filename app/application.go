package app

import (
	"github.com/ghifar/bookstore-users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	logger.Info("Zap Logger Configured")
	mapUrls()
	router.Run(":8080")
}
