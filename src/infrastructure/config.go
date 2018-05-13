package infrastructure

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	URL       string `mapstructure:"url"`
	AuthToken string `mapstructure:"auth_token"`
}

type DataConfig struct {
	DBPath string `mapstructure:"db_path"`
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Data   DataConfig   `mapstructure:"data"`
}

func ReadConfig() (*Config, error) {
	config := &Config{}

	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	err := v.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
