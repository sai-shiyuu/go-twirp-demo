package user

import (
	db "go-web/repostory/user"
)

func GetUser(id int64) (*db.User, error) {
	return db.GetUser(id)
}

func GetUsers(name string) ([]*db.User, error) {
	return db.GetUsers(name)
}
