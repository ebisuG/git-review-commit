package llm

import (
	"fmt"
	"strings"

	"github.com/mozilla-ai/any-llm-go/providers"
	"github.com/mozilla-ai/any-llm-go/providers/anthropic"
	"github.com/mozilla-ai/any-llm-go/providers/openai"
)

type factory func() (providers.Provider, error)

var registry = map[string]factory{
	"openai":    func() (providers.Provider, error) { return openai.New() },
	"anthropic": func() (providers.Provider, error) { return anthropic.New() },
}

func Resolve(spec string) (providers.Provider, string, error) {
	//spec is supposed to be provider/model format
	name, model, ok := strings.Cut(spec, "/")
	if !ok {
		return nil, "", fmt.Errorf("must be provider/model: %q", spec)
	}

	f, ok := registry[name]
	if !ok {
		return nil, "", fmt.Errorf("unknown provider %q", name)
	}

	p, err := f()
	if err != nil {
		return nil, "", err
	}
	return p, model, nil
}
