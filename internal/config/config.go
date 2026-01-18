package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL string
	Port   string
}

func Load() (*Config, error) {

	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: .env file not found using environment variables")
	}

	var config *Config = &Config{
		DB_URL: os.Getenv("DB_URL"),
		Port:   os.Getenv("PORT"),
	}

	return config, nil

}
