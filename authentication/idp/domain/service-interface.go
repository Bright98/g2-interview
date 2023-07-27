package domain

type ServiceInterface interface {
	LoginService(loginInfo *LoginInfo) (string, *Errors)
}
