package domain

// user
type Users struct {
	Id       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Status   int8   `json:"status" bson:"status"` //1: active, -1: removed
}

// todo
// todo list
type TodoLists struct {
	Id          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Status      int8   `json:"status" bson:"status"`
}

// todo item
type TodoItems struct {
	Id         string `json:"id" bson:"_id"`
	TodoListID string `json:"todo_list_id" bson:"todo_list_id"`
	Title      string `json:"title" bson:"title"`
	Priority   int64  `json:"priority" bson:"priority"`
	Status     int8   `json:"status" bson:"status"`
}

// idp
type LoginInfo struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

// error
type Errors struct {
	Key   string `json:"key" bson:"key"`
	Error string `json:"error" bson:"error"`
}
