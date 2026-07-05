package main

import (
	"fmt"

	"github.com/ebisuG/git-review-commit-v1/internal/ai"
	"github.com/ebisuG/git-review-commit-v1/internal/cli"
)

func main() {
	prompt, err := cli.PromptBuild()
	if err != nil {
		fmt.Println(err)
	}
	client := NewClient()
	comment, err := client.Review(prompt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(comment)
}

func NewClient() ai.Reviewer {
	client := ai.NewOpenAiClient("dummy-key")
	return client
}
