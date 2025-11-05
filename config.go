package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadEventFromGitHub() (*Event, error) {
	evPath := os.Getenv("GITHUB_EVENT_PATH")
	if evPath == "" {
		return nil, fmt.Errorf("GITHUB_EVENT_PATH is not set")
	}

	ev, err := os.ReadFile(evPath)
	if err != nil {
		return nil, err
	}

	var event Event
	err = json.Unmarshal(ev, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
