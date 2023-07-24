package grpc

import (
	"g2/api-gateway/domain"
	pb "g2/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

type GrpcClient struct {
	UserClient pb.UserServiceClient
}

func NewGrpcClient() *GrpcClient {
	//load env file
	err := domain.LoadEnvFile()
	if err != nil {
		log.Fatalln(err)
	}

	grpcClient := &GrpcClient{}

	//user grpc
	client, err := UserGrpcClientConnection()
	if err != nil {
		log.Fatalln(err)
	}
	grpcClient.UserClient = client

	return grpcClient
}

func UserGrpcClientConnection() (pb.UserServiceClient, error) {
	conn, err := grpc.Dial(os.Getenv("USER_GRPC_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewUserServiceClient(conn)
	return client, nil
}
