package domain

type ServiceInterface interface {
	CheckSSOTokenValidationService(ssoToken string) (*IdpClaim, *Errors)
	InsertSSOTokenService(idpToken string) (string, *Errors)
}
