package gpt3

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-chatgpt/gpt3/srv"
)

// HandleStreamCompletion 流式输出GPT回复
func HandleStreamCompletion(ctx *gin.Context) {
	message := ctx.Query("message")
	if message == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "请输入message参数",
		})
		return
	}
	ctx.Header("Content-Type", "text/event-stream; charset=utf-8")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Transfer-Encoding", "chunked")

	srv.StreamCompletion(message, ctx)
}
