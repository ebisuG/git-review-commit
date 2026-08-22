package prompt

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/ebisuG/git-review-commit-v2/internal/cli"
)

type PromptData struct {
	UserInput      string
	BasePrompt     string
	OutputTemplate string
	GitLog         string
	GitDiff        string
}

func NewUserInput(command cli.Command) (string, error) {
	const ct = cli.CommandTemplate
	tmpl, err := template.New("input").Parse(ct)
	if err != nil {
		return "", err
	}
	var inputBuf bytes.Buffer
	err = tmpl.Execute(&inputBuf, command)
	if err != nil {
		return "", err
	}
	s := inputBuf.String()
	return s, nil
}

func newDiff() (string, error) {
	diff, err := buildDiff()
	if err != nil {
		fmt.Println("Cannot Find Git Diff")
		return "", nil
	}

	templater, err := template.New("diff").Parse(diffTemplate)
	var diffBuf bytes.Buffer
	err = templater.Execute(&diffBuf, diff)
	if err != nil {
		fmt.Println("Cannot Output Git Diff")
		return "", nil
	}

	return diffBuf.String(), nil
}

func NewPromptData(userInput string) (PromptData, error) {
	gitLog, err := readGitLog()
	if err != nil {
		return PromptData{}, err
	}
	gitDiff, err := newDiff()
	if err != nil {
		fmt.Println("Cannot Get Git Diff")
		fmt.Println(err)
	}
	return PromptData{UserInput: userInput, BasePrompt: basePrompt, OutputTemplate: outputTemplate, GitLog: gitLog, GitDiff: gitDiff}, nil
}

func NewPrompt(data PromptData) (string, error) {
	prompt, err := template.New("prompt").Parse(prompTemplate)
	if err != nil {
		return "", err
	}
	var promptBuf bytes.Buffer
	err = prompt.Execute(&promptBuf, data)
	if err != nil {
		return "", err
	}
	t := promptBuf.String()
	return t, nil
}
