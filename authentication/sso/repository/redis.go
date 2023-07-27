package repository

import (
	"context"
	"g2/auth/sso/domain"
	"g2/auth/sso/variables"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	RedisCtx    context.Context
)

func (r Repository) SetSSOTokenRedis(ssoToken, idpToken string) *domain.Errors {
	res := RedisClient.Set(RedisCtx, ssoToken, idpToken, redis.KeepTTL)
	_, err := res.Result()
	if err != nil {
		return domain.SetError(variables.CantInsertErr, err.Error())
	}
	return nil
}
func (r Repository) GetSSOTokenRedis(ssoToken string) (string, *domain.Errors) {
	res := RedisClient.Get(RedisCtx, ssoToken)
	value, err := res.Result()
	if err != nil {
		return "", domain.SetError(variables.NotFoundErr, err.Error())
	}
	return value, nil
}
func (r Repository) RemoveUserTokenRedis(ssoToken string) *domain.Errors {
	res := RedisClient.Del(RedisCtx, ssoToken)
	result, err := res.Result()
	if (result == 0) || (err != nil) {
		return domain.SetError(variables.CantRemoveErr, "")
	}
	return nil
}
