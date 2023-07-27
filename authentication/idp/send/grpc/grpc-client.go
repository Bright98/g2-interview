package grpc

import (
	userProto "g2/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

type GrpcClient struct {
	UserClient userProto.UserServiceClient
}

func NewGrpcClient() *GrpcClient {
	grpcClient := &GrpcClient{}

	//user grpc
	userClient, err := UserGrpcClientConnection()
	if err != nil {
		log.Fatalln(err)
	}
	grpcClient.UserClient = userClient

	return grpcClient
}

func UserGrpcClientConnection() (userProto.UserServiceClient, error) {
	conn, err := grpc.Dial(os.Getenv("USER_GRPC_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := userProto.NewUserServiceClient(conn)
	return client, nil
}
