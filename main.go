package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	/*authorized := r.Group("/api", gin.BasicAuth(gin.Accounts{
		"user1": "testdigibee",
		"user2": "testaz",
	}))*/

	authorized := r.Group("/api")
	authorized.GET("/ping", ping)

	//r.Use(static.Serve("/", static.LocalFile("./views", true)))
	r.Run()
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
