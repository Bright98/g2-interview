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
		return "", err
	}
	return todoList.Id, nil
}
func (d *DomainService) EditTodoListService(todoList *TodoLists) *Errors {
	err := d.Repo.EditTodoListRepository(todoList)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}
	return nil
}
func (d *DomainService) RemoveTodoListService(id string) *Errors {
	err := d.Repo.RemoveTodoListRepository(id)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}
	return nil
}
func (d *DomainService) GetTodoListByIDService(id string) (*TodoLists, *Errors) {
	return d.Repo.GetTodoListByIDRepository(id)
}
func (d *DomainService) GetTodoListListService(skip, limit int64) ([]TodoLists, *Errors) {
	return d.Repo.GetTodoListListRepository(skip, limit)
}

// todo item
func (d *DomainService) InsertTodoItemService(todoItem *TodoItems) (string, *Errors) {
	todoItem.Id = GenerateID()
	todoItem.Status = variables.ActiveStatus
	err := d.Repo.InsertTodoItemRepository(todoItem)
	if err != nil {
		fmt.Println("error: ", err)
		return "", err
	}
	return todoItem.Id, nil
}
func (d *DomainService) EditTodoItemService(todoItem *TodoItems) *Errors {
	err := d.Repo.EditTodoItemRepository(todoItem)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}
	return nil
}
func (d *DomainService) RemoveTodoItemService(id string) *Errors {
	err := d.Repo.RemoveTodoItemRepository(id)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}
	return nil
}
func (d *DomainService) GetTodoItemByIDService(id string) (*TodoItems, *Errors) {
	return d.Repo.GetTodoItemByIDRepository(id)
}
func (d *DomainService) GetTodoItemListService(todoListID string) ([]TodoItems, *Errors) {
	return d.Repo.GetTodoItemListRepository(todoListID)
}
