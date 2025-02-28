package controler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mongoClient "github.com/victormelos/curso-youtube/src/configuration/database/mongodb"
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/middleware"
	"github.com/victormelos/curso-youtube/src/model/repository/mongodb"
	"github.com/victormelos/curso-youtube/src/model/service"
	"go.uber.org/zap"
)

func DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller")

	// Verificar autenticação
	if !middleware.IsAuthenticated(c) {
		errRest := rest_err.NewUnauthorizedError("Usuário não autenticado")
		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Query("userId")
	if userId == "" {
		errRest := rest_err.NewBadRequestError("userId é obrigatório")
		c.JSON(errRest.Code, errRest)
		return
	}

	// Verificar autorização
	currentUser := middleware.GetCurrentUser(c)
	if !currentUser.IsAdmin && currentUser.ID != userId {
		errRest := rest_err.NewForbiddenError("Sem permissão para deletar este usuário")
		c.JSON(errRest.Code, errRest)
		return
	}

	repository := mongodb.NewUserRepository(mongoClient.MongoDBClient)
	domainService := service.NewUserDomainService(repository)

	// Verificar se o usuário existe antes de deletar
	_, err := domainService.FindByID(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if err := domainService.Delete(userId); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Usuário deletado com sucesso", zap.String("userId", userId))
	c.Status(http.StatusNoContent)
}
