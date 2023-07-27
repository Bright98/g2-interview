package grpc

import (
	"g2/api-gateway/domain"
	idpProto "g2/proto/idp"
	ssoProto "g2/proto/sso"
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
	IdpClient  idpProto.IdpServiceClient
	SSOClient  ssoProto.SSOServiceClient
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

	//idp grpc
	idpClient, err := IdpGrpcClientConnection()
	if err != nil {
		log.Fatalln(err)
	}
	grpcClient.IdpClient = idpClient

	//sso grpc
	ssoClient, err := SSOGrpcClientConnection()
	if err != nil {
		log.Fatalln(err)
	}
	grpcClient.SSOClient = ssoClient

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
func IdpGrpcClientConnection() (idpProto.IdpServiceClient, error) {
	conn, err := grpc.Dial(os.Getenv("IDP_GRPC_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := idpProto.NewIdpServiceClient(conn)
	return client, nil
}
func SSOGrpcClientConnection() (ssoProto.SSOServiceClient, error) {
	conn, err := grpc.Dial(os.Getenv("SSO_GRPC_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := ssoProto.NewSSOServiceClient(conn)
	return client, nil
}
