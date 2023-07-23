package domain

import "g2/user/variables"

type domainService struct {
	Repo RepositoryInterface
}

func Service(repo RepositoryInterface) *domainService {
	return &domainService{Repo: repo}
}

// users
func (d domainService) InsertUserService(user *Users) (string, *Errors) {
	//check email not repeated
	_user, _ := d.Repo.GetUserByEmailRepository(user.Email)
	if _user != nil {
		return "", SetError(variables.EmailExistErr, "")
	}

	user.Id = GenerateID()
	user.Password = HashString(user.Password)
	user.Status = variables.ActiveStatus

	return user.Id, d.Repo.InsertUserRepository(user)
}
func (d domainService) EditUserService(user *Users) *Errors {
	return d.Repo.EditUserRepository(user)
}
func (d domainService) RemoveUserService(id string) *Errors {
	return d.Repo.RemoveUserRepository(id)
}
func (d domainService) GetUserByIDService(id string) (*Users, *Errors) {
	return d.Repo.GetUserByIDRepository(id)
}
func (d domainService) GetUserListService() ([]Users, *Errors) {
	return d.Repo.GetUserListRepository()
}
