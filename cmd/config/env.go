package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBuser string
	DBpass string
	DBAddr string
	DBName string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "127.0.0.1"),
		Port:       getEnv("PORT", "8080"),
		DBuser:     getEnv("DB_USER", "root"),
		DBpass:     getEnv("DB_PASS", "password"),
		DBAddr:     getEnv("DB_ADDR", "127.0.0.1:3306"),
		DBName:     getEnv("DB_NAME", "ecom"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
