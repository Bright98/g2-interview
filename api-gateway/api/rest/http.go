package rest

import (
	"context"
	"g2/api-gateway/messaging/events"
	"g2/api-gateway/send/grpc"
	ssoGrpc "g2/proto/sso"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type RestHandler struct {
	Grpc *grpc.GrpcClient
	Msg  *events.RabbitConsumer
}

func NewRestApi(grpcClient *grpc.GrpcClient, message *events.RabbitConsumer) *RestHandler {
	return &RestHandler{Grpc: grpcClient, Msg: message}
}

func (r *RestHandler) CheckAuth(c *gin.Context, ctx context.Context) (bool, *ssoGrpc.IdpClaim) {
	ssoToken := c.GetHeader("Authorization")
	validity, _ := r.Grpc.SSOClient.CheckSSOValidation(ctx, &ssoGrpc.TokenRequest{Token: ssoToken})
	if validity.Error != nil {
		c.JSON(http.StatusUnauthorized, bson.M{"data": nil, "error": validity.Error})
		return false, nil
	}
	return true, validity.Data
}
