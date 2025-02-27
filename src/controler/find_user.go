package controler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mongoClient "github.com/victormelos/curso-youtube/src/configuration/database/mongodb"
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/controler/model/request"
	"github.com/victormelos/curso-youtube/src/model/repository/mongodb"
	"github.com/victormelos/curso-youtube/src/model/service"
)

func FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById controller")
	userId := c.Query("userId")
	if userId == "" {
		errRest := rest_err.NewBadRequestError("userId é obrigatório")
		c.JSON(errRest.Code, errRest)
		return
	}

	repository := mongodb.NewUserRepository(mongoClient.MongoDBClient)
	domainService := service.NewUserDomainService(repository)

	userDomain, err := domainService.FindByID(userId)
	if err != nil {
		logger.Error("Error trying to call FindByID service", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, request.UserResponse{
		Name:  userDomain.Name,
		Email: userDomain.Email,
		Age:   userDomain.Age,
	})
}

func FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller")
	email := c.Query("email")
	if email == "" {
		errRest := rest_err.NewBadRequestError("email é obrigatório")
		c.JSON(errRest.Code, errRest)
		return
	}

	repository := mongodb.NewUserRepository(mongoClient.MongoDBClient)
	domainService := service.NewUserDomainService(repository)

	userDomain, err := domainService.FindByEmail(email)
	if err != nil {
		logger.Error("Error trying to call FindByEmail service", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, request.UserResponse{
		Name:  userDomain.Name,
		Email: userDomain.Email,
		Age:   userDomain.Age,
	})
}
