package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    string
	Backends      []string
	RedisAddr     string
	CacheTTL      time.Duration
	JWTSecret     string
	SkipAuthPaths []string
}

func LoadConfig() *Config {
	viper.SetConfigFile("config/config.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return &Config{
		ServerPort:    viper.GetString("server.port"),
		Backends:      viper.GetStringSlice("server.backends"),
		RedisAddr:     viper.GetString("redis.addr"),
		CacheTTL:      viper.GetDuration("cache.ttl"),
		JWTSecret:     viper.GetString("auth.jwt_secret"),
		SkipAuthPaths: viper.GetStringSlice("auth.skip_auth_paths"),
	}
}
