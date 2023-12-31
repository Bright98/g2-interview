package domain

import (
	"crypto/sha256"
	"fmt"
	pb "g2/proto/user"
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
func MapDomainGrpcError(err *Errors) *pb.ErrorResponse {
	if err == nil {
		return nil
	}
	return &pb.ErrorResponse{Key: err.Key, Error: err.Error}
}
func MapUserToGrpcUser(user *Users) *pb.User {
	if user == nil {
		return nil
	}
	userRes := &pb.User{}
	userRes.Id = user.Id
	userRes.Name = user.Name
	userRes.Email = user.Email
	userRes.Password = ""
	userRes.Status = int32(user.Status)
	return userRes
}
func (d *DomainService) InsertErrorLogFunction(err *Errors, collection, desc string) {
	log := &ErrorLogs{}
	log.Key = err.Key
	log.Error = err.Error
	log.Collection = collection
	log.Description = desc
	d.Repo.InsertErrorLogRepository(log)
}
