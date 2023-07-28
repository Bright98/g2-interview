package domain

type RepositoryInterface interface {
	InsertTodoListRepository(todoList *TodoLists) *Errors
	EditTodoListRepository(todoList *TodoLists) *Errors
	RemoveTodoListRepository(id, userID string) *Errors
	GetTodoListByIDRepository(id, userID string) (*TodoLists, *Errors)
	GetTodoListListRepository(userID string, skip, limit int64) ([]TodoLists, *Errors)
	InsertTodoItemRepository(todoItem *TodoItems) *Errors
	EditTodoItemRepository(todoItem *TodoItems) *Errors
	RemoveTodoItemRepository(id, userID string) *Errors
	GetTodoItemByIDRepository(id, userID string) (*TodoItems, *Errors)
	GetTodoItemListRepository(todoListID, userID string) ([]TodoItems, *Errors)
	InsertErrorLogRepository(errorLog *ErrorLogs) *Errors
}
