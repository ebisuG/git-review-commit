package prompt

import (
	"bytes"
	"text/template"

	"github.com/ebisuG/git-review-commit-v2/internal/cli"
)

type PromptData struct {
	UserInput      string
	BasePrompt     string
	OutputTemplate string
	GitLog         string
}

func newUserInput() (string, error) {
	const userInput = cli.UserInput
	input, err := template.New("input").Parse(userInput)
	if err != nil {
		return "", err
	}
	args, err := cli.Parse()
	if err != nil {
		return "", err
	}
	var inputBuf bytes.Buffer
	err = input.Execute(&inputBuf, *args)
	if err != nil {
		return "", err
	}
	s := inputBuf.String()
	return s, nil
}

func NewPromptData() (PromptData, error) {
	userInput, err := newUserInput()
	if err != nil {
		return PromptData{}, err
	}
	gitLog, err := readGitLog()
	if err != nil {
		return PromptData{}, err
	}
	return PromptData{UserInput: userInput, BasePrompt: basePrompt, OutputTemplate: outputTemplate, GitLog: gitLog}, nil
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
