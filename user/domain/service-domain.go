package domain

import (
	"fmt"
	"g2/user/variables"
)

type DomainService struct {
	Repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *DomainService {
	return &DomainService{Repo: repo}
}

// users
func (d *DomainService) InsertUserService(user *Users) (string, *Errors) {
	//check email not repeated
	_user, _ := d.Repo.GetUserByEmailRepository(user.Email)
	if _user != nil {
		return "", SetError(variables.EmailExistErr, "")
	}

	user.Id = GenerateID()
	user.Password = HashString(user.Password)
	user.Status = variables.ActiveStatus

	err := d.Repo.InsertUserRepository(user)
	if err != nil {
		fmt.Println("error: ", err)
		return "", err
	}
	return user.Id, nil
}
func (d *DomainService) EditUserService(user *Users) *Errors {
	user.Password = HashString(user.Password)
	user.Status = variables.ActiveStatus

	err := d.Repo.EditUserRepository(user)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}
	return nil
}
func (d *DomainService) RemoveUserService(id string) *Errors {
	err := d.Repo.RemoveUserRepository(id)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}
	return nil
}
func (d *DomainService) GetUserByIDService(id string) (*Users, *Errors) {
	return d.Repo.GetUserByIDRepository(id)
}
func (d *DomainService) GetUserListService(skip, limit int64) ([]Users, *Errors) {
	users, err := d.Repo.GetUserListRepository(skip, limit)
	return users, err
}
