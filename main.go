package main

import (
	"tg-bot/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", services.GetAnswer)
	r.POST("/post", services.PostAnswer)

	r.Run(":8080") // порт в конфиги
}
