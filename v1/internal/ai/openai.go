package ai

import (
	"context"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

func NewOpenAiClient(key string) *OpenAiClient {
	client := OpenAiClient(openai.NewClient(
		option.WithAPIKey(key),
	))
	return &client
}

type OpenAiClient openai.Client

func (c *OpenAiClient) Review(question string) (string, error) {
	ctx := context.Background()
	res, err := c.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(question)},
		Model: openai.ChatModelGPT5_4Mini,
	})

	if err != nil {
		return "", err
	}

	return res.OutputText(), nil

}

var _ Reviewer = (*OpenAiClient)(nil)
