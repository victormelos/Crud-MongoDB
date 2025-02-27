package controler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mongoClient "github.com/victormelos/curso-youtube/src/configuration/database/mongodb"
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
	"github.com/victormelos/curso-youtube/src/model/repository/mongodb"
	"github.com/victormelos/curso-youtube/src/model/service"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required"`
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
	}

	repository := mongodb.NewUserRepository(mongoClient.MongoDBClient)
	domainService := service.NewUserDomainService(repository)

	result, restErr := domainService.Create(userDomain)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    result.ID,
		"name":  result.Name,
		"email": result.Email,
		"age":   result.Age,
	})
}
