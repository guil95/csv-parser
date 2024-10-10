package config

import (
	"context"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sethvargo/go-envconfig"
)

var AppConfig MyConfig

type MyConfig struct {
	Environment   string `env:"ENVIRONMENT"`
	LogLevel      string `env:"LOG_LEVEL"`
	MaxFileLength int    `env:"MAX_FILE_LENGTH"`
}

func init() {
	if err := envconfig.Process(context.Background(), &AppConfig); err != nil {
		log.Fatal(err)
	}
}
