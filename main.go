package main

import (
	"github.com/gin-gonic/gin"
)

var sslkey string = "./certs/server.key"
var sslcert string = "./certs/server.crt"

func main() {
	r := gin.Default()

	authorized := r.Group("/api")
	authorized.GET("/ping", ping)

	//r.Use(static.Serve("/", static.LocalFile("./views", true)))
	r.RunTLS(":8080", sslcert, sslkey)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
