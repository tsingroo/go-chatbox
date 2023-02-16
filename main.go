package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Run(":80")
}

// func handleUserReqs() {
// 	cli := gogpt.NewClient("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")
// 	ctx := context.Background()

// 	req := gogpt.CompletionRequest{
// 		Model:     gogpt.GPT3TextDavinci003,
// 		MaxTokens: 3000,
// 		Prompt:    "This is a test",
// 		Stream:    true,
// 	}

// 	stream, err := cli.CreateCompletionStream(ctx, req)
// 	if err != nil {
// 		return
// 	}
// 	defer stream.Close()

// 	for {
// 		response, err := stream.Recv()
// 		if errors.Is(err, io.EOF) {
// 			fmt.Println("Stream finished")
// 			break
// 		}

// 		if err != nil {
// 			fmt.Printf("Stream error: %v\n", err)
// 			return
// 		}

// 		if len(response.Choices) > 0 {
// 			fmt.Print(response.Choices[0].Text)
// 		}

// 	}
// }
