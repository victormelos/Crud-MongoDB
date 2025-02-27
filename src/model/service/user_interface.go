package service

import (
	"time"

	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/model/domain"
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
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
}

type UserDomainServiceInterface interface {
	UserDomainInterface
}

type UserDomainService struct {
	domain.UserDomainInterface
	userRepository domain.UserRepositoryInterface
}

func NewUserDomainService(userDomain domain.UserDomainInterface, repository domain.UserRepositoryInterface) *UserDomainService {
	return &UserDomainService{
		UserDomainInterface: userDomain,
		userRepository:      repository,
	}
}

func NewUserDomain(password, email, name string, age int) UserDomainInterface {
	return &userDomain{
		Password: password,
		Email:    email,
		Name:     name,
		Age:      age,
	}
}

type userDomain struct {
	Password  string    `bson:"password"`
	Email     string    `bson:"email"`
	Name      string    `bson:"name"`
	Age       int       `bson:"age"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
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
	return ud.Password
}

func (ud *userDomain) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ud.Password = string(hash)
	return nil
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetAge() int {
	return ud.Age
}

func (ud *userDomain) SetCreatedAt(time time.Time) {
	ud.CreatedAt = time
}

func (ud *userDomain) SetUpdatedAt(time time.Time) {
	ud.UpdatedAt = time
}
