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
	fileName string
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

	yamlFile, err := os.ReadFile(filepath.Join(configDir, y.fileName))
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

func NewYamlLoader(fileName string) *YamlLoader {
	return &YamlLoader{fileName: fileName}
}

var _ Loader = (*YamlLoader)(nil)
