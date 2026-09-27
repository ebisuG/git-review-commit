package config

import (
	"fmt"
	"testing"
)

func TestFindConfig(t *testing.T) {
	result, err := findConfigDir()
	if err != nil {
		t.Errorf("failed to find directory executable located")
	} else {
		t.Logf("test succeded, temporal directory is : %v", result)
	}
}

func TestNewYamlLoader(t *testing.T) {
	loader := NewYamlLoader("config.yaml")
	_, err := loader.Load()
	if err != nil {
		t.Log("test succeeded because there is no config.yaml file")
	} else {
		t.Errorf("this test should be error.")
		fmt.Println(err)
	}
}
