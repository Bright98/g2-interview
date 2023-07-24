package main

import (
	"g2/api-gateway/api/rest"
	"g2/api-gateway/send/grpc"
)

func init() {
	//handle directory connection
	grpcClient := grpc.NewGrpcClient()
	_ = rest.NewRestApi(grpcClient)

}
func main() {

}
