package rest

import (
	"context"
	"encoding/json"
	"g2/api-gateway/domain"
	"g2/api-gateway/variables"
	pb "g2/proto/todo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
	"time"
)

// todo list
// commands
func (r *RestHandler) InsertTodoList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//check auth
	validity, _ := r.CheckAuth(c, ctx)
	if !validity {
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
	todoList := &domain.TodoLists{}
	err = json.Unmarshal(body, todoList)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	err = r.Msg.PublishMessage(
		body,
		variables.InsertTodoListQueueName,
		variables.InsertTodoListBindingKey,
	)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, bson.M{"error": nil, "data": "will be insert"})
}
func (r *RestHandler) EditTodoList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//check auth
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}

	id := c.Param("todo-list-id")
	if id == "" {
		_err := domain.SetError(variables.InvalidationErr, "id is empty")
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
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
	todoList := &domain.TodoLists{}
	err = json.Unmarshal(body, todoList)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	todoList.Id = id
	todoList.UserID = claim.GetUserId()

	//ready new todo list for send message
	newBody, err := json.Marshal(todoList)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	err = r.Msg.PublishMessage(
		newBody,
		variables.EditTodoListQueueName,
		variables.EditTodoListBindingKey,
	)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, bson.M{"error": nil, "data": "will be edit"})
}
func (r *RestHandler) RemoveTodoList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//check auth
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}

	id := c.Param("todo-list-id")
	if id == "" {
		_err := domain.SetError(variables.InvalidationErr, "id is empty")
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	//prepare body for message
	todoList := &domain.TodoLists{}
	todoList.Id = id
	todoList.UserID = claim.GetUserId()
	body, err := json.Marshal(todoList)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, "id is empty")
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	err = r.Msg.PublishMessage(
		body,
		variables.RemoveTodoListQueueName,
		variables.RemoveTodoListBindingKey,
	)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, bson.M{"error": nil, "data": "will be remove"})
}

// queries
func (r *RestHandler) GetTodoListByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//check auth
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}

	id := c.Param("todo-list-id")
	res, _ := r.Grpc.TodoClient.GetTodoListByID(ctx, &pb.IDRequest{Id: id, UserId: claim.GetUserId()})
	c.JSON(http.StatusOK, res)
}
func (r *RestHandler) GetTodoListList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//check auth
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}

	skip, limit := domain.GetSkipLimitFromQuery(c)
	res, _ := r.Grpc.TodoClient.GetTodoListList(
		ctx,
		&pb.SkipLimitRequest{Skip: skip, Limit: limit, UserId: claim.GetUserId()},
	)
	c.JSON(http.StatusOK, res)
}

// todo items
// commands
func (r *RestHandler) InsertTodoItem(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//check auth
	validity, _ := r.CheckAuth(c, ctx)
	if !validity {
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
	todoItem := &domain.TodoItems{}
	err = json.Unmarshal(body, todoItem)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	err = r.Msg.PublishMessage(
		body,
		variables.InsertTodoItemQueueName,
		variables.InsertTodoItemBindingKey,
	)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, bson.M{"error": nil, "data": "will be insert"})
}
func (r *RestHandler) EditTodoItem(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//check auth
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}

	id := c.Param("todo-item-id")
	if id == "" {
		_err := domain.SetError(variables.InvalidationErr, "id is empty")
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
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
	todoItem := &domain.TodoItems{}
	err = json.Unmarshal(body, todoItem)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	todoItem.Id = id
	todoItem.UserID = claim.GetUserId()

	//ready new todo item for send message
	newBody, err := json.Marshal(todoItem)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	err = r.Msg.PublishMessage(
		newBody,
		variables.EditTodoItemQueueName,
		variables.EditTodoItemBindingKey,
	)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, bson.M{"error": nil, "data": "will be edit"})
}
func (r *RestHandler) RemoveTodoItem(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//check auth
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}

	id := c.Param("todo-item-id")
	if id == "" {
		_err := domain.SetError(variables.InvalidationErr, "id is empty")
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	//prepare body for message
	todoItem := &domain.TodoItems{}
	todoItem.Id = id
	todoItem.UserID = claim.GetUserId()
	body, err := json.Marshal(todoItem)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, "id is empty")
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}

	err = r.Msg.PublishMessage(
		body,
		variables.RemoveTodoItemQueueName,
		variables.RemoveTodoItemBindingKey,
	)
	if err != nil {
		_err := domain.SetError(variables.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"error": _err, "data": nil})
		return
	}
	c.JSON(http.StatusOK, bson.M{"error": nil, "data": "will be remove"})
}

// queries
func (r *RestHandler) GetTodoItemByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//check auth
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}

	id := c.Param("todo-item-id")
	res, _ := r.Grpc.TodoClient.GetTodoItemByID(ctx, &pb.IDRequest{Id: id, UserId: claim.GetUserId()})
	c.JSON(http.StatusOK, res)
}
func (r *RestHandler) GetTodoItemList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//check auth
	validity, claim := r.CheckAuth(c, ctx)
	if !validity {
		return
	}

	id := c.Param("todo-item-id")
	res, _ := r.Grpc.TodoClient.GetTodoItemList(ctx, &pb.IDRequest{Id: id, UserId: claim.GetUserId()})
	c.JSON(http.StatusOK, res)
}
