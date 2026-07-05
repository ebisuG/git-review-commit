package main

import (
	"fmt"

	"github.com/ebisuG/git-review-commit-v1/internal/ai"
	"github.com/ebisuG/git-review-commit-v1/internal/cli"
	"github.com/ebisuG/git-review-commit-v1/internal/config"
)

func main() {
	prompt, err := cli.PromptBuild()
	if err != nil {
		fmt.Println(err)
	}
	config, err := NewLoader().Load()
	if err != nil {
		fmt.Println(err)
	}
	client := NewClient(config.ApiKey)
	comment, err := client.Review(prompt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(comment)
}

func NewClient(key string) ai.Reviewer {
	client := ai.NewOpenAiClient(key)
	return client
}

func NewLoader() config.Loader {
	loader := config.NewYamlLoader("config.yaml")
	return &loader
}
