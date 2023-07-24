package main

import (
	"g2/user/api/grpc"
	"g2/user/domain"
	"g2/user/messaging/actions"
	"g2/user/repository"
	"log"
	"os"
	"strconv"
)

var (
	rabbitHandler actions.RabbitHandler
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
	_ = grpc.NewGrpcServer(service)
	rabbitHandler = actions.NewRabbitHandler(service)

	//get mongo requirements from env file
	timeout := os.Getenv("MONGO_TIMEOUT")
	mongoUrl := os.Getenv("MONGO_URL")
	database := os.Getenv("MONGO_DATABASE")
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//mongo connection
	err = repository.MongoConnection(mongoUrl, database, timeoutInt)
	if err != nil {
		log.Fatalln(err)
	}

	//get grpc requirements from env file
	port := os.Getenv("PORT")

	//grpc connection
	err = grpc.GrpcServerConnection(port)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	//actions
	actions.RabbitmqListenToActions(rabbitHandler)
}
