package llm

import (
	"context"
	"fmt"
	"strings"

	"github.com/ebisuG/git-review-commit-v2/internal/config"
	"github.com/ebisuG/git-review-commit-v2/internal/review"
	anyllm "github.com/mozilla-ai/any-llm-go"
	"github.com/mozilla-ai/any-llm-go/providers"
	"github.com/mozilla-ai/any-llm-go/providers/anthropic"
	"github.com/mozilla-ai/any-llm-go/providers/deepseek"
	"github.com/mozilla-ai/any-llm-go/providers/gemini"
	"github.com/mozilla-ai/any-llm-go/providers/groq"
	"github.com/mozilla-ai/any-llm-go/providers/llamacpp"
	"github.com/mozilla-ai/any-llm-go/providers/llamafile"
	"github.com/mozilla-ai/any-llm-go/providers/mistral"
	"github.com/mozilla-ai/any-llm-go/providers/ollama"
	"github.com/mozilla-ai/any-llm-go/providers/openai"
	"github.com/mozilla-ai/any-llm-go/providers/zai"
)

type factory func(apiKey string) (providers.Provider, error)

var registry = map[string]factory{
	"openai":    func(apiKey string) (providers.Provider, error) { return openai.New(anyllm.WithAPIKey(apiKey)) },
	"anthropic": func(apiKey string) (providers.Provider, error) { return anthropic.New(anyllm.WithAPIKey(apiKey)) },
	"gemini":    func(apiKey string) (providers.Provider, error) { return gemini.New(anyllm.WithAPIKey(apiKey)) },
	"mistral":   func(apiKey string) (providers.Provider, error) { return mistral.New(anyllm.WithAPIKey(apiKey)) },
	"deepseek":  func(apiKey string) (providers.Provider, error) { return deepseek.New(anyllm.WithAPIKey(apiKey)) },
	"groq":      func(apiKey string) (providers.Provider, error) { return groq.New(anyllm.WithAPIKey(apiKey)) },
	"zai":       func(apiKey string) (providers.Provider, error) { return zai.New(anyllm.WithAPIKey(apiKey)) },
	"ollama":    func(_ string) (providers.Provider, error) { return ollama.New() },
	"llamacpp":  func(_ string) (providers.Provider, error) { return llamacpp.New() },
	"llamafile": func(_ string) (providers.Provider, error) { return llamafile.New() },
}

func Resolve(spec string, apiKey string) (providers.Provider, string, error) {
	//spec is supposed to be provider/model format
	name, model, ok := strings.Cut(spec, "/")
	if !ok {
		return nil, "", fmt.Errorf("must be provider/model: %q", spec)
	}

	f, ok := registry[name]
	if !ok {
		return nil, "", fmt.Errorf("unknown provider %q", name)
	}

	p, err := f(apiKey)
	if err != nil {
		return nil, "", err
	}
	return p, model, nil
}

func NewAnyLlmGoReviewer(loader config.Loader) (review.Reviewer, error) {
	//TODO:refine returned error
	conf, err := loader.Load()
	if err != nil {
		return nil, err
	}
	p, model, err := Resolve(conf.ProviderAndModel, conf.ApiKey)
	if err != nil {
		return nil, err
	}

	return &LlmClient{
			params:   anyllm.CompletionParams{Model: model},
			provider: p},
		nil
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
