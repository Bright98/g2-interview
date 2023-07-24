package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func LoadEnvFile() error {
	return godotenv.Load(".env")
}
func GetServerPort() string {
	return ":" + os.Getenv("PORT")
}
func GetSkipLimitFromQuery(c *gin.Context) (int64, int64) {
	skip, ok := c.GetQuery("skip")
	if !ok {
		skip = "1"
	}
	limit, ok := c.GetQuery("limit")
	if !ok {
		limit = "10"
	}
	_skip, err := strconv.ParseInt(skip, 10, 64)
	if err != nil {
		_skip = 0
	}
	_limit, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		_limit = 0
	}

	return _skip, _limit
}
