package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	cli := gogpt.NewClient("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")
	ctx := context.Background()
	handleUserReqs(cli, ctx)
}

// 读取用户输入
func readInput() string {
	var input string
	fmt.Print("请输入：")
	fmt.Scanln(&input)
	return input
}

func handleUserReqs(cli *gogpt.Client, ctx context.Context) {
	userInput := readInput()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 10000,
		Prompt:    userInput,
		Stream:    true,
	}
	userInput = ""
	stream, err := cli.CreateCompletionStream(ctx, req)
	if err != nil {
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			break
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return
		}

		if len(response.Choices) > 0 {
			fmt.Print(response.Choices[0].Text)
		}

	}

	if userInput == "exit" {
		return
	} else {
		handleUserReqs(cli, ctx)
	}

}
