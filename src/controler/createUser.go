package controler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/controler/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("invalid json body, error=%s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}
