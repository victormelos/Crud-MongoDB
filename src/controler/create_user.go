package controler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mongoClient "github.com/victormelos/curso-youtube/src/configuration/database/mongodb"
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/validation"
	"github.com/victormelos/curso-youtube/src/controler/model/request"
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

	domain := service.NewUserDomain(
		userRequest.Password,
		userRequest.Email,
		userRequest.Name,
		userRequest.Age,
	)

	repository := mongodb.NewUserRepository(mongoClient.MongoDBClient)
	domainService := service.NewUserDomainService(domain, repository)

	if err := domainService.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := request.UserResponse{
		Name:  domainService.GetName(),
		Email: domainService.GetEmail(),
		Age:   domainService.GetAge(),
	}

	c.JSON(http.StatusCreated, response)
}
