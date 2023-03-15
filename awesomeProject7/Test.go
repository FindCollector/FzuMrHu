package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-openai"
)

func main() {
	r := gin.Default()
	r.GET("/chat", handleRequest)
	r.Run(":8080")
}

func handleRequest(c *gin.Context) {
	comment, hascmt := c.GetQuery("comment")
	key, haskey := c.GetQuery("key")

	if !hascmt || !haskey {
		return
	}

	clnt := gogpt.NewClient(key)
	ctx := context.Background()

	req := gogpt.ChatCompletionRequest{
		Model:     gogpt.GPT3Dot5Turbo,
		MaxTokens: 100,
		Messages: []gogpt.ChatCompletionMessage{
			{Role: "user", Content: comment},
		},
	}

	resp, err := clnt.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Println("Error:" + err.Error())
		return
	}

	c.String(200, resp.Choices[0].Message.Content)
}
