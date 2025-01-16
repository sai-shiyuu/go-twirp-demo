package login

import (
	"context"
	"fmt"
	redis "go-web/common/cache"
	utils "go-web/common/utils"
	udb "go-web/repostory/user"

	"github.com/google/uuid"
)

func Login(Id int64, password string) (token string, err error) {
	user, err := udb.GetUser(Id)
	if err != nil {
		return
	}
	sha := utils.GenerateSHA256Hash(password)
	if sha != user.Password {
		return "", nil
	}
	token = uuid.NewString()
	r := redis.GetRedisClient()
	key := fmt.Sprintf("login:token:%d", user.Id)
	r.Del(context.Background(), key)
	r.Set(context.Background(), key, token, 0)
	return token, nil
}

func Register(name string, password string) (num int64, err error) {
	sha := utils.GenerateSHA256Hash(password)
	u := udb.User{
		Name:     name,
		Password: sha,
	}
	return udb.AddUser(u)
}
