package main

import (
	"go-fullstack/config"
	"go-fullstack/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*.html")
	r.Static("/public", "./public")

	config.ConnectDB()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
