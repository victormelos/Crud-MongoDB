package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/victormelos/curso-youtube/src/configuration/database/mongodb"
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/controler/routes"
)

func main() {
	logger.Info("Starting application")

	// Carrega o .env, mas não falha se não existir
	godotenv.Load()

	// Inicializa a conexão com o MongoDB
	_, err := mongodb.NewMongoDBConnection()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
