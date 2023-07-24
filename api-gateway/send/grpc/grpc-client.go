package grpc

import (
	pb "g2/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

type GrpcClient struct {
	UserClient pb.UserServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	grpcClient := &GrpcClient{}
	var err error

	grpcClient.UserClient, err = UserGrpcClientConnection()
	return grpcClient, err
}

func UserGrpcClientConnection() (pb.UserServiceClient, error) {
	conn, err := grpc.Dial(os.Getenv("USER_GRPC_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewUserServiceClient(conn)
	return client, nil
}
