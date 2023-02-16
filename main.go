package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	// 判断是否为windows系统
	if runtime.GOOS == "windows" {
		// windows下先执行chcp 65001命令，否则中文会乱码
		execChcpCommand()
	}

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
		MaxTokens: 3000,
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

// execChcpCommand 执行chcp 65001命令
func execChcpCommand() error {
	cmd := exec.Command("chcp", "65001")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
