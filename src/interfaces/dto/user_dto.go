package dto

// UserRequest representa a estrutura de dados para criação de usuário
type UserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100,containsany=!@#$%*&"`
	Age      int    `json:"age" binding:"required,min=1,max=130"`
}

// UserUpdateRequest representa a estrutura de dados para atualização de usuário
type UserUpdateRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=100"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,min=1,max=130"`
}

// UserResponse representa a estrutura de dados para resposta de usuário
type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
