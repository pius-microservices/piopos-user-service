package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	BaseURL   string
	Mode      string
	JwtSecret string

	UserDBPort     string
	UserDBUsername string
	UserDBPassword string
	UserDBName     string
	UserDBHost     string
}

var (
	envConfig *Config
	once      sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file")
		}

		envConfig = &Config{
			Port:      os.Getenv("APP_PORT"),
			BaseURL:   os.Getenv("BASE_URL"),
			JwtSecret: os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
			Mode:      os.Getenv("MODE"),

			UserDBPort:     os.Getenv("USER_SERVICE_DB_PORT"),
			UserDBUsername: os.Getenv("USER_SERVICE_DB_USERNAME"),
			UserDBPassword: os.Getenv("USER_SERVICE_DB_PASSWORD"),
			UserDBName:     os.Getenv("USER_SERVICE_DB_NAME"),
			UserDBHost:     os.Getenv("USER_SERVICE_DB_HOST"),
		}
	})

	return envConfig
}
