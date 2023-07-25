package repository

import (
	"context"
	"g2/user/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Repository struct {
}

func NewRepository() domain.RepositoryInterface {
	repo := &Repository{}
	return repo
}

// mongo
func MongoConnection(mongoUrl string, database string, timeout int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	MongoTimeout = time.Duration(timeout) * time.Second
	MongoDatabase = client.Database(database)

	return nil
}
