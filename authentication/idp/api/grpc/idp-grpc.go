package grpc

import (
	"context"
	"g2/auth/idp/domain"
	pb "g2/proto/idp"
)

func (s GrpcServer) Login(ctx context.Context, in *pb.LoginInfoRequest) (*pb.TokenResponse, error) {
	idpToken, err := s.domain.LoginService(domain.MapIdpGrpcToLoginInfo(in))
	return &pb.TokenResponse{Data: idpToken, Error: domain.MapDomainGrpcError(err)}, nil
}
