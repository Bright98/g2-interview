package domain

type ServiceInterface interface {
	InsertTodoListService(todoList *TodoLists) (string, *Errors)
	EditTodoListService(todoList *TodoLists) *Errors
	RemoveTodoListService(id, userID string) *Errors
	GetTodoListByIDService(id, userID string) (*TodoLists, *Errors)
	GetTodoListListService(userID string, skip, limit int64) ([]TodoLists, *Errors)
	InsertTodoItemService(todoItem *TodoItems) (string, *Errors)
	EditTodoItemService(todoItem *TodoItems) *Errors
	RemoveTodoItemService(id, userID string) *Errors
	GetTodoItemByIDService(id, userID string) (*TodoItems, *Errors)
	GetTodoItemListService(todoListID, userID string) ([]TodoItems, *Errors)
}
