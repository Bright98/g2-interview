package main

import (
	grpcApi "g2/auth/idp/api/grpc"
	"g2/auth/idp/domain"
	grpcClient "g2/auth/idp/send/grpc"
	"log"
	"os"
)

var (
	grpcHandler *grpcApi.GrpcServer
	port        string
)

func init() {
	//load env file
	err := domain.LoadEnvFile()
	if err != nil {
		log.Fatalln(err.Error())
	}

	_grpcClient := grpcClient.NewGrpcClient()
	service := domain.NewService(_grpcClient)
	grpcHandler = grpcApi.NewGrpcServer(service)

	//get grpc requirements from env file
	port = os.Getenv("PORT")
}
func main() {
	//grpc connection
	grpcApi.GrpcServerConnection(grpcHandler, port)
}
