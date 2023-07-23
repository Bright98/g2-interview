package domain

type ServiceInterface interface {
	InsertUserService(user *Users) (string, *Errors)
	EditUserService(user *Users) *Errors
	RemoveUserService(id string) *Errors
	GetUserByIDService(id string) (*Users, *Errors)
	GetUserListService() ([]Users, *Errors)
}
