package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName: getEnv("APP_NAME", "KnowledgeAI"),
		Port:    getEnv("PORT", "8080"),
	}

	log.Println("Configuration loaded")

	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
