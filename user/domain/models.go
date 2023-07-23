package domain

type Users struct {
	Id       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Status   int8   `json:"status" bson:"status"` //1: active, -1: removed
}

// error
type Errors struct {
	Key   string `json:"key" bson:"key"`
	Error string `json:"error" bson:"error"`
}
