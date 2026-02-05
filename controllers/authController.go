package controllers

import (
	"go-fullstack/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var id int
	err := config.DB.QueryRow(
		"SELECT id FROM users WHERE email=? AND password=?",
		email, password,
	).Scan(&id)

	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"error": "Login gagal",
		})
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}

func Logout(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
