package main

import (
	"g2/user/domain"
	"g2/user/repository"
	"log"
	"os"
	"strconv"
)

func init() {
	//load env file
	err := domain.LoadEnvFile()
	if err != nil {
		log.Fatalln(err.Error())
	}

	//get mongo requirements from env file
	timeout := os.Getenv("MONGO_TIMEOUT")
	mongoUrl := os.Getenv("MONGO_URL")
	database := os.Getenv("MONGO_DATABASE")
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//mongo connection
	err = repository.MongoConnection(mongoUrl, database, timeoutInt)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
}
