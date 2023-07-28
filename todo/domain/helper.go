package domain

import (
	pb "g2/proto/todo"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvFile() error {
	return godotenv.Load(".env")
}
func SetError(key string, err string) *Errors {
	return &Errors{Key: key, Error: err}
}
func GenerateID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
func GetServerPort() string {
	return ":" + os.Getenv("PORT")
}
func MapDomainGrpcError(err *Errors) *pb.ErrorResponse {
	if err == nil {
		return nil
	}
	return &pb.ErrorResponse{Key: err.Key, Error: err.Error}
}
func MapTodoListToGrpcTodoList(todoList *TodoLists) *pb.TodoList {
	if todoList == nil {
		return nil
	}
	todoListRes := &pb.TodoList{}
	todoListRes.Id = todoList.Id
	todoListRes.Name = todoList.Name
	todoListRes.Description = todoList.Description
	todoListRes.UserId = ""
	todoListRes.Status = int32(todoList.Status)
	return todoListRes
}
func MapTodoItemToGrpcTodoItem(todoItem *TodoItems) *pb.TodoItem {
	if todoItem == nil {
		return nil
	}
	todoItemRes := &pb.TodoItem{}
	todoItemRes.Id = todoItem.Id
	todoItemRes.TodoListId = todoItem.TodoListID
	todoItemRes.Title = todoItem.Title
	todoItemRes.Priority = todoItem.Priority
	todoItemRes.UserId = ""
	todoItemRes.Status = int32(todoItem.Status)
	return todoItemRes
}
func (d *DomainService) InsertErrorLogFunction(err *Errors, collection, desc string) {
	log := &ErrorLogs{}
	log.Key = err.Key
	log.Error = err.Error
	log.Collection = collection
	log.Description = desc
	d.Repo.InsertErrorLogRepository(log)
}
