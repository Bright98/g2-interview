package main

import (
	"g2/api-gateway/api/rest"
	"g2/api-gateway/send/grpc"
	"log"
)

func init() {
	//grpc connection
	grpcClient, err := grpc.NewGrpcClient()
	if err != nil {
		log.Fatalln(err)
	}
	_ = rest.NewRestApi(grpcClient)

}
func main() {

}
