package domain

type RepositoryInterface interface {
	SetSSOTokenRedis(ssoToken, idpToken string) *Errors
	GetSSOTokenRedis(ssoToken string) (string, *Errors)
	RemoveUserTokenRedis(ssoToken string) *Errors
}
