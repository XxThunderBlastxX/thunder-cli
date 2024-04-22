package config

import (
	"bytes"
	"embed"

	"github.com/spf13/viper"
)

//go:embed .env
var envFile embed.FS

type Config struct {
	// API
	BaseApiUrl string `mapstructure:"THUNDER_BASE_API_URL"`

	// Auth
	AuthDomain   string `mapstructure:"AUTH_DOMAIN"`
	AuthClientId string `mapstructure:"AUTH_CLIENT_ID"`
	AuthScope    string `mapstructure:"AUTH_SCOPE"`
	AuthAudience string `mapstructure:"AUTH_AUDIENCE"`
}

func NewAppConfig() (*Config, error) {
	config := Config{}

	// Read the embedded .env file
	envContent, err := envFile.ReadFile(".env")
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadConfig(bytes.NewBuffer(envContent)); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
