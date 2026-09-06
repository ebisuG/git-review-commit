package cli

import (
	"errors"
	"fmt"
	"os"
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
	for i := 0; i < len(options)-1; i += 2 {
		option := Option{Flag: input[i], Value: input[i+1]}
		flagAndValues = append(flagAndValues, option)
	}
	return Command{Instruction: instruction, Options: flagAndValues}, nil
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

type GetInput func() []string

func GetArgs() []string {
	return os.Args
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
