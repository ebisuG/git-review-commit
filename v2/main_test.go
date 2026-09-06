package main

import (
	"strings"
	"testing"

	"github.com/ebisuG/git-review-commit-v2/internal/review"
)

var gitStyleInput1 = strings.Split(`git-review -m "test-title" -m "test-body"`, " ")
var gitStyleInput2 = strings.Split(`git-review -m "test-title"`, " ")

func TestMain(t *testing.T) {
	app := NewMockApp(gitStyleInput1)
	isSuccessed := app.Run()
	if isSuccessed != nil {
		t.Errorf("App failed running")
	}
}

func TestMain2(t *testing.T) {
	app := NewMockApp(gitStyleInput2)
	isSuccessed := app.Run()
	if isSuccessed != nil {
		t.Errorf("App failed running")
	}
}

func NewMockApp(input []string) *App {
	return &App{
		reviewer:      NewStubReviewer(),
		promptBuilder: BuildPrompt,
		interpreter:   NewInterpreter(),
		validater:     NewValidater(),
		input:         input,
	}
}

type StubReviewer struct{}

func (c *StubReviewer) Review(question string) (string, error) {
	return "Dummy Review Message", nil
}

var _ review.Reviewer = (*StubReviewer)(nil)

func NewStubReviewer() *StubReviewer {
	return &StubReviewer{}
}
