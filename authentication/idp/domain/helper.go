package domain

import (
	"g2/auth/idp/variables"
	idpProto "g2/proto/idp"
	userProto "g2/proto/user"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

const TokenExpireTime = time.Hour * 24

func LoadEnvFile() error {
	return godotenv.Load(".env")
}
func GetServerPort() string {
	return ":" + os.Getenv("PORT")
}
func SetError(key string, err string) *Errors {
	return &Errors{Key: key, Error: err}
}
func MapDomainGrpcError(err *Errors) *idpProto.ErrorResponse {
	if err == nil {
		return nil
	}
	return &idpProto.ErrorResponse{Key: err.Key, Error: err.Error}
}
func MapLoginInfoToUserGrpc(info *LoginInfo) *userProto.LoginInfoRequest {
	if info == nil {
		return nil
	}
	infoRes := &userProto.LoginInfoRequest{}
	infoRes.Email = info.Email
	infoRes.Password = info.Password
	return infoRes
}
func MapIdpGrpcToLoginInfo(info *idpProto.LoginInfoRequest) *LoginInfo {
	if info == nil {
		return nil
	}
	loginInfo := &LoginInfo{}
	loginInfo.Email = info.Email
	loginInfo.Password = info.Password
	return loginInfo
}

// jwt auth
func tokenExpirationTime() int64 {
	return time.Now().Add(TokenExpireTime).Unix()
}
func jwtFunc(_ *jwt.Token) (interface{}, error) {
	return []byte("access_token"), nil
}
func createClaims(idpClaim *IdpClaim) jwt.MapClaims {
	accessTokenExpireTime := tokenExpirationTime()
	claims := jwt.MapClaims{}
	claims["user_id"] = idpClaim.UserID
	claims["authorized"] = true
	claims["exp"] = accessTokenExpireTime
	return claims
}
func createAccessTokenByClaim(claim jwt.Claims) (string, error) {
	signMethod := jwt.SigningMethodHS256
	token := jwt.New(signMethod)
	token.Claims = claim
	tokenKey, err := jwtFunc(token)
	signedKey, err := token.SignedString(tokenKey)
	if err != nil {
		return "", err
	}
	return signedKey, nil

}
func CreateJwtToken(idpClaim *IdpClaim) (string, *Errors) {
	claim := createClaims(idpClaim)
	token, err := createAccessTokenByClaim(claim)
	if err != nil {
		return "", SetError(variables.CantCreateTokenErr, err.Error())
	}
	return token, nil
}
