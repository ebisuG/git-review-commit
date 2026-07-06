package config

// Loader returns configuration

type Loader interface {
	Load() (*Config, error)
}

type Config struct {
	ApiKey string
}
