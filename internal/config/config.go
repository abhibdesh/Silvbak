package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	DBName   string
	Port     string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file is found")
	}
	return &Config{
		MongoURI: mustGet("MONGO_URL"),
		DBName:   getOrDefault("DB_NAME", "silvbak"),
		Port:     getOrDefault("PORT", "8080"),
	}
}

func mustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("required env var %s is not set", key)
	}
	return v
}

func getOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
