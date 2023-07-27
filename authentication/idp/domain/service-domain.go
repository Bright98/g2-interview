package domain

import (
	"context"
	"g2/auth/idp/send/grpc"
	"time"
)

type DomainService struct {
	Api *grpc.GrpcClient
}

func NewService(grpcClient *grpc.GrpcClient) *DomainService {
	return &DomainService{Api: grpcClient}
}

func (d *DomainService) LoginService(loginInfo *LoginInfo) (string, *Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//	send email, password to user service and get user id to generate token
	response, _ := d.Api.UserClient.GetUserIDByLoginInfo(ctx, MapLoginInfoToUserGrpc(loginInfo))

	//	generate idp token
	idpClaim := &IdpClaim{}
	idpClaim.UserID = response.Data
	token, err := CreateJwtToken(idpClaim)
	if err != nil {
		return "", err
	}

	return token, nil
}
