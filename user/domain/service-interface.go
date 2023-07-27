package domain

type ServiceInterface interface {
	InsertUserService(user *Users) (string, *Errors)
	EditUserService(user *Users) *Errors
	RemoveUserService(id string) *Errors
	GetUserByIDService(id string) (*Users, *Errors)
	GetUserListService(skip, limit int64) ([]Users, *Errors)
	GetUserIDByLoginInfo(email, password string) (string, *Errors)
}
