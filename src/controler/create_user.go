package controler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mongoClient "github.com/victormelos/curso-youtube/src/configuration/database/mongodb"
	jwtConfig "github.com/victormelos/curso-youtube/src/configuration/jwt"
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
	"github.com/victormelos/curso-youtube/src/model/repository/mongodb"
	"github.com/victormelos/curso-youtube/src/model/service"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
	Age      int    `json:"age" binding:"required,min=1,max=130"`
	IsAdmin  bool   `json:"is_admin"`
}

type CreateUserResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Token   string `json:"token"`
	IsAdmin bool   `json:"is_admin"`
}

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser")

	var userInput CreateUserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		errRest := rest_err.NewBadRequestError("Dados inválidos no corpo da requisição")
		c.JSON(errRest.Code, errRest)
		return
	}

	userDomain := &user.UserDomain{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
		Age:      userInput.Age,
		IsAdmin:  userInput.IsAdmin,
	}

	repository := mongodb.NewUserRepository(mongoClient.MongoDBClient)
	domainService := service.NewUserDomainService(repository)

	result, restErr := domainService.Create(userDomain)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	// Gerar token JWT
	token, err := jwtConfig.GenerateToken(
		result.ID,
		result.Name,
		result.Email,
		result.IsAdmin,
	)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, CreateUserResponse{
		ID:      result.ID,
		Name:    result.Name,
		Email:   result.Email,
		Age:     result.Age,
		Token:   token,
		IsAdmin: result.IsAdmin,
	})
}
