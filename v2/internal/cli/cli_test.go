package cli

import (
	"fmt"
	"testing"
)

type InterpretCase struct {
	input  []string
	expect Command
}

var ok1 = InterpretCase{
	input: []string{"git-review", "-m", `test-title:

first line of message body
second line of message body
third line of message body`},
	expect: Command{
		Instruction: "git-review",
		Options: []Option{
			{Flag: "-m", Value: "test-title:"},
			{Flag: "-m", Value: `first line of message body
second line of message body
third line of message body`},
		}},
}

func TestGitReviewInterpreterInterpret(t *testing.T) {
	gitReviewInterpreter := NewGitReviewInterpreter()
	command, err := gitReviewInterpreter.Interpret(ok1.input)
	if err != nil {
		t.Errorf("Failed to interpret : %d", err)
	}
	fmt.Println("command : ", command)

}
