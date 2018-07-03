package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuchao/GoPixel/fileOrg"
	"github.com/zhuchao/GoPixel/index"
)

var DB = make(map[string]string)
var indexHtmlPtr *[]byte

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/index", func(c *gin.Context) {
		c.Writer.Write(*indexHtmlPtr)
		c.Writer.Flush()
	})
	r.Static("/asserts", "asserts")
	r.POST("/upload", fileOrg.Upload)
	return r
}

func main() {
	indexHtmlPtr = index.Init()
	r := setupRouter()
	go index.RenderTask(&indexHtmlPtr)
	r.Run(":8080")
}
