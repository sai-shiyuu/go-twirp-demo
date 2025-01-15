package user

import (
	"errors"
	"go-web/common/db"
	"go-web/common/log"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = log.GetLogger()
}

type User struct {
	Id    int32
	Name  string
	Ctime string
	Mtime string
}

func GetUser(id int32) (*User, error) {
	row, err := db.QueryRow("select * from t_user where id = ?", id)
	if err != nil {
		logger.Error(err)
		return nil, errors.New("get user error")
	}

	u := User{}
	err = row.Scan(&u.Id, &u.Name, &u.Ctime, &u.Mtime)
	if err != nil {
		logger.Error(err)
		return nil, errors.New("user scan error")
	}
	return &u, err
}

func GetUsers(name string) ([]*User, error) {
	rows, err := db.QueryRows("select * from t_user where name = ?", name)
	if err != nil {
		logger.Error(err)
		return nil, errors.New("get users error")
	}
	defer rows.Close()
	users := []*User{}
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.Id, &u.Name, &u.Ctime, &u.Mtime)
		if err != nil {
			logger.Error(err)
			return nil, errors.New("user scan error")
		}
		users = append(users, &u)
	}
	return users, err
}
