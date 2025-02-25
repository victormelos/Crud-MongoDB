package request

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100,containsany=!@#$%*&"`
	Age      int    `json:"age" binding:"required,min=1,max=130"`
}

func (user *UserRequest) Validate() error {
	return validate.Struct(user)
}
