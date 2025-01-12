package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Services map[string]ServiceConfig `yaml:"services"`
	Global   GlobalConfig             `yaml:"global"`
}

type ServiceConfig struct {
	BaseURL           string `yaml:"baseUrl"`
	IPv4PublicAddress string `yaml:"ipv4PublicAddress,omitempty"`
	APIKey            string `yaml:"apiKey"`
	APIVersion        string `yaml:"apiVersion"`
	Timeout           string `yaml:"timeout"`
}

type GlobalConfig struct {
	RetryCount int    `yaml:"retryCount"`
	LogLevel   string `yaml:"logLevel"`
	Port       string `yaml:"port"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
