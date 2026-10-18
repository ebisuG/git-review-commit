package llm

import (
	"os"
	"testing"

	"github.com/ebisuG/git-review-commit-v2/internal/config"
)

func TestNewAnyLlmGoReviewer(t *testing.T) {
	openaiLoader := NewOpenaiStubLoader()
	_, err := NewAnyLlmGoReviewer(openaiLoader)
	if err != nil {
		t.Errorf("failed to resolve openai/ChatModelGPT5_4Mini : %v", err)
	}

	anthropicLoader := NewAnthropicStubLoader()
	_, err = NewAnyLlmGoReviewer(anthropicLoader)
	if err != nil {
		t.Errorf("failed to resolve anthropic/ModelClaudeHaiku5_5 : %v", err)
	}
}

type OpenaiStubLoader struct{}
type AnthropicStubLoader struct{}

func (*OpenaiStubLoader) Load() (*config.Config, error) {
	key := os.Getenv("OPENAI_API_KEY")
	conf := config.Config{
		ApiKey:           key,
		ProviderAndModel: "openai/ChatModelGPT5_4Mini",
	}
	return &conf, nil
}

var _ config.Loader = (*OpenaiStubLoader)(nil)

func NewOpenaiStubLoader() config.Loader {
	return &OpenaiStubLoader{}
}

func (*AnthropicStubLoader) Load() (*config.Config, error) {
	key := os.Getenv("ANTHROPIC_API_KEY")
	conf := config.Config{
		ApiKey:           key,
		ProviderAndModel: "anthropic/ModelClaudeHaiku5_5",
	}
	return &conf, nil
}

var _ config.Loader = (*AnthropicStubLoader)(nil)

func NewAnthropicStubLoader() config.Loader {
	return &AnthropicStubLoader{}
}
