package main

import (
	"bytes"
	"fmt"
	"text/template"

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
	loader := config.NewYamlLoader("config.yaml", "review-style.txt")
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

func NewUserInput(command cli.Command, loader config.Loader) (prompt.UserInput, error) {
	const ct = cli.CommandTemplate
	tmpl, err := template.New("input").Parse(ct)
	if err != nil {
		return prompt.UserInput{}, err
	}
	var inputBuf bytes.Buffer
	err = tmpl.Execute(&inputBuf, command)
	if err != nil {
		return prompt.UserInput{}, err
	}
	s := inputBuf.String()

	conf, err := loader.Load()

	if err != nil {
		return prompt.UserInput{}, err
	}

	var userInput prompt.UserInput
	userInput.UserDraft = s
	userInput.UserDraft = conf.ReviewStyle
	return userInput, nil
}

type PromptBuilder func(command cli.Command, loader config.Loader) string

func BuildPrompt(command cli.Command, loader config.Loader) string {

	userDraft, err := NewUserInput(command, loader)
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
	loader        config.Loader
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

	prompt := a.promptBuilder(command, a.loader)

	result, err := a.reviewer.Review(prompt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result)
	return nil
}

func NewApp() *App {
	loader := NewLoader()

	return &App{
		reviewer:      NewReviewer(loader),
		promptBuilder: BuildPrompt,
		interpreter:   NewInterpreter(),
		validater:     NewValidater(),
		input:         cli.GetArgs(),
		loader:        loader,
	}
}

var _ Runner = (*App)(nil)

func main() {
	app := NewApp()
	app.Run()
}
