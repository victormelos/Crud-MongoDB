package user

import (
	"time"

	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

// UserDomain representa a entidade de usuário no domínio
type UserDomain struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Age       int       `bson:"age"`
	IsAdmin   bool      `bson:"is_admin"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (ud *UserDomain) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ud.Password = string(hash)
	return nil
}

// UserRepositoryInterface define o contrato para operações de repositório
type UserRepositoryInterface interface {
	Create(user *UserDomain) (*UserDomain, *rest_err.RestErr)
	FindByEmail(email string) (*UserDomain, *rest_err.RestErr)
	FindByID(id string) (*UserDomain, *rest_err.RestErr)
	Update(id string, user *UserDomain) *rest_err.RestErr
	Delete(id string) *rest_err.RestErr
	FindAll() ([]*UserDomain, *rest_err.RestErr)
}

// UserServiceInterface define o contrato para operações de serviço
type UserServiceInterface interface {
	Create(user *UserDomain) (*UserDomain, *rest_err.RestErr)
	FindByEmail(email string) (*UserDomain, *rest_err.RestErr)
	FindByID(id string) (*UserDomain, *rest_err.RestErr)
	Update(id string, user *UserDomain) *rest_err.RestErr
	Delete(id string) *rest_err.RestErr
	FindAll() ([]*UserDomain, *rest_err.RestErr)
	ValidatePassword(password string) error
}
