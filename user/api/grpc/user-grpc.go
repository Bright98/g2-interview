package grpc

import (
	"context"
	pb "g2/proto/user"
	"g2/user/domain"
)

// write for necessary usage (not used in api-gateway)
func (s GrpcServer) InsertUser(ctx context.Context, in *pb.User) (*pb.InsertedIDResponse, error) {
	request := &domain.Users{}
	request.Name = in.GetName()
	request.Email = in.GetEmail()
	request.Password = in.GetPassword()

	insertedID, err := s.domain.InsertUserService(request)
	response := &pb.IDRequest{Id: insertedID}
	return &pb.InsertedIDResponse{Data: response, Error: domain.MapDomainGrpcError(err)}, nil
}
func (s GrpcServer) EditUser(ctx context.Context, in *pb.User) (*pb.ErrorResponse, error) {
	request := &domain.Users{}
	request.Id = in.GetId()
	request.Name = in.GetName()
	request.Email = in.GetEmail()
	request.Password = in.GetPassword()
	request.Status = int8(in.GetStatus())

	err := s.domain.EditUserService(request)
	return &pb.ErrorResponse{Key: err.Key, Error: err.Error}, nil
}
func (s GrpcServer) RemoveUser(ctx context.Context, in *pb.IDRequest) (*pb.ErrorResponse, error) {
	err := s.domain.RemoveUserService(in.GetId())
	return &pb.ErrorResponse{Key: err.Key, Error: err.Error}, nil
}

// use in api gateway api
func (s GrpcServer) GetUserByID(ctx context.Context, in *pb.IDRequest) (*pb.UserResponse, error) {
	user, err := s.domain.GetUserByIDService(in.GetId())
	return &pb.UserResponse{Data: domain.MapUserToGrpcUser(user), Error: domain.MapDomainGrpcError(err)}, nil
}
func (s GrpcServer) GetUserList(ctx context.Context, in *pb.SkipLimitRequest) (*pb.UserListResponse, error) {
	users, err := s.domain.GetUserListService()
	var usersRes []*pb.User
	for _, user := range users {
		usersRes = append(usersRes, domain.MapUserToGrpcUser(&user))
	}
	return &pb.UserListResponse{Data: usersRes, Error: domain.MapDomainGrpcError(err)}, nil
}
