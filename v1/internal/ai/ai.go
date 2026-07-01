package ai

import (
	"context"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

type AiClients struct {
	OpenAiClient openai.Client
}

type OpenAiReviewer interface {
	OpenAiReview(question string) (string, error)
}

func NewOpenAiClient(key string) openai.Client {
	client := openai.NewClient(
		option.WithAPIKey(key),
	)
	return client
}

func NewAiClients() *AiClients {
	openAiKey := ""
	return &AiClients{OpenAiClient: NewOpenAiClient(openAiKey)}
}

var _ OpenAiReviewer = (*AiClients)(nil)

func (c *AiClients) OpenAiReview(question string) (string, error) {
	// func (q *Question) Review(teacher openai.Client) (string, error) {
	ctx := context.Background()
	res, err := c.OpenAiClient.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(question)},
		Model: openai.ChatModelGPT5_4Mini,
	})

	if err != nil {
		return "", err
	}

	return res.OutputText(), nil

}
