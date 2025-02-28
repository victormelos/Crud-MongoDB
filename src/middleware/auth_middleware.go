package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	jwtConfig "github.com/victormelos/curso-youtube/src/configuration/jwt"
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
)

const (
	userKey = "currentUser"
)

type CurrentUser struct {
	ID      string
	Name    string
	Email   string
	IsAdmin bool
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Iniciando middleware de autenticação")

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			err := rest_err.NewUnauthorizedError("Token não fornecido")
			c.JSON(err.Code, err)
			c.Abort()
			return
		}

		// O token deve vir no formato: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			err := rest_err.NewUnauthorizedError("Token mal formatado")
			c.JSON(err.Code, err)
			c.Abort()
			return
		}

		// Validar o token
		claims, err := jwtConfig.ValidateToken(parts[1])
		if err != nil {
			c.JSON(err.Code, err)
			c.Abort()
			return
		}

		// Criar o usuário atual com as informações do token
		currentUser := &CurrentUser{
			ID:      claims.ID,
			Name:    claims.Name,
			Email:   claims.Email,
			IsAdmin: claims.IsAdmin,
		}

		// Armazenar o usuário no contexto
		c.Set(userKey, currentUser)
		logger.Info("Usuário autenticado com sucesso")
		c.Next()
	}
}

func IsAuthenticated(c *gin.Context) bool {
	_, exists := c.Get(userKey)
	return exists
}

func GetCurrentUser(c *gin.Context) *CurrentUser {
	value, exists := c.Get(userKey)
	if !exists {
		return nil
	}
	return value.(*CurrentUser)
}
