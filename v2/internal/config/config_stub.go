package config

type StubLoader struct{}

func (*StubLoader) Load() (*Config, error) {
	StubConfig := &Config{
		ApiKey:           "dummy",
		ProviderAndModel: "openai/gpt5",
	}
	return StubConfig, nil
}

func NewStubLoader() *StubLoader {
	return &StubLoader{}
}

var _ Loader = (*StubLoader)(nil)
