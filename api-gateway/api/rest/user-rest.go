package rest

import (
	"context"
	"encoding/json"
	"g2/api-gateway/domain"
	"g2/api-gateway/variables"
	pb "g2/proto/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
	"time"
)

// commands
func (r *RestHandler) InsertUser(c *gin.Context) {
	//get request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	//request body validation
	user := &domain.Users{}
	err = json.Unmarshal(body, user)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	err = r.Msg.PublishMessage(
		body,
		variables.InsertUserQueueName,
		variables.InsertUserBindingKey,
	)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, bson.M{"error": nil, "data": "will be insert"})
}
func (r *RestHandler) EditUser(c *gin.Context) {
	id := c.Param("user-id")
	if id == "" {
		_err := domain.SetError(variables.InvalidationErr, "id is empty")
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	//check auth
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}
	if claim.UserId != id {
		_err := domain.SetError(variables.AccessErr, "can't access to another account")
		c.JSON(http.StatusForbidden, bson.M{"error": _err, "data": nil})
		return
	}

	//get request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	//request body validation
	user := &domain.Users{}
	err = json.Unmarshal(body, user)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	user.Id = id

	//ready new user for send message
	newBody, err := json.Marshal(user)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	err = r.Msg.PublishMessage(
		newBody,
		variables.EditUserQueueName,
		variables.EditUserBindingKey,
	)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, bson.M{"error": nil, "data": "will be edit"})
}
func (r *RestHandler) RemoveUser(c *gin.Context) {
	id := c.Param("user-id")
	if id == "" {
		_err := domain.SetError(variables.InvalidationErr, "id is empty")
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	//check auth
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}
	if claim.UserId != id {
		_err := domain.SetError(variables.AccessErr, "can't access to another account")
		c.JSON(http.StatusForbidden, bson.M{"error": _err, "data": nil})
		return
	}

	//prepare body for message
	user := &domain.Users{}
	user.Id = id
	body, err := json.Marshal(user)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, "id is empty")
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	err = r.Msg.PublishMessage(
		body,
		variables.RemoveUserQueueName,
		variables.RemoveUserBindingKey,
	)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, bson.M{"error": nil, "data": "will be remove"})
}

// queries
func (r *RestHandler) GetUserByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id := c.Param("user-id")

	//check auth
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}
	if claim.UserId != id {
		_err := domain.SetError(variables.AccessErr, "can't access to another account")
		c.JSON(http.StatusForbidden, bson.M{"error": _err, "data": nil})
		return
	}

	res, _ := r.Grpc.UserClient.GetUserByID(ctx, &pb.IDRequest{Id: id})
	c.JSON(http.StatusOK, res)
}
func (r *RestHandler) GetUserList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	skip, limit := domain.GetSkipLimitFromQuery(c)
	res, _ := r.Grpc.UserClient.GetUserList(ctx, &pb.SkipLimitRequest{Skip: skip, Limit: limit})
	c.JSON(http.StatusOK, res)
}
