package controler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victormelos/curso-youtube/src/configuration/validation"
	"github.com/victormelos/curso-youtube/src/controler/model/request"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
}

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying marshal object", zap.Error(err))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("Request to create user", zap.Any("user", userRequest))

	response := request.UserResponse{
		Id:    "teste",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusOK, response)
}
