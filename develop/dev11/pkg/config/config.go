package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config
type config struct {
	Port string `yaml:"port"`
}

func (c *config) LoadConfig(filename string) (config config, err error) {
	// Чтение содержимого файла
	file, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	// Разбор YAML
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, nil
	}

	return config, nil
}

func NewConfig() config {
	return config{}
}
