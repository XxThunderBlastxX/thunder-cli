package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	BaseApiUrl string

	AuthDomain   string
	AuthClientId string
	AuthScope    string
	AuthAudience string

	Viper *viper.Viper
}

func NewAppConfig() (*Config, error) {
	var v = viper.New()

	var (
		thunderAPIBaseUrl = "https://api.koustav.dev"
		authDomain        = "thunder.jp.auth0.com"
		authClientID      = "X43uDaR6gjwJEFPEdP7jNGxXTlCPAjfa"
		authAudience      = "https://thunder.jp.auth0.com/api/v2/"
		authScope         = "openid"
	)

	config := Config{}

	config.AuthDomain = authDomain
	config.AuthClientId = authClientID
	config.AuthScope = authScope
	config.AuthAudience = authAudience
	config.BaseApiUrl = thunderAPIBaseUrl

	//Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	//Get absolute path of config directory
	dirPath, err := filepath.Abs(homeDir + "/.config/thunder-cli")
	if err != nil {
		return nil, err
	}

	// Check if directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	v.AddConfigPath("/$HOME/.config/thunder-cli")
	v.SetConfigName("config")
	v.SetConfigType("json")

	config.Viper = v

	return &config, nil
}
