package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func NewConfig() (*Config, error) {
	var cfg Config
	err := godotenv.Load("config/local-dev.env")
	if err != nil {
		log.Fatal(err)
	}

	err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	return &cfg, nil
}

type Config struct {
	YoutubeAPIKey       string `env:"YOUTUBE_API_KEY"`
	YoutubeClientID     string `env:"YOUTUBE_CLIENT_ID"`
	YoutubeClientSecret string `env:"YOUTUBE_CLIENT_SECRET"`
}
