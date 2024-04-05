package config

import (
	"embed"
	"io/fs"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var f embed.FS

type Config struct {
	Name        string `yaml:"name"`
	Link        string `yaml:"link"`
	Description string `yaml:"description"`
	Auth        struct {
		Domain   string `yaml:"domain"`
		ClientId string `yaml:"client_id"`
		Audience string `yaml:"audience"`
		Scope    string `yaml:"scope"`
	} `yaml:"auth"`
}

func NewAppConfig() (*Config, error) {
	config := Config{}

	data, err := fs.ReadFile(f, "config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
