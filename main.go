package main

import (
	"tg-bot/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/get", services.GetAnsw)

	r.POST("/post", services.PostAnsw)

	r.Run(":8080")
}
