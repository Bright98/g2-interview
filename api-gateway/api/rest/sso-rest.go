package rest

import (
	"context"
	"encoding/json"
	"g2/api-gateway/domain"
	"g2/api-gateway/variables"
	ssoProto "g2/proto/sso"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
	"time"
)

func (r *RestHandler) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//get request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	//request body validation
	loginInfo := &domain.LoginInfo{}
	err = json.Unmarshal(body, loginInfo)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	//create idp token
	idpRes, _ := r.Grpc.IdpClient.Login(ctx, domain.MapIdpGrpcToLoginInfo(loginInfo))

	//create sso Token
	idpToken := &ssoProto.TokenRequest{Token: idpRes.Data}
	ssoRes, _ := r.Grpc.SSOClient.InsertSSOToken(ctx, idpToken)

	c.JSON(http.StatusOK, ssoRes)
}
