package repository

import (
	"context"
	"g2/todo/domain"
	"g2/todo/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	MongoTimeout  time.Duration
	MongoDatabase *mongo.Database
)

// todo list
func (r Repository) InsertTodoListRepository(todoList *domain.TodoLists) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.TodoListCollection)
	_, err := collection.InsertOne(ctx, todoList)
	if err != nil {
		return domain.SetError(variables.CantInsertErr, err.Error())
	}
	return nil
}
func (r Repository) EditTodoListRepository(todoList *domain.TodoLists) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.TodoListCollection)
	filter := bson.M{"_id": todoList.Id, "status": bson.M{"$ne": variables.RemovedStatus}}
	update := bson.M{"$set": todoList}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.SetError(variables.CantInsertErr, err.Error())
	}
	return nil
}
func (r Repository) RemoveTodoListRepository(id string) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.TodoListCollection)
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
func (r Repository) GetTodoListByIDRepository(id string) (*domain.TodoLists, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.TodoListCollection)
	filter := bson.M{"_id": id, "status": bson.M{"$ne": variables.RemovedStatus}}
	todoList := &domain.TodoLists{}
	err := collection.FindOne(ctx, filter).Decode(todoList)
	if err != nil {
		return nil, domain.SetError(variables.ServiceUnknownErr, err.Error())
	}
	return todoList, nil
}
func (r Repository) GetTodoListListRepository(skip, limit int64) ([]domain.TodoLists, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.TodoListCollection)
	filter := bson.M{"status": bson.M{"$ne": variables.RemovedStatus}}
	if skip != 0 {
		skip = (skip - 1) * limit
	}
	option := options.FindOptions{Skip: &skip, Limit: &limit}
	res, err := collection.Find(ctx, filter, &option)
	if err != nil {
		return nil, domain.SetError(variables.ServiceUnknownErr, err.Error())
	}

	var result []domain.TodoLists
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

// todo item
func (r Repository) InsertTodoItemRepository(todoItem *domain.TodoItems) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.TodoItemCollection)
	_, err := collection.InsertOne(ctx, todoItem)
	if err != nil {
		return domain.SetError(variables.CantInsertErr, err.Error())
	}
	return nil
}
func (r Repository) EditTodoItemRepository(todoItem *domain.TodoItems) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.TodoItemCollection)
	filter := bson.M{"_id": todoItem.Id, "status": bson.M{"$ne": variables.RemovedStatus}}
	update := bson.M{"$set": todoItem}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.SetError(variables.CantInsertErr, err.Error())
	}
	return nil
}
func (r Repository) RemoveTodoItemRepository(id string) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.TodoItemCollection)
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
func (r Repository) GetTodoItemByIDRepository(id string) (*domain.TodoItems, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.TodoItemCollection)
	filter := bson.M{"_id": id, "status": bson.M{"$ne": variables.RemovedStatus}}
	todoItem := &domain.TodoItems{}
	err := collection.FindOne(ctx, filter).Decode(todoItem)
	if err != nil {
		return nil, domain.SetError(variables.ServiceUnknownErr, err.Error())
	}
	return todoItem, nil
}
func (r Repository) GetTodoItemListRepository(todoListID string) ([]domain.TodoItems, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(variables.TodoItemCollection)
	filter := bson.M{"todo_list_id": todoListID, "status": bson.M{"$ne": variables.RemovedStatus}}
	res, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, domain.SetError(variables.ServiceUnknownErr, err.Error())
	}

	var result []domain.TodoItems
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
