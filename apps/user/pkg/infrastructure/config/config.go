package config

import "os"

type Config struct {
	Host string
	Port string

	PG_HOST     string
	PG_PORT     string
	PG_USER     string
	PG_PASSWORD string
	PG_DATABASE string

	LogLevel string
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func Get() Config {
	return Config{
		// server config
		Host: getEnvOrDefault("HOST", "0.0.0.0"),
		Port: getEnvOrDefault("PORT", "8080"),

		// postgres config
		PG_HOST:     getEnvOrDefault("PG_HOST", "localhost"),
		PG_PORT:     getEnvOrDefault("PG_PORT", "5432"),
		PG_USER:     getEnvOrDefault("PG_USER", "fitsphere"),
		PG_PASSWORD: getEnvOrDefault("PG_PASSWORD", "fitsphere"),
		PG_DATABASE: getEnvOrDefault("PG_DATABASE", "fitsphere"),

		// logger config
		LogLevel: getEnvOrDefault("LOG_LEVEL", "info"),
	}
}
