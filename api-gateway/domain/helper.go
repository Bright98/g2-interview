package domain

import (
	"github.com/joho/godotenv"
)

func LoadEnvFile() error {
	return godotenv.Load(".env")
}
