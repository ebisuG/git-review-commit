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

func Parse() (*Args, error) {
	if len(os.Args) != 5 {
		return &Args{}, fmt.Errorf("invalid argument format")
	}
	return &Args{Title: os.Args[2], Body: os.Args[4]}, nil
}

type PromptData struct {
	UserInput      string
	BasePrompt     string
	OutputTemplate string
}

// For now use string, use template package when sending the request
func PromptBuild() (string, error) {
	const userInput = `git commit -m "{{.Title}}" -m "{{.Body}}"`

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

	const basePrompt = "You are experienced software engineer. Review below git commit message draft from non native English speakers. In order to grow up strong point, keep the orignal message as much as possible. Reflect user's original structure of message as you can, too. Populate reviewed version into next template for output."
	const outputTemplate = " - Sounds native English speaker version : \n - More precise and concise version : \n - Guessed missing context version : "

	data := PromptData{
		BasePrompt:     basePrompt,
		OutputTemplate: outputTemplate,
		UserInput:      s,
	}

	const prompTemplate = "Main instruction:\n  {{.BasePrompt}}\nUser input:\n{{.UserInput}}\nOutput templates:\n{{.OutputTemplate}}"
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
