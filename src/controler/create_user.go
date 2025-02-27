package controler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	mongoClient "github.com/victormelos/curso-youtube/src/configuration/database/mongodb"
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
	"github.com/victormelos/curso-youtube/src/model/repository/mongodb"
	"github.com/victormelos/curso-youtube/src/model/service"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser")

	name := c.Query("name")
	email := c.Query("email")
	password := c.Query("password")
	ageStr := c.Query("age")

	if name == "" || email == "" || password == "" || ageStr == "" {
		errRest := rest_err.NewBadRequestError("Todos os campos são obrigatórios")
		c.JSON(errRest.Code, errRest)
		return
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		errRest := rest_err.NewBadRequestError("Idade inválida")
		c.JSON(errRest.Code, errRest)
		return
	}

	userDomain := &user.UserDomain{
		Name:     name,
		Email:    email,
		Password: password,
		Age:      age,
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
