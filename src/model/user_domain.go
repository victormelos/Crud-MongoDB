package model

import (
	"errors"

	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

func NewUserDomain(password, email, name string, age int) *UserDomain {
	return &UserDomain{
		Password: password,
		Email:    email,
		Name:     name,
		Age:      age,
	}
}

type UserDomain struct {
	Password string
	Email    string
	Name     string
	Age      int
}

func (user *UserDomain) EncoderConfigyptPassword() error {
	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	return nil
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser() *rest_err.RestErr
	FindUser() *rest_err.RestErr
	DeleteUser() *rest_err.RestErr
}
