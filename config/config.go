package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// LoadDBConfig loads database configuration from environment variables.
func LoadDBConfig() *DBConfig {
	return &DBConfig{
		Host:     mustGetEnv("DB_HOST"),
		Port:     mustGetEnv("DB_PORT"),
		User:     mustGetEnv("DB_USER"),
		Password: mustGetEnv("DB_PASSWORD"),
		DBName:   mustGetEnv("DB_NAME"),
		SSLMode:  mustGetEnv("DB_SSLMODE"),
	}
}

// GetPostgresDSN constructs the PostgreSQL connection string in URL format.
func GetPostgresDSN(cfg *DBConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)
}

// mustGetEnv retrieves the value of the environment variable or raises an error if not found.
func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Sprintf("Environment variable %s is required but not set", key))
	}
	return value
}
