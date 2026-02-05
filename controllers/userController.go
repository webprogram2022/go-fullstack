package controllers

import (
	"go-fullstack/config"
	"go-fullstack/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	rows, _ := config.DB.Query("SELECT id, name, email FROM users")

	var users []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"users": users,
	})
}

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	config.DB.Exec(
		"INSERT INTO users(name,email,password) VALUES(?,?,?)",
		name, email, password,
	)

	c.Redirect(http.StatusFound, "/dashboard")
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	config.DB.Exec("DELETE FROM users WHERE id=?", id)

	c.Redirect(http.StatusFound, "/dashboard")
}
