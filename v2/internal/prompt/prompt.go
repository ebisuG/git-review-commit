package prompt

import (
	"bytes"
	"fmt"
	"text/template"
)

type PromptData struct {
	UserInput      UserInput
	BasePrompt     string
	OutputTemplate string
	GitLog         string
	GitDiff        string
}

type UserInput struct {
	UserDraft                string
	UserPreferredReviewStyle string
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

func NewPromptData(userInput UserInput) (PromptData, error) {
	gitLog, err := readGitLog()
	if err != nil {
		return PromptData{}, err
	}
	gitDiff, err := newDiff()
	if err != nil {
		fmt.Println("Cannot Get Git Diff")
		fmt.Println(err)
	}

	//read user preffered prompt and add it to PromptData
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
