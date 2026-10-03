package llm

import (
	"context"

	"github.com/ebisuG/git-review-commit-v2/internal/config"
	"github.com/ebisuG/git-review-commit-v2/internal/review"
	anyllm "github.com/mozilla-ai/any-llm-go"
	"github.com/mozilla-ai/any-llm-go/providers"
)

func NewReviewer(loader config.Loader) review.Reviewer {
	conf, err := loader.Load()
	if err != nil {
		return nil
	}
	p, model, err := Resolve(conf.ProviderAndModel)
	return &LlmClient{params: anyllm.CompletionParams{Model: model}, provider: p}
}

type LlmClient struct {
	params   anyllm.CompletionParams
	provider providers.Provider
}

var _ review.Reviewer = (*LlmClient)(nil)

func (l *LlmClient) Review(question string) (string, error) {
	ctx := context.Background()

	messages := []anyllm.Message{
		{Role: anyllm.RoleUser, Content: question},
	}

	l.params.Messages = messages

	response, err := l.provider.Completion(ctx, l.params)

	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.ContentString(), nil

}
