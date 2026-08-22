package main

import (
	"fmt"

	"github.com/ebisuG/git-review-commit-v2/internal/cli"
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

func NewInterpreter() cli.Interpreter {
	interpreter := cli.NewGitReviewInterpreter()
	return &interpreter
}

func NewValidater() cli.Validater {
	validater := cli.NewGitReviewValidater()
	return &validater
}

func BuildPrompt() string {
	input := cli.GetArgs()
	validater := NewValidater()
	err := validater.Validate(input)
	if err != nil {
		fmt.Println(err)
	}

	interpreter := NewInterpreter()
	command, err := interpreter.Interpret(input)
	if err != nil {
		fmt.Println(err)
	}

	userDraft, err := prompt.NewUserInput(command)
	if err != nil {
		fmt.Println(err)
	}

	promptData, err := prompt.NewPromptData(userDraft)
	if err != nil {
		fmt.Println(err)
	}
	prompt, err := prompt.NewPrompt(promptData)
	if err != nil {
		fmt.Println(err)
	}
	return prompt
}

type App struct {
	reviewer review.Reviewer
	prompt   string
}

type Runner interface {
	Run() error
}

func (a *App) Run() error {
	result, err := a.reviewer.Review(a.prompt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result)
	return nil
}

func NewApp(loader config.Loader) *App {
	config, err := loader.Load()
	if err != nil {
		fmt.Println(err)
	}
	return &App{reviewer: NewReviewer(config.ApiKey), prompt: BuildPrompt()}
}

var _ Runner = (*App)(nil)

func main() {
	loader := NewLoader()
	app := NewApp(loader)
	app.Run()
}
