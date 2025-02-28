package user

import (
	"time"

	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository user.UserRepositoryInterface
}

// NewUserService cria uma nova instância do serviço de usuário
func NewUserService(repository user.UserRepositoryInterface) user.UserServiceInterface {
	return &userService{
		repository: repository,
	}
}

func (s *userService) Create(userDomain *user.UserDomain) (*user.UserDomain, *rest_err.RestErr) {
	if err := s.ValidatePassword(userDomain.Password); err != nil {
		return nil, rest_err.NewBadRequestError("Senha inválida")
	}

	userDomain.CreatedAt = time.Now()
	userDomain.UpdatedAt = time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao criptografar senha")
	}
	userDomain.Password = string(hashedPassword)

	userCreated, err := s.repository.Create(userDomain)
	if err != nil {
		return nil, err.(*rest_err.RestErr)
	}

	return userCreated, nil
}

func (s *userService) FindByEmail(email string) (*user.UserDomain, *rest_err.RestErr) {
	return s.repository.FindByEmail(email)
}

func (s *userService) FindByID(id string) (*user.UserDomain, *rest_err.RestErr) {
	return s.repository.FindByID(id)
}

func (s *userService) Update(id string, userDomain *user.UserDomain) *rest_err.RestErr {
	userDomain.UpdatedAt = time.Now()
	return s.repository.Update(id, userDomain)
}

func (s *userService) Delete(id string) *rest_err.RestErr {
	return s.repository.Delete(id)
}

func (s *userService) ValidatePassword(password string) error {
	if len(password) < 6 || len(password) > 100 {
		return rest_err.NewBadRequestError("A senha deve ter entre 6 e 100 caracteres")
	}
	return nil
}

func (s *userService) FindAll() ([]*user.UserDomain, *rest_err.RestErr) {
	return s.repository.FindAll()
}
