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

func Login(uname string, pass string) (string, error) {
	u := model.User{}

	u.Username = uname
	u.Password = pass

	return model.ValidateLogin(u.Username, u.Password)
}
