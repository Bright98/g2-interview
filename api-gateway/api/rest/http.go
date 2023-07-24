package rest

import (
	"context"
	"g2/api-gateway/domain"
	"g2/api-gateway/send/grpc"
	pb "g2/proto/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

type RestHandler struct {
	Grpc *grpc.GrpcClient
}

func NewRestApi(grpcClient *grpc.GrpcClient) *RestHandler {
	return &RestHandler{Grpc: grpcClient}
}

func (r *RestHandler) GetUserByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id := c.Param("user-id")

	res, _ := r.Grpc.UserClient.GetUserByID(ctx, &pb.IDRequest{Id: id})
	c.JSON(http.StatusOK, bson.M{"data": res})
}
func (r *RestHandler) GetUserList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	skip, limit := domain.GetSkipLimitFromQuery(c)
	res, _ := r.Grpc.UserClient.GetUserList(ctx, &pb.SkipLimitRequest{Skip: skip, Limit: limit})
	c.JSON(http.StatusOK, bson.M{"data": res})
}
