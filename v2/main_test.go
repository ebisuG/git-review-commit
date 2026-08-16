package main

import (
	"testing"

	"github.com/ebisuG/git-review-commit-v2/internal/review"
)

func TestMain(t *testing.T) {
	app := NewMockApp()
	isSuccessed := app.Run()
	if isSuccessed != nil {
		t.Errorf("App failed running")
	}
}

func NewMockApp() *App {
	return &App{reviewer: NewStubReviewer(), prompt: BuildPrompt()}
}

type StubReviewer struct{}

func (c *StubReviewer) Review(question string) (string, error) {
	return "Dummy Review Message", nil
}

var _ review.Reviewer = (*StubReviewer)(nil)

func NewStubReviewer() *StubReviewer {
	return &StubReviewer{}
}
