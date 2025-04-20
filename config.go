package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetAPIKey() string {
	return os.Getenv("API_KEY")
}

//8XDV7IIC6AF9EN5TQVHZTSRX648HN43I1