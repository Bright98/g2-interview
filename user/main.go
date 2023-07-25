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
	grpcHandler   *grpc.GrpcServer
	port          string
)

func init() {
	//load env file
	err := domain.LoadEnvFile()
	if err != nil {
		log.Fatalln(err.Error())
	}

	//get rabbitmq requirements from env file
	rabbitAddress := os.Getenv("RABBITMQ_URL")

	//handle directory connection
	repo := repository.NewRepository()
	service := domain.NewService(repo)
	grpcHandler = grpc.NewGrpcServer(service)
	rabbitHandler = actions.NewRabbitHandler(rabbitAddress, service)

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
	port = os.Getenv("PORT")

}

func main() {
	//actions
	actions.RabbitmqListenToActions(rabbitHandler)

	//grpc connection
	grpc.GrpcServerConnection(grpcHandler, port)

	//Gin := gin.Default()
	////define routes
	//Gin.GET("/api/user/users/id/:user-id/test", restHandler.GetUserByIDTest)
	//
	//err := Gin.Run(domain.GetServerPort())
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
}
