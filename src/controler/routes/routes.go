package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victormelos/curso-youtube/src/controler"
	"github.com/victormelos/curso-youtube/src/middleware"
)

func InitRoutes(r *gin.RouterGroup) {
	// Rotas públicas
	r.POST("/createUser", controler.CreateUser)

	// Rotas protegidas
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/getUserById", controler.FindUserById)
		protected.GET("/getUserByEmail", controler.FindUserByEmail)
		protected.GET("/getAllUsers", controler.FindAllUsers)
		protected.PUT("/updateUser", controler.UpdateUser)
		protected.DELETE("/deleteUser", controler.DeleteUser)
	}
}
