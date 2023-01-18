package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(name string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
	return os.Getenv(name)
}
