package grpc

import (
	"context"
	pb "g2/proto/todo"
	"g2/todo/domain"
)

// write for necessary usage (not used in api-gateway)
func (s GrpcServer) InsertTodoList(ctx context.Context, in *pb.TodoList) (*pb.InsertedIDResponse, error) {
	request := &domain.TodoLists{}
	request.Name = in.GetName()
	request.Description = in.GetDescription()

	insertedID, err := s.domain.InsertTodoListService(request)
	response := &pb.IDRequest{Id: insertedID}
	return &pb.InsertedIDResponse{Data: response, Error: domain.MapDomainGrpcError(err)}, nil
}
func (s GrpcServer) EditTodoList(ctx context.Context, in *pb.TodoList) (*pb.ErrorResponse, error) {
	request := &domain.TodoLists{}
	request.Id = in.GetId()
	request.Name = in.GetName()
	request.Description = in.GetDescription()
	request.Status = int8(in.GetStatus())

	err := s.domain.EditTodoListService(request)
	return &pb.ErrorResponse{Key: err.Key, Error: err.Error}, nil
}
func (s GrpcServer) RemoveTodoList(ctx context.Context, in *pb.IDRequest) (*pb.ErrorResponse, error) {
	err := s.domain.RemoveTodoListService(in.GetId())
	return &pb.ErrorResponse{Key: err.Key, Error: err.Error}, nil
}

// use in api gateway
func (s GrpcServer) GetTodoListByID(ctx context.Context, in *pb.IDRequest) (*pb.TodoListResponse, error) {
	todoList, err := s.domain.GetTodoListByIDService(in.GetId())
	return &pb.TodoListResponse{
		Data:  domain.MapTodoListToGrpcTodoList(todoList),
		Error: domain.MapDomainGrpcError(err),
	}, nil
}
func (s GrpcServer) GetTodoListList(ctx context.Context, in *pb.SkipLimitRequest) (*pb.TodoListListResponse, error) {
	todoLists, err := s.domain.GetTodoListListService(in.Skip, in.Limit)
	var todoListsRes []*pb.TodoList
	for _, todo := range todoLists {
		todoListsRes = append(todoListsRes, domain.MapTodoListToGrpcTodoList(&todo))
	}
	return &pb.TodoListListResponse{Data: todoListsRes, Error: domain.MapDomainGrpcError(err)}, nil
}
