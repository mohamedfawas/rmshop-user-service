package config

import "os"

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	// Server config
	cfg.Server.Host = getEnvOrDefault("SERVER_HOST", "0.0.0.0")
	cfg.Server.Port = getEnvOrDefault("SERVER_PORT", "50053")

	// Database config
	cfg.Database.Host = getEnvOrDefault("DB_HOST", "localhost")
	cfg.Database.Port = getEnvOrDefault("DB_PORT", "5432")
	cfg.Database.User = getEnvOrDefault("DB_USER", "postgres")
	cfg.Database.Password = getEnvOrDefault("DB_PASSWORD", "postgres")
	cfg.Database.DBName = getEnvOrDefault("DB_NAME", "rmshop")
	cfg.Database.SSLMode = getEnvOrDefault("DB_SSLMODE", "disable")

	return cfg, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
