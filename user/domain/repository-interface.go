package domain

type RepositoryInterface interface {
	InsertUserRepository(user *Users) *Errors
	EditUserRepository(user *Users) *Errors
	RemoveUserRepository(id string) *Errors
	GetUserByIDRepository(id string) (*Users, *Errors)
	GetUserListRepository(skip, limit int64) ([]Users, *Errors)
	GetUserByEmailPasswordRepository(email, password string) (*Users, *Errors)
	GetUserByEmailRepository(email string) (*Users, *Errors)
}
