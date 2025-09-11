package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	RedisAddr       string
	RedisPassword   string
	PublishInterval time.Duration
	Channel         string
}

func FromEnv() *Config {
	intervalStr := getEnv("PUBLISH_INTERVAL", "5")
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		interval = 5
	}
	return &Config{
		RedisAddr:       getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:   getEnv("REDIS_PASSWORD", ""),
		PublishInterval: time.Duration(interval) * time.Second,
		Channel:         getEnv("REDIS_CHANNEL", "tokens"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
