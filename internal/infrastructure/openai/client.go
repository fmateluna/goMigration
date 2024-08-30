package openai

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	apiClient *openai.Client
}

func NewOpenAIClient(apiKey string) *Client {
	return &Client{
		apiClient: openai.NewClient(apiKey),
	}
}

func (c *Client) RequestImprovement(ctx context.Context, prompt string) (string, error) {
	messages := []openai.ChatCompletionMessage{
		{Role: "system", Content: "Eres un developer senior."},
		{Role: "user", Content: prompt},
	}

	resp, err := c.apiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT4o20240513,
		Messages: messages,
	})
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
