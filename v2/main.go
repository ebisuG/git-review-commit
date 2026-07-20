package main

import (
	"fmt"

	"github.com/ebisuG/git-review-commit-v2/internal/config"
	"github.com/ebisuG/git-review-commit-v2/internal/prompt"
	"github.com/ebisuG/git-review-commit-v2/internal/review"
)

func NewReviewer(key string) review.Reviewer {
	client := review.NewOpenAiClient(key)
	return client
}

func NewLoader() config.Loader {
	loader := config.NewYamlLoader("config.yaml")
	return loader
}

func BuildPrompt() string {
	promptData, err := prompt.NewPromptData()
	if err != nil {
		fmt.Println(err)
	}
	prompt, err := prompt.NewPrompt(promptData)
	if err != nil {
		fmt.Println(err)
	}
	return prompt
}

func main() {
	config, err := NewLoader().Load()
	if err != nil {
		fmt.Println(err)
	}
	reviewr := NewReviewer(config.ApiKey)
	prompt := BuildPrompt()
	result, err := reviewr.Review(prompt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
