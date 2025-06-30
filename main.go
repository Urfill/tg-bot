package main

import (
	"github.com/gin-gonic/gin"
	"tg-bot/services"
)

func main() {
	r := gin.Default()
	//r.GET("/get", clients.TgBotClient)
	r.GET("/get", services.GetAnswer)
	r.POST("/post", services.PostAnswer)

	r.Run(":8080") // порт в конфиги
}
