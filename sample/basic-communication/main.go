package main

import (
	"context"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

func main() {
	ctx := context.Background()
	client := openai.NewClient(
		option.WithAPIKey("dummy-api-key"),
	)

	question := "What is cat in Japanese?"

	res, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(question)},
		Model: openai.ChatModelGPT5_4Mini,
	})

	if err != nil {
		panic(err)
	}

	println(res.OutputText())
}
