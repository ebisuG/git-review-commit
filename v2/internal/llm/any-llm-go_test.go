package llm

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ebisuG/git-review-commit-v2/internal/config"
	"github.com/joho/godotenv"
)

func TestNewAnyLlmGoReviewer(t *testing.T) {
	openaiLoader := NewOpenaiStubLoader()
	openAiReviewer, err := NewAnyLlmGoReviewer(openaiLoader)
	if err != nil {
		t.Errorf("failed to resolve openai/gpt-5.4-mini : %w", err)
	}
	res, err := openAiReviewer.Review("Say hello, and what is your name?")
	if err != nil {
		t.Errorf("execution of review failed : %w", err)
	}
	fmt.Println("OpenAI : ", res)

	anthropicLoader := NewAnthropicStubLoader()
	anthropicReviewer, err := NewAnyLlmGoReviewer(anthropicLoader)
	if err != nil {
		t.Errorf("failed to resolve anthropic/claude-haiku-5-5 : %w", err)
	}
	res, err = anthropicReviewer.Review("Say hello, and what is your name??")
	if err != nil {
		t.Errorf("execution of review failed : %w", err)
	}
	fmt.Println("Anthropic : ", res)
}

type OpenaiStubLoader struct{}
type AnthropicStubLoader struct{}

func (*OpenaiStubLoader) Load() (*config.Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("failed loading .env : %v", err)
	}
	key, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok || key == "" {
		return &config.Config{}, fmt.Errorf("not found API key from env")
	}
	conf := config.Config{
		ApiKey:           key,
		ProviderAndModel: "openai/gpt-5.4-mini",
	}
	return &conf, nil
}

var _ config.Loader = (*OpenaiStubLoader)(nil)

func NewOpenaiStubLoader() config.Loader {
	return &OpenaiStubLoader{}
}

func (*AnthropicStubLoader) Load() (*config.Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("failed loading .env : %v", err)
	}
	key, ok := os.LookupEnv("ANTHROPIC_API_KEY")
	if !ok || key == "" {
		return &config.Config{}, fmt.Errorf("not found API key from env")
	}
	conf := config.Config{
		ApiKey:           key,
		ProviderAndModel: "anthropic/claude-haiku-5-5",
	}
	return &conf, nil
}

var _ config.Loader = (*AnthropicStubLoader)(nil)

func NewAnthropicStubLoader() config.Loader {
	return &AnthropicStubLoader{}
}
