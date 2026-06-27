package cli

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type Args struct {
	Title string
	Body  string
}

func ValidateArgs() error {
	const argsInMMstyle = 5
	argsLength := len(os.Args)
	if argsLength < 5 {
		return fmt.Errorf("less argument %d", argsLength)
	} else if argsLength > 5 {
		return fmt.Errorf("too many argument %d", argsLength)
	} else {
		return nil
	}
}

func Parse() (*Args, error) {
	err := ValidateArgs()
	if err != nil {
		return &Args{}, err
	}
	return &Args{Title: os.Args[2], Body: os.Args[4]}, nil
}

type PromptData struct {
	UserInput      string
	BasePrompt     string
	OutputTemplate string
}

func PromptBuild() (string, error) {

	input, err := template.New("input").Parse(userInput)
	if err != nil {
		return "", err
	}
	var inputBuf bytes.Buffer
	args, err := Parse()
	if err != nil {
		return "", err
	}
	err = input.Execute(&inputBuf, *args)
	if err != nil {
		return "", err
	}
	s := inputBuf.String()

	data := PromptData{
		BasePrompt:     basePrompt,
		OutputTemplate: outputTemplate,
		UserInput:      s,
	}

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
	fmt.Println(t)
	return t, nil
}
