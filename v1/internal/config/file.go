package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	ApiKey string `yaml:"API_KEY"`
}

type YamlLoader struct {
	path string
}

func (y *YamlLoader) Load() (*Config, error) {
	filename, _ := filepath.Abs(string(y.path))
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return &Config{}, err
	}
	var yamlConfig YamlConfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		return &Config{}, err
	}
	var config Config
	config = Config{ApiKey: yamlConfig.ApiKey}
	return &config, nil
}

func NewYamlLoader(path string) *YamlLoader {
	return &YamlLoader{path: path}
}

var _ Loader = (*YamlLoader)(nil)
