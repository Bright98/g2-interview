package grpc

import (
	"context"
	pb "g2/proto/todo"
	"g2/todo/domain"
)

// write for necessary usage (not used in api-gateway)
func (s GrpcServer) InsertTodoItem(ctx context.Context, in *pb.TodoItem) (*pb.InsertedIDResponse, error) {
	request := &domain.TodoItems{}
	request.TodoListID = in.GetTodoListId()
	request.Title = in.GetTitle()
	request.Priority = in.GetPriority()

	insertedID, err := s.domain.InsertTodoItemService(request)
	response := &pb.IDRequest{Id: insertedID}
	return &pb.InsertedIDResponse{Data: response, Error: domain.MapDomainGrpcError(err)}, nil
}
func (s GrpcServer) EditTodoItem(ctx context.Context, in *pb.TodoItem) (*pb.ErrorResponse, error) {
	request := &domain.TodoItems{}
	request.Id = in.GetId()
	request.TodoListID = in.GetTodoListId()
	request.Title = in.GetTitle()
	request.Priority = in.GetPriority()
	request.Status = int8(in.GetStatus())

	err := s.domain.EditTodoItemService(request)
	return &pb.ErrorResponse{Key: err.Key, Error: err.Error}, nil
}
func (s GrpcServer) RemoveTodoItem(ctx context.Context, in *pb.IDRequest) (*pb.ErrorResponse, error) {
	err := s.domain.RemoveTodoItemService(in.GetId())
	return &pb.ErrorResponse{Key: err.Key, Error: err.Error}, nil
}

// use in api gateway
func (s GrpcServer) GetTodoItemByID(ctx context.Context, in *pb.IDRequest) (*pb.TodoItemResponse, error) {
	TodoItem, err := s.domain.GetTodoItemByIDService(in.GetId())
	return &pb.TodoItemResponse{
		Data:  domain.MapTodoItemToGrpcTodoItem(TodoItem),
		Error: domain.MapDomainGrpcError(err),
	}, nil
}
func (s GrpcServer) GetTodoItemList(ctx context.Context, in *pb.IDRequest) (*pb.TodoItemListResponse, error) {
	TodoItems, err := s.domain.GetTodoItemListService(in.GetId())
	var TodoItemsRes []*pb.TodoItem
	for _, todo := range TodoItems {
		TodoItemsRes = append(TodoItemsRes, domain.MapTodoItemToGrpcTodoItem(&todo))
	}
	return &pb.TodoItemListResponse{Data: TodoItemsRes, Error: domain.MapDomainGrpcError(err)}, nil
}
