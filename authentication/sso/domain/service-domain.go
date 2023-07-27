package domain

import "g2/auth/sso/variables"

type DomainService struct {
	Repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *DomainService {
	return &DomainService{Repo: repo}
}

func (d *DomainService) CheckSSOTokenValidationService(ssoToken string) (*IdpClaim, *Errors) {
	idpToken, err := d.Repo.GetSSOTokenRedis(ssoToken)
	if err != nil {
		return nil, err
	}

	// check idp validation
	validation := jwtTokenValidation(idpToken)
	if !validation {
		return nil, SetError(variables.TokenExpiredErr, "")
	}

	// get token claim
	idpClaim, err := getUserInfoFromToken(idpToken)
	if err != nil {
		return nil, err
	}
	return idpClaim, nil
}
func (d *DomainService) InsertSSOTokenService(idpToken string) (string, *Errors) {
	ssoToken := GenerateID()
	err := d.Repo.SetSSOTokenRedis(ssoToken, idpToken)
	return ssoToken, err
}
