package main

import (
	"log"
	"os"

	"go-fullstack/config"
	"go-fullstack/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not loaded")
	}

	r := gin.Default()

	r.LoadHTMLGlob("views/*.html")
	r.Static("/public", "./public")

	config.ConnectDB()

	routes.SetupRoutes(r)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
