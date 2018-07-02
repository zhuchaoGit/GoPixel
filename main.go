package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuchao/GoPixel/fileOrg"
)

var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.POST("/upload", fileOrg.Upload)
	r.Static("/asserts", "asserts")
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
