package main

import (
	gpt3controller "go-chatgpt/gpt3/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	engine.GET("/stream", gpt3controller.HandleStreamCompletion)

	engine.Run(":8090")
}
