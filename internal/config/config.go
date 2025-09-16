package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
	BDName string
	BDKey  string
}

func NewConfig() *Config {
	if err := godotenv.Load("./settings.env"); err != nil {
		log.Panic(err)
	}

	// key, exists := os.LookupEnv("TGBOT_API_KEY")
	// if exists == false {
	// 	log.Panic("The API key", exists)
	// }

	return &Config{
		APIKey: os.Getenv("TGBOT_API_KEY"),
	}
}
