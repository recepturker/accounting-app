package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/recepturker/accounting-app/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var router *gin.Engine = gin.Default()
	routes.RegisterRoutes(router)

	var port string = os.Getenv("APP_PORT")
	log.Fatal(router.Run(":" + port))
}
