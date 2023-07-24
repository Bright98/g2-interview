package main

import (
	"g2/api-gateway/api/rest"
	"g2/api-gateway/domain"
	"g2/api-gateway/send/grpc"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	Gin         *gin.Engine
	RestHandler *rest.RestHandler
)

func init() {
	//load env file
	err := domain.LoadEnvFile()
	if err != nil {
		log.Fatalln(err)
	}

	//handle directory connection
	grpcClient := grpc.NewGrpcClient()
	RestHandler = rest.NewRestApi(grpcClient)

	//initialize gin
	Gin = gin.Default()
}
func main() {
	//define routes
	Gin.GET("/api/user/users/id/:user-id", RestHandler.GetUserByID)
	Gin.GET("/api/user/users", RestHandler.GetUserList)

	//run gin
	err := Gin.Run(domain.GetServerPort())
	if err != nil {
		log.Fatalln(err.Error())
	}
}
