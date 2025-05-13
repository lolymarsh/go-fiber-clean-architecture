package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	App   *App
	Db    *Database
	Redis *Redis
	Auth  *Auth
}

var (
	once sync.Once
)

func LoadConfig() *Config {
	godotenv.Load()
	configInstance := &Config{
		App: &App{
			Port:         getEnv("APP_PORT", "8080"),
			BodyLimit:    getEnvAsInt("APP_BODY_LIMIT", 20),
			IdleTimeout:  getEnvAsInt("APP_IDLE_TIMEOUT", 60),
			AllowOrigins: getEnv("APP_ALLOW_ORIGINS", "http://localhost,http://example.com"),
		},
		Db: &Database{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "3306"),
			User:            getEnv("DB_USER", "root"),
			Password:        getEnv("DB_PASSWORD", ""),
			Name:            getEnv("DB_NAME", "test"),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: getEnvAsInt("DB_CONN_MAX_LIFETIME", 60),
		},
		Redis: &Redis{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		Auth: &Auth{
			SecretKey: getEnv("AUTH_SECRET_KEY", "secret_key"),
			Expire:    getEnvAsInt("AUTH_EXPIRE", 3600),
		},
	}

	return configInstance
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		return value == "true" || value == "1"
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
