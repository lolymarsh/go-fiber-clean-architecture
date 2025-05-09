package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	App   *App
	Db    *Database
	Auth  *Auth
	Redis *Redis
}

var (
	once sync.Once
)

func LoadConfig() (*Config, error) {
	configInstance := &Config{}
	once.Do(func() {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")

		viper.AutomaticEnv()

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("failed to read .env file: %w", err))
		}
		if err := viper.Unmarshal(configInstance); err != nil {
			panic(fmt.Errorf("failed to unmarshal config: %w", err))
		}
	})

	return configInstance, nil
}
