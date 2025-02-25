package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victormelos/curso-youtube/src/controler"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controler.FindUserById)
	r.GET("/getUserByEmail/:email", controler.FindUserByEmail)
	r.POST("/createUser", controler.CreateUser)
	r.PUT("/updateUser/:userId", controler.UpdateUser)
	r.DELETE("/deleteUser/:userId", controler.DeleteUser)
}
