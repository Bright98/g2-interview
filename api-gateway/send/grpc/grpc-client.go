package grpc

import (
	"g2/api-gateway/domain"
	todoProto "g2/proto/todo"
	userProto "g2/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

type GrpcClient struct {
	UserClient userProto.UserServiceClient
	TodoClient todoProto.TodoServiceClient
}

func NewGrpcClient() *GrpcClient {
	//load env file
	err := domain.LoadEnvFile()
	if err != nil {
		log.Fatalln(err)
	}

	grpcClient := &GrpcClient{}

	//user grpc
	userClient, err := UserGrpcClientConnection()
	if err != nil {
		log.Fatalln(err)
	}
	grpcClient.UserClient = userClient

	//todo grpc
	todoClient, err := TodoGrpcClientConnection()
	if err != nil {
		log.Fatalln(err)
	}
	grpcClient.TodoClient = todoClient

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
func TodoGrpcClientConnection() (todoProto.TodoServiceClient, error) {
	conn, err := grpc.Dial(os.Getenv("TODO_GRPC_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := todoProto.NewTodoServiceClient(conn)
	return client, nil
}
