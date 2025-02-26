package controler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victormelos/curso-youtube/src/configuration/validation"
	"github.com/victormelos/curso-youtube/src/controler/model/request"
	"github.com/victormelos/curso-youtube/src/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var (
	logger     *zap.Logger
	UserDomain model.UserDomainInterface
)

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()

}

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying marshal object")
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("Request to create user", zap.Any("user", userRequest))

	domain := model.NewUserDomain(
		userRequest.Password,
		userRequest.Email,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := request.UserResponse{
		Name:  domain.GetName(),
		Email: domain.GetEmail(),
		Age:   domain.GetAge(),
	}

	c.JSON(http.StatusCreated, response)
}

func EncoderConfigyptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
