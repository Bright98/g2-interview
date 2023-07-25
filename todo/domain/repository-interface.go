package domain

type RepositoryInterface interface {
	InsertTodoListRepository(todoList *TodoLists) *Errors
	EditTodoListRepository(todoList *TodoLists) *Errors
	RemoveTodoListRepository(id string) *Errors
	GetTodoListByIDRepository(id string) (*TodoLists, *Errors)
	GetTodoListListRepository(skip, limit int64) ([]TodoLists, *Errors)
	InsertTodoItemRepository(todoItem *TodoItems) *Errors
	EditTodoItemRepository(todoItem *TodoItems) *Errors
	RemoveTodoItemRepository(id string) *Errors
	GetTodoItemByIDRepository(id string) (*TodoItems, *Errors)
	GetTodoItemListRepository(todoListID string) ([]TodoItems, *Errors)
}
