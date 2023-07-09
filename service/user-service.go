package service

import (
	"gin/model"
)

func Register(uname string, pass string) (*model.User, error) {
	u := model.User{}

	u.Username = uname
	u.Password = pass

	return u.SaveUser()
}
