package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	ApiKey           string `yaml:"API_KEY"`
	ProviderAndModel string `yaml:"PROVIDER_AND_MODEL"`
}

type YamlLoader struct {
	yamlFileName        string
	reviewStyleFileName string
}

func findConfigDir() (string, error) {
	toolLocatedPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	toolLocatedPath, err = filepath.EvalSymlinks(toolLocatedPath)
	if err != nil {
		return "", err
	}
	return filepath.Dir(toolLocatedPath), nil
}

func (y *YamlLoader) Load() (*Config, error) {
	configDir, err := findConfigDir()
	if err != nil {
		return &Config{}, err
	}

	yamlFile, err := os.ReadFile(filepath.Join(configDir, y.yamlFileName))
	if err != nil {
		return &Config{}, err
	}
	var yamlConfig YamlConfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		return &Config{}, err
	}

	reviewStyleByte, err := os.ReadFile(filepath.Join(configDir, y.reviewStyleFileName))
	if err != nil {
		return &Config{}, err
	}

	reviewStyle := string(reviewStyleByte)

	var config Config
	config = Config{ApiKey: yamlConfig.ApiKey, ProviderAndModel: yamlConfig.ProviderAndModel, ReviewStyle: reviewStyle}
	return &config, nil
}

func NewYamlLoader(yamlFileName string, reviewStyleFileName string) *YamlLoader {
	return &YamlLoader{yamlFileName: yamlFileName, reviewStyleFileName: reviewStyleFileName}
}

var _ Loader = (*YamlLoader)(nil)
