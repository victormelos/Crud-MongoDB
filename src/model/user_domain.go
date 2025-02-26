package model

import (
	"errors"

	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

type UserDomain struct {
	password string
	email    string
	name     string
	age      int
}

func NewUserDomain(password, email, name string, age int) *UserDomain {
	return &UserDomain{
		password: password,
		email:    email,
		name:     name,
		age:      age,
	}
}

// Getters
func (ud *UserDomain) GetEmail() string { return ud.email }
func (ud *UserDomain) GetName() string  { return ud.name }
func (ud *UserDomain) GetAge() int      { return ud.age }

// Setters com validação
func (ud *UserDomain) SetEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	ud.email = email
	return nil
}

func (ud *UserDomain) SetName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	ud.name = name
	return nil
}

func (ud *UserDomain) SetAge(age int) error {
	if age <= 0 || age > 130 {
		return errors.New("invalid age")
	}
	ud.age = age
	return nil
}

func (ud *UserDomain) EncoderConfigyptPassword() error {
	if ud.password == "" {
		return errors.New("password cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	ud.password = string(hash)
	return nil
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser() *rest_err.RestErr
	FindUser() *rest_err.RestErr
	DeleteUser() *rest_err.RestErr

	// Adicionando os métodos de acesso
	GetEmail() string
	GetName() string
	GetAge() int
	SetEmail(string) error
	SetName(string) error
	SetAge(int) error
}
