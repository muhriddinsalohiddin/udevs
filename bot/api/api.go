package api

import (
	"github.com/gin-gonic/gin"
	_"github.com/muhriddinsalohiddin/udevs/bot/api/docs"
	"github.com/muhriddinsalohiddin/udevs/bot/api/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Message Sender Bot
// @version 1.0
// @description Telegram Bot which sends messages to channels and groups
// @BasePath /
func Option() *gin.Engine {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/send", handlers.SendMessageAPI)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
