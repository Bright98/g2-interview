package domain

import (
	"fmt"
	"g2/todo/variables"
)

type DomainService struct {
	Repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *DomainService {
	return &DomainService{Repo: repo}
}

// todo list
func (d *DomainService) InsertTodoListService(todoList *TodoLists) (string, *Errors) {
	todoList.Id = GenerateID()
	todoList.Status = variables.ActiveStatus
	err := d.Repo.InsertTodoListRepository(todoList)
	if err != nil {
		fmt.Println("error: ", err)
		d.InsertErrorLogFunction(err, variables.TodoListCollection, "insert error")
		return "", err
	}
	return todoList.Id, nil
}
func (d *DomainService) EditTodoListService(todoList *TodoLists) *Errors {
	todoList.Status = variables.ActiveStatus
	err := d.Repo.EditTodoListRepository(todoList)
	if err != nil {
		fmt.Println("error: ", err)
		d.InsertErrorLogFunction(err, variables.TodoListCollection, "edit error")
		return err
	}
	return nil
}
func (d *DomainService) RemoveTodoListService(id, userID string) *Errors {
	err := d.Repo.RemoveTodoListRepository(id, userID)
	if err != nil {
		fmt.Println("error: ", err)
		d.InsertErrorLogFunction(err, variables.TodoListCollection, "remove error")
		return err
	}
	return nil
}
func (d *DomainService) GetTodoListByIDService(id, userID string) (*TodoLists, *Errors) {
	return d.Repo.GetTodoListByIDRepository(id, userID)
}
func (d *DomainService) GetTodoListListService(userID string, skip, limit int64) ([]TodoLists, *Errors) {
	return d.Repo.GetTodoListListRepository(userID, skip, limit)
}

// todo item
func (d *DomainService) InsertTodoItemService(todoItem *TodoItems) (string, *Errors) {
	//check todo list existence
	_, err := d.Repo.GetTodoListByIDRepository(todoItem.TodoListID, todoItem.UserID)
	if err != nil {
		fmt.Println("error: ", err)
		d.InsertErrorLogFunction(err, variables.TodoItemCollection, "find todo list")
		return "", err
	}

	todoItem.Id = GenerateID()
	todoItem.Status = variables.ActiveStatus
	err = d.Repo.InsertTodoItemRepository(todoItem)
	if err != nil {
		fmt.Println("error: ", err)
		d.InsertErrorLogFunction(err, variables.TodoItemCollection, "insert error")
		return "", err
	}
	return todoItem.Id, nil
}
func (d *DomainService) EditTodoItemService(todoItem *TodoItems) *Errors {
	//check todo list existence
	_, err := d.Repo.GetTodoListByIDRepository(todoItem.TodoListID, todoItem.UserID)
	if err != nil {
		fmt.Println("error: ", err)
		d.InsertErrorLogFunction(err, variables.TodoItemCollection, "find todo list")
		return err
	}

	todoItem.Status = variables.ActiveStatus
	err = d.Repo.EditTodoItemRepository(todoItem)
	if err != nil {
		fmt.Println("error: ", err)
		d.InsertErrorLogFunction(err, variables.TodoItemCollection, "edit error")
		return err
	}
	return nil
}
func (d *DomainService) RemoveTodoItemService(id, userID string) *Errors {
	err := d.Repo.RemoveTodoItemRepository(id, userID)
	if err != nil {
		fmt.Println("error: ", err)
		d.InsertErrorLogFunction(err, variables.TodoItemCollection, "remove error")
		return err
	}
	return nil
}
func (d *DomainService) GetTodoItemByIDService(id, userID string) (*TodoItems, *Errors) {
	return d.Repo.GetTodoItemByIDRepository(id, userID)
}
func (d *DomainService) GetTodoItemListService(todoListID, userID string) ([]TodoItems, *Errors) {
	return d.Repo.GetTodoItemListRepository(todoListID, userID)
}
