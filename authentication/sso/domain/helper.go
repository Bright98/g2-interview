package domain

import (
	"encoding/json"
	"g2/auth/sso/variables"
	pb "g2/proto/sso"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func LoadEnvFile() error {
	return godotenv.Load(".env")
}
func SetError(key string, err string) *Errors {
	return &Errors{Key: key, Error: err}
}
func GenerateID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
func MapDomainGrpcError(err *Errors) *pb.ErrorResponse {
	if err == nil {
		return nil
	}
	return &pb.ErrorResponse{Key: err.Key, Error: err.Error}
}
func MapSSOValidationToGrpc(validation *IdpClaim) *pb.IdpClaim {
	if validation == nil {
		return nil
	}
	//fill claim
	idpClaim := &pb.IdpClaim{}
	idpClaim.UserId = validation.UserID

	return idpClaim
}
func jwtFunc(_ *jwt.Token) (interface{}, error) {
	return []byte("access_token"), nil
}
func jwtTokenValidation(accessToken string) bool {
	token, err := jwt.Parse(accessToken, jwtFunc)
	if err != nil {
		return false
	}
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	return ok
}
func getUserInfoFromToken(accessToken string) (*IdpClaim, *Errors) {
	token, err := jwt.Parse(accessToken, jwtFunc)
	if err != nil {
		return nil, SetError(variables.TokenIsWrongErr, err.Error())
	}

	claim := make(map[string]any)
	claimByte, err := json.Marshal(token.Claims)
	if err != nil {
		return nil, SetError(variables.TokenIsWrongErr, err.Error())
	}

	err = json.Unmarshal(claimByte, &claim)
	if err != nil {
		return nil, SetError(variables.TokenIsWrongErr, err.Error())
	}

	userID, ok := claim["user_id"]
	if !ok {
		return nil, SetError(variables.TokenIsWrongErr, err.Error())
	}

	idpClaim := &IdpClaim{}
	idpClaim.UserID = userID.(string)

	return idpClaim, nil
}
