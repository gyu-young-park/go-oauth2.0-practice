package config

import (
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

const ENV_FILE = ".env"

var ConfigMap *Config

type Config struct {
	Port              string
	GoogleOauthConfig *oauth2.Config
}

func SetupConfig() {
	err := godotenv.Load(ENV_FILE)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ConfigMap = &Config{
		Port:              ":8080",
		GoogleOauthConfig: startGooleOauthConfig(),
	}
}
