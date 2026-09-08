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

first line of message body.
second line of message body.
third line of message body.`},
	expect: Command{
		Instruction: "git-review",
		Options: []Option{
			{Flag: "-m", Value: "test-title:"},
			{Flag: "-m", Value: `first line of message body. second line of message body. third line of message body. `},
		}},
}

func TestGitReviewInterpreterInterpret(t *testing.T) {
	gitReviewInterpreter := NewGitReviewInterpreter()
	command, err := gitReviewInterpreter.Interpret(ok1.input)
	if err != nil {
		t.Errorf("Failed to interpret : %d", err)
	}
	if command.Instruction != ok1.expect.Instruction {
		t.Errorf("Failed to parse main command.")
	}
	if command.Options[0].Flag != ok1.expect.Options[0].Flag ||
		command.Options[0].Value != ok1.expect.Options[0].Value ||
		command.Options[1].Flag != ok1.expect.Options[1].Flag ||
		command.Options[1].Value != ok1.expect.Options[1].Value {

		t.Errorf("Failed to parse options.")
	}
	fmt.Println("command : ", command)

}
