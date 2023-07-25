package rest

import (
	"g2/api-gateway/messaging/events"
	"g2/api-gateway/send/grpc"
)

type RestHandler struct {
	Grpc *grpc.GrpcClient
	Msg  *events.RabbitConsumer
}

func NewRestApi(grpcClient *grpc.GrpcClient, message *events.RabbitConsumer) *RestHandler {
	return &RestHandler{Grpc: grpcClient, Msg: message}
}
