package controler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mongoClient "github.com/victormelos/curso-youtube/src/configuration/database/mongodb"
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/validation"
	"github.com/victormelos/curso-youtube/src/controler/model/request"
	"github.com/victormelos/curso-youtube/src/domain/user"
	"github.com/victormelos/curso-youtube/src/model/repository/mongodb"
	"github.com/victormelos/curso-youtube/src/model/service"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying marshal object", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("Request to create user")

	userDomain := &user.UserDomain{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Age:      userRequest.Age,
	}

	repository := mongodb.NewUserRepository(mongoClient.MongoDBClient)
	domainService := service.NewUserDomainService(repository)

	result, err := domainService.Create(userDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, request.UserResponse{
		Name:  result.Name,
		Email: result.Email,
		Age:   result.Age,
	})
}
