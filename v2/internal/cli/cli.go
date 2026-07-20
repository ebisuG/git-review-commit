package cli

import (
	"fmt"
	"os"
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
