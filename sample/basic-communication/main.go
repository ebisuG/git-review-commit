package main

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	anthropicOption "github.com/anthropics/anthropic-sdk-go/option"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

func main() {
	// checkOpenai()
	checkAnthropic()
}

func checkOpenai() {
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

func checkAnthropic() {

	client := anthropic.NewClient(
		anthropicOption.WithAPIKey("dummy-api-key"), // defaults to os.LookupEnv("ANTHROPIC_API_KEY")
	)
	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock("What is a quaternion?")),
		},
		Model: anthropic.ModelClaudeHaiku5_5,
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", message.Content)
}
