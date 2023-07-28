package grpc

import (
	"context"
	"g2/auth/sso/domain"
	pb "g2/proto/sso"
)

func (s GrpcServer) CheckSSOValidation(ctx context.Context, in *pb.TokenRequest) (*pb.SSOValidationResponse, error) {
	//in: sso token
	ssoValidation, err := s.domain.CheckSSOTokenValidationService(in.Token)
	grpcErr := domain.MapDomainGrpcError(err)
	grpcData := domain.MapSSOValidationToGrpc(ssoValidation)
	return &pb.SSOValidationResponse{Data: grpcData, Error: grpcErr}, nil

}
func (s GrpcServer) InsertSSOToken(ctx context.Context, in *pb.TokenRequest) (*pb.TokenResponse, error) {
	//in: idp token
	ssoToken, err := s.domain.InsertSSOTokenService(in.Token)
	return &pb.TokenResponse{Data: ssoToken, Error: domain.MapDomainGrpcError(err)}, nil
}
