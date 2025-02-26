package service

import (
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser() *rest_err.RestErr
	FindUser() *rest_err.RestErr
	DeleteUser() *rest_err.RestErr
	GetPassword() string
	GetEmail() string
	GetName() string
	GetAge() int
	EncryptPassword() error
}

type UserDomainServiceInterface interface {
	UserDomainInterface
}

type UserDomainService struct {
	UserDomainInterface
}

func NewUserDomainService(domain UserDomainInterface) *UserDomainService {
	return &UserDomainService{
		UserDomainInterface: domain,
	}
}

func NewUserDomain(password, email, name string, age int) UserDomainInterface {
	return &userDomain{
		password: password,
		email:    email,
		name:     name,
		age:      age,
	}
}

type userDomain struct {
	password string
	email    string
	name     string
	age      int
}

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	return nil
}

func (ud *userDomain) UpdateUser() *rest_err.RestErr {
	return nil
}

func (ud *userDomain) FindUser() *rest_err.RestErr {
	return nil
}

func (ud *userDomain) DeleteUser() *rest_err.RestErr {
	return nil
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ud.password = string(hash)
	return nil
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int {
	return ud.age
}
