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

	req := openai.CompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 3000,
		Prompt:    question,
		Stream:    true,
	}

	stream, err := cli.CreateCompletionStream(ctx, req)
	if err != nil {
		return
	}
	defer stream.Close()

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
			originTxt := response.Choices[0].Text
			htmlTxt := strings.ReplaceAll(originTxt, "\n", "<br>")
			ctx.Writer.Write([]byte(htmlTxt))
			ctx.Writer.Flush()
		}

	}
}
