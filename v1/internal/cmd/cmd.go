package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

type Args struct {
	Title string
	Body  string
}

func Parse() *Args {
	return &Args{os.Args[2], os.Args[4]}
}

//TODO:validate user's input format
// func Validate() {
// }

type PromptData struct {
	UserInput      string
	BasePrompt     string
	OutputTemplate string
}

// For now use string, use template package when sending the request
func PromptBuild() {
	const userInput = `git commit -m "{{.Title}}" -m "{{.Body}}"`
	// const userInput = " - Commit title:{{.Title}}\n - Commit message body:{{.Body}}"

	input, err := template.New("input").Parse(userInput)
	if err != nil {
		panic(err)
	}
	var inputBuf bytes.Buffer
	args := Parse()
	err = input.Execute(&inputBuf, *args)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	var promptBuf bytes.Buffer
	err = prompt.Execute(&promptBuf, data)
	if err != nil {
		panic(err)
	}
	t := promptBuf.String()
	fmt.Println(t)

}
