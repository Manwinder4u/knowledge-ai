package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	Port        string
	DatabaseURL string

	JWTSecret      string
	JWTExpiryHours int

	UploadPath string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName:     getEnv("APP_NAME", "KnowledgeAI"),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),

		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpiryHours: getEnvAsInt("JWT_EXPIRY_HOURS", 24),

		UploadPath: getEnv("UPLOAD_PATH", "storage/uploads"),
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

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return i

}
