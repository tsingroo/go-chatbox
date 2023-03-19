package srv

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

func StreamCompletion(question string, ctx *gin.Context) {
	// 先获取http代理client
	proxyClient, err := getProxyHttpClient()
	if err != nil {
		fmt.Println("获取代理失败", err)
		ctx.Writer.Write([]byte(err.Error()))
		return
	}

	cliCfg := openai.DefaultConfig("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")
	cliCfg.HTTPClient = proxyClient

	cli := openai.NewClientWithConfig(cliCfg)

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 3000,
		Stream:    true,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	}

	stream, err := cli.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Println("cli CreateChatCompletionStream error", err)
		ctx.Writer.Write([]byte(err.Error()))
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("stream.Recv EOF", err)
			return
		}

		if err != nil {
			fmt.Println("srv StreamCompletion error", err)
			ctx.Writer.Write([]byte(err.Error()))
			return
		}

		if len(response.Choices) > 0 {
			respText := response.Choices[0].Delta.Content
			ctx.Writer.Write([]byte(respText))
			ctx.Writer.Flush()
		}

	}
}

// getProxyHttpClient 获取http代理client
func getProxyHttpClient() (*http.Client, error) {
	proxyUrl, err := url.Parse("http://localhost:20171")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	client := &http.Client{Transport: transport}

	return client, nil
}
