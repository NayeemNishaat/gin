package model

import (
	"errors"
	"gin/lib"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
}

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) BeforeSave() error {
	// Part: Hash the Pass
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	// Part: Trim Username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidateLogin(username string, password string) (string, error) {
	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := lib.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func GetUserByID(userId uint) (User, error) {
	var u User

	if err := DB.First(&u, userId).Error; err != nil {
		return u, errors.New("User not found")
	}

	// u.PrepareGive()

	return u, nil
}
