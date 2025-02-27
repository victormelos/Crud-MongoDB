package controler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mongoClient "github.com/victormelos/curso-youtube/src/configuration/database/mongodb"
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/model/repository/mongodb"
	"github.com/victormelos/curso-youtube/src/model/service"
)

func DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller")
	userId := c.Query("userId")
	if userId == "" {
		errRest := rest_err.NewBadRequestError("userId é obrigatório")
		c.JSON(errRest.Code, errRest)
		return
	}

	repository := mongodb.NewUserRepository(mongoClient.MongoDBClient)
	domainService := service.NewUserDomainService(repository)

	if err := domainService.Delete(userId); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusNoContent)
}
