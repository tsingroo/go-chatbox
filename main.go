package main

import (
	gpt3controller "go-chatgpt/gpt3/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/stream", gpt3controller.HandleStreamCompletion)

	engine.Run(":8080")
}
