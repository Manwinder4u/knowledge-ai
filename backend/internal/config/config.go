package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	Port        string
	DatabaseURL string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName:     getEnv("APP_NAME", "KnowledgeAI"),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}

	log.Println("Configuration loaded")

	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
