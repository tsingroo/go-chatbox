package srv

import (
	"errors"
	"io"

	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func StreamCompletion(question string, ctx *gin.Context) {
	cli := gogpt.NewClient("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")

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
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
			ctx.Writer.Flush()
			return
		}

		if err != nil {
			// fmt.Printf("Stream error: %v\n", err)
			return
		}

		if len(response.Choices) > 0 {
			ctx.Writer.Write([]byte(response.Choices[0].Text))
		}

	}
}
