package domain

type TodoLists struct {
	Id          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	UserID      string `json:"user_id" bson:"user_id"`
	Status      int8   `json:"status" bson:"status"`
}
type TodoItems struct {
	Id         string `json:"id" bson:"_id"`
	TodoListID string `json:"todo_list_id" bson:"todo_list_id"`
	Title      string `json:"title" bson:"title"`
	Priority   int64  `json:"priority" bson:"priority"`
	UserID     string `json:"user_id" bson:"user_id"`
	Status     int8   `json:"status" bson:"status"`
}

// error
type Errors struct {
	Key   string `json:"key" bson:"key"`
	Error string `json:"error" bson:"error"`
}
