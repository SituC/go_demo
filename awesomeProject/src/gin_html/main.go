package main

import (
	"github.com/gin-gonic/gin"

)

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func DoLogin(c *gin.Context) {
	username := c.PostForm("username")
	// password := c.PostForm("password")
	// 取post参数可以设置默认值
	password := c.DefaultPostForm("password", "123")

	c.HTML(200, "index.html", gin.H{
		"username": username,
		"password": password,
	})
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("../html/*")
	e.GET("/login", Login)
	e.POST("/login", DoLogin)
	e.Run(":8080")
}