package user

import (
	db "go-web/repostory/user"
)

func GetUser(id int32) (*db.User, error) {
	return db.GetUser(id)
}

func GetUsers(name string) ([]*db.User, error) {
	return db.GetUsers(name)
}
