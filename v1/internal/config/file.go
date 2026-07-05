package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	ApiKey string `yaml:"API_KEY"`
}

type YamlPath string

func (y *YamlPath) Load() (*Config, error) {
	filename, _ := filepath.Abs(string(*y))
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
	config.ApiKey = Config(yamlConfig).ApiKey
	return &config, nil
}

func NewYamlLoader(path string) YamlPath {
	return YamlPath(path)
}

var _ Loader = (*YamlPath)(nil)
