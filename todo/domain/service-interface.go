package domain

type ServiceInterface interface {
	InsertTodoListService(todoList *TodoLists) (string, *Errors)
	EditTodoListService(todoList *TodoLists) *Errors
	RemoveTodoListService(id string) *Errors
	GetTodoListByIDService(id string) (*TodoLists, *Errors)
	GetTodoListListService(skip, limit int64) ([]TodoLists, *Errors)
	InsertTodoItemService(todoItem *TodoItems) (string, *Errors)
	EditTodoItemService(todoItem *TodoItems) *Errors
	RemoveTodoItemService(id string) *Errors
	GetTodoItemByIDService(id string) (*TodoItems, *Errors)
	GetTodoItemListService(todoListID string) ([]TodoItems, *Errors)
}
