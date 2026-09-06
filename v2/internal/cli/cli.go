package cli

import (
	"errors"
	"os"
	"strings"
)

type Args struct {
	Title string
	Body  string
}

type Command struct {
	Instruction string
	Options     []Option
}
type Option struct {
	Flag  string
	Value string
}

type Interpreter interface {
	Interpret(input []string) (Command, error)
}

type GitReviewInterpreter struct{}

var _ Interpreter = (*GitReviewInterpreter)(nil)

func NewGitReviewInterpreter() GitReviewInterpreter {
	return GitReviewInterpreter{}
}

func (g *GitReviewInterpreter) Interpret(input []string) (Command, error) {
	instruction := input[0]
	options := input[1:]
	var flagAndValues []Option

	//case for -m has both title and body
	//ex) git-review -m "title:
	//
	//main body message 1
	//main body message 2"
	if len(options) == 2 {
		strings.ReplaceAll(options[1], "\r\n", "\n")
		if strings.Contains(options[1], "\n") {
			title := strings.Split(options[1], "\n")[0]
			body := strings.Split(options[1], "\n")[1]
			formattedInput := []string{"-m", title, "-m", body}
			flagAndValues = convertStringToOption(formattedInput)
			return Command{Instruction: instruction, Options: flagAndValues}, nil
		}
	}

	//case for -m <title> and -m <body> pattern
	flagAndValues = convertStringToOption(options)
	return Command{Instruction: instruction, Options: flagAndValues}, nil
}

func convertStringToOption(s []string) []Option {
	var flagAndValues []Option
	for i := 0; i < len(s)-1; i += 2 {
		option := Option{Flag: s[i], Value: s[i+1]}
		flagAndValues = append(flagAndValues, option)
	}
	return flagAndValues
}

type Validater interface {
	Validate(input []string) error
}

type GitReviewValidater struct{}

var _ Validater = (*GitReviewValidater)(nil)

func NewGitReviewValidater() GitReviewValidater {
	return GitReviewValidater{}
}

func (v *GitReviewValidater) Validate(input []string) error {
	if len(input) == 3 {
		if input[1] != "-m" {
			return errors.New("first flag is invalid")
		}
		if len(input[2]) == 0 {
			return errors.New("At lease, write git commit message title")
		}
		return nil
	}

	if len(input) == 5 {
		if input[1] != "-m" {
			return errors.New("first flag is invalid")
		}
		if len(input[2]) == 0 {
			return errors.New("At lease, write git commit message title")
		}
		if input[3] != "-m" {
			return errors.New("second flag is invalid")
		}
		return nil
	}
	return errors.New("Follow format : git-review -m <title> -m <body>")
}

func GetArgs() []string {
	return os.Args
}
