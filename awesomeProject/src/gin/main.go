package main

import (
	"github.com/gin-gonic/gin"

)

func Hello(c *gin.Context) {
	// c.String(200, "Hello World, %s", "lwl")
	c.JSON(200, gin.H{
		"name": "Hello World",
		"age": "20",
	})
}

func main() {
	e := gin.Default()
	e.GET("/hello", Hello)
	e.Run(":8080")
}