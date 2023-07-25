package main

import (
	"g2/api-gateway/api/rest"
	"g2/api-gateway/domain"
	"g2/api-gateway/messaging/events"
	"g2/api-gateway/send/grpc"
	"github.com/gin-gonic/gin"
	"log"
	"os"
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

	//get rabbitmq requirements from env file
	rabbitAddress := os.Getenv("RABBITMQ_URL")

	//handle directory connection
	grpcClient := grpc.NewGrpcClient()
	rabbitmq := events.NewRabbitMQ(rabbitAddress)
	RestHandler = rest.NewRestApi(grpcClient, rabbitmq)

	//initialize gin
	Gin = gin.Default()
}
func main() {
	//define routes

	//users
	Gin.GET("/api/user/users/id/:user-id", RestHandler.GetUserByID)
	Gin.GET("/api/user/users", RestHandler.GetUserList)
	Gin.POST("/api/user/users", RestHandler.InsertUser)
	Gin.PUT("/api/user/users/id/:user-id", RestHandler.EditUser)
	Gin.DELETE("/api/user/users/id/:user-id", RestHandler.RemoveUser)

	//todo list
	Gin.GET("/api/todo/todo-lists/id/:todo-list-id", RestHandler.GetTodoListByID)
	Gin.GET("/api/todo/todo-lists", RestHandler.GetTodoListList)
	Gin.POST("/api/todo/todo-lists", RestHandler.InsertTodoList)
	Gin.PUT("/api/todo/todo-lists/id/:todo-list-id", RestHandler.EditTodoList)
	Gin.DELETE("/api/todo/todo-lists/id/:todo-list-id", RestHandler.RemoveTodoList)

	//todo items
	Gin.GET("/api/todo/todo-items/id/:todo-item-id", RestHandler.GetTodoItemByID)
	Gin.GET("/api/todo/todo-lists/id/:todo-list-id/todo-items", RestHandler.GetTodoItemList)
	Gin.POST("/api/todo/todo-items", RestHandler.InsertTodoItem)
	Gin.PUT("/api/todo/todo-items/id/:todo-item-id", RestHandler.EditTodoItem)
	Gin.DELETE("/api/todo/todo-items/id/:todo-item-id", RestHandler.RemoveTodoItem)

	//run gin
	err := Gin.Run(domain.GetServerPort())
	if err != nil {
		log.Fatalln(err.Error())
	}
}
