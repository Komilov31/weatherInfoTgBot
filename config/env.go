package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TgBotToken      string
	WeatherApiToken string
	DBUser          string
	DBPassword      string
	DBHost          string
	DBPort          string
	DBName          string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Something went wrong with loading .env file")
	}

	return Config{
		TgBotToken:      getEnv("tg_bot_token", ""),
		WeatherApiToken: getEnv("weather_api_token", ""),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPassword:      getEnv("DB_PASSWORD", ""),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBName:          getEnv("DB_NAME", "postgres"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
