package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost          string
	DBPort          string
	DBName          string
	DBUser          string
	DBPassword      string
	ServerPort      string
	UserServiceHost string
	UserServicePort string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBName:          getEnv("DB_NAME", "post_db"),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPassword:      getEnv("DB_PASSWORD", ""),
		ServerPort:      getEnv("SERVER_PORT", "50052"),
		UserServiceHost: getEnv("USER_SERVICE_HOST", "localhost"),
		UserServicePort: getEnv("USER_SERVICE_PORT", "50051"),
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func (c *Config) GetUserServiceAddress() string {
	return fmt.Sprintf("%s:%s", c.UserServiceHost, c.UserServicePort)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
