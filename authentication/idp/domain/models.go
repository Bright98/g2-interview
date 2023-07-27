package domain

type LoginInfo struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
type IdpClaim struct {
	UserID string `json:"user-id" bson:"user-id"`
}

// error
type Errors struct {
	Key   string `json:"key" bson:"key"`
	Error string `json:"error" bson:"error"`
}
