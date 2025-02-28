package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
)

var (
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	JWT_EXPIRES_IN = 24 * time.Hour // 24 horas
)

func init() {
	if JWT_SECRET_KEY == "" {
		JWT_SECRET_KEY = "sua_chave_secreta_aqui" // Em produção, use variável de ambiente
	}
}

type jwtCustomClaims struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func GenerateToken(id, name, email string, isAdmin bool) (string, *rest_err.RestErr) {
	claims := jwtCustomClaims{
		ID:      id,
		Name:    name,
		Email:   email,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXPIRES_IN)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", rest_err.NewInternalServerError("Erro ao gerar token")
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwtCustomClaims, *rest_err.RestErr) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedError("Token inválido")
	}

	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedError("Token inválido")
	}

	return claims, nil
}
