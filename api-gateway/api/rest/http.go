package rest

import (
	"g2/api-gateway/send/grpc"
)

type RestHandler struct {
	Grpc *grpc.GrpcClient
}

func NewRestApi(grpcClient *grpc.GrpcClient) *RestHandler {
	return &RestHandler{Grpc: grpcClient}
}
