package srv

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

func StreamCompletion(question string, ctx *gin.Context) {
	cli := openai.NewClient("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")

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 3000,
		Stream:    true,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Hello!",
			},
		},
	}

	stream, err := cli.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Println("srv StreamCompletion error", err)
		ctx.Writer.Write([]byte(err.Error()))
		return
	}
	defer stream.Close()
	respText := ""

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			fmt.Println("srv StreamCompletion error", err)
			return
		}

		if len(response.Choices) > 0 {
			respText = respText + response.Choices[0].Delta.Content
			htmlTxt := strings.ReplaceAll(respText, "\n", "<br>")
			ctx.Writer.Write([]byte(htmlTxt))
			ctx.Writer.Flush()
		}

	}
}
