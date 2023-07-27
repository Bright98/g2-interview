package main

import (
	"g2/auth/sso/api/grpc"
	"g2/auth/sso/domain"
	"g2/auth/sso/repository"
	"log"
	"os"
	"strconv"
)

var (
	grpcHandler *grpc.GrpcServer
	port        string
)

func init() {
	//load env file
	err := domain.LoadEnvFile()
	if err != nil {
		log.Fatalln(err.Error())
	}

	//handle directory connection
	repo := repository.NewRepository()
	service := domain.NewService(repo)
	grpcHandler = grpc.NewGrpcServer(service)

	//get redis requirements from env file
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	db := os.Getenv("REDIS_DB")
	dbInt, err := strconv.Atoi(db)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//redis connection
	err = repository.RedisConnection(redisAddress, redisPassword, dbInt)
	if err != nil {
		log.Fatalln(err)
	}

	//get grpc requirements from env file
	port = os.Getenv("PORT")

}
func main() {
	//grpc connection
	grpc.GrpcServerConnection(grpcHandler, port)
}
