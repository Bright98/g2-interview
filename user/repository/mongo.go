package repository

import (
	"context"
	"g2/user/domain"
	"g2/user/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	MongoTimeout  time.Duration
	MongoDatabase *mongo.Database
)

func (r Repository) InsertUserRepository(user *domain.Users) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.UserCollection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return domain.SetError(variables.CantInsertErr, err.Error())
	}
	return nil
}
func (r Repository) EditUserRepository(user *domain.Users) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.UserCollection)
	filter := bson.M{"_id": user.Id, "status": bson.M{"$ne": variables.RemovedStatus}}
	update := bson.M{"$set": user}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.SetError(variables.CantInsertErr, err.Error())
	}
	return nil
}
func (r Repository) RemoveUserRepository(id string) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.UserCollection)
	filter := bson.M{"_id": id, "status": bson.M{"$ne": variables.RemovedStatus}}
	update := bson.M{"$set": bson.M{"status": variables.RemovedStatus}}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.SetError(variables.CantRemoveErr, err.Error())
	}
	if res.MatchedCount == 0 {
		return domain.SetError(variables.NotFoundErr, "")
	}
	return nil
}
func (r Repository) GetUserByIDRepository(id string) (*domain.Users, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.UserCollection)
	filter := bson.M{"_id": id, "status": bson.M{"$ne": variables.RemovedStatus}}
	user := &domain.Users{}
	err := collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, domain.SetError(variables.ServiceUnknownErr, err.Error())
	}
	return user, nil
}
func (r Repository) GetUserListRepository() ([]domain.Users, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.UserCollection)
	filter := bson.M{"status": bson.M{"$ne": variables.RemovedStatus}}
	res, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, domain.SetError(variables.ServiceUnknownErr, err.Error())
	}

	var result []domain.Users
	err = res.All(ctx, &result)
	if err != nil {
		return nil, domain.SetError(variables.ServiceUnknownErr, err.Error())
	}

	err = res.Close(ctx)
	if err != nil {
		return nil, domain.SetError(variables.ServiceUnknownErr, err.Error())
	}

	return result, nil
}
func (r Repository) GetUserByEmailPasswordRepository(email, password string) (*domain.Users, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.UserCollection)
	user := &domain.Users{}
	filter := bson.M{"email": email, "password": password, "status": bson.M{"$ne": variables.RemovedStatus}}
	err := collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, domain.SetError(variables.ServiceUnknownErr, err.Error())
	}
	return user, nil
}
func (r Repository) GetUserByEmailRepository(email string) (*domain.Users, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.UserCollection)
	user := &domain.Users{}
	filter := bson.M{"email": email, "status": bson.M{"$ne": variables.RemovedStatus}}
	err := collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, domain.SetError(variables.ServiceUnknownErr, err.Error())
	}
	return user, nil
}
