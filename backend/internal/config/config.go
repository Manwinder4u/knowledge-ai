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

	UploadPath    string
	MaxUploadSize int64

	OllamaURL       string
	EmbeddingModel  string
	OllamaChatModel string
}

func Load() *Config {
	_ = godotenv.Load()

	maxUploadSize, err := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "20971520"), 10, 64)
	if err != nil {
		log.Fatal("invalid MAX_UPLOAD_SIZE")
	}

	cfg := &Config{
		AppName:     getEnv("APP_NAME", "KnowledgeAI"),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),

		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpiryHours: getEnvAsInt("JWT_EXPIRY_HOURS", 24),

		UploadPath:    getEnv("UPLOAD_PATH", "storage/uploads"),
		MaxUploadSize: maxUploadSize,

		OllamaURL:       getEnv("OLLAMA_URL", "http://localhost:11434"),
		EmbeddingModel:  getEnv("EMBEDDING_MODEL", "nomic-embed-text"),
		OllamaChatModel: getEnv("OLLAMA_CHAT_MODEL", "qwen2.5:7b"),
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
