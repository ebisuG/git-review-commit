package main

import (
	"fmt"

	"github.com/ebisuG/git-review-commit-v2/internal/cli"
	"github.com/ebisuG/git-review-commit-v2/internal/config"
	"github.com/ebisuG/git-review-commit-v2/internal/llm"
	"github.com/ebisuG/git-review-commit-v2/internal/prompt"
	"github.com/ebisuG/git-review-commit-v2/internal/review"
)

func NewReviewer(conf config.Loader) review.Reviewer {
	reviewer, err := llm.NewAnyLlmGoReviewer(conf)
	if err != nil {
		fmt.Errorf("failed to get a reviewer : %v", err)
	}
	return reviewer
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

type PromptBuilder func(command cli.Command) string

func BuildPrompt(command cli.Command) string {

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
	reviewer      review.Reviewer
	promptBuilder PromptBuilder
	interpreter   cli.Interpreter
	validater     cli.Validater
	input         []string
}

type Runner interface {
	Run() error
}

func (a *App) Run() error {
	err := a.validater.Validate(a.input)
	if err != nil {
		fmt.Println(err)
	}

	command, err := a.interpreter.Interpret(a.input)
	if err != nil {
		fmt.Println(err)
	}

	prompt := a.promptBuilder(command)

	result, err := a.reviewer.Review(prompt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result)
	return nil
}

func NewApp(loader config.Loader) *App {

	return &App{
		reviewer:      NewReviewer(loader),
		promptBuilder: BuildPrompt,
		interpreter:   NewInterpreter(),
		validater:     NewValidater(),
		input:         cli.GetArgs(),
	}
}

var _ Runner = (*App)(nil)

func main() {
	loader := NewLoader()
	app := NewApp(loader)
	app.Run()
}
