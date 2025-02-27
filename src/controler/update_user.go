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

type UpdateUserInput struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Age      *int    `json:"age,omitempty"`
	Password *string `json:"password,omitempty"`
}

func UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller")

	userId := c.Query("userId")
	if userId == "" {
		errRest := rest_err.NewBadRequestError("userId é obrigatório")
		c.JSON(errRest.Code, errRest)
		return
	}

	var userInput UpdateUserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		errRest := rest_err.NewBadRequestError("Dados inválidos no corpo da requisição")
		c.JSON(errRest.Code, errRest)
		return
	}

	repository := mongodb.NewUserRepository(mongoClient.MongoDBClient)
	domainService := service.NewUserDomainService(repository)

	// Cria um userDomain apenas com os campos que foram fornecidos
	userDomain := &user.UserDomain{}

	// Só inclui os campos que foram explicitamente enviados no JSON
	if userInput.Name != nil {
		userDomain.Name = *userInput.Name
	}
	if userInput.Email != nil {
		userDomain.Email = *userInput.Email
	}
	if userInput.Age != nil {
		userDomain.Age = *userInput.Age
	}
	if userInput.Password != nil {
		userDomain.Password = *userInput.Password
	}

	if restErr := domainService.Update(userId, userDomain); restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusNoContent)
}
