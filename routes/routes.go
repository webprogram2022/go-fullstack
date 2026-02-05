package routes

import (
	"go-fullstack/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", controllers.ShowLogin)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	r.GET("/dashboard", controllers.Dashboard)
	r.POST("/add-user", controllers.AddUser)
	r.GET("/delete/:id", controllers.DeleteUser)

}
