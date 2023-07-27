package domain

type IdpClaim struct {
	UserID string `json:"user_id" bson:"user_id"`
}

// error
type Errors struct {
	Key   string `json:"key" bson:"key"`
	Error string `json:"error" bson:"error"`
}
