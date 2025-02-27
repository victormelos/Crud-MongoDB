package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victormelos/curso-youtube/src/controler"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById", controler.FindUserById)
	r.GET("/getUserByEmail", controler.FindUserByEmail)
	r.POST("/createUser", controler.CreateUser)
	r.PUT("/updateUser", controler.UpdateUser)
	r.DELETE("/deleteUser", controler.DeleteUser)
}
