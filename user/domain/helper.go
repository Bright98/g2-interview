package domain

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvFile() error {
	return godotenv.Load(".env")
}
func SetError(key string, err string) *Errors {
	return &Errors{Key: key, Error: err}
}
func GenerateID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
func HashString(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}
func GetServerPort() string {
	return ":" + os.Getenv("PORT")
}
