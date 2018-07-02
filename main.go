package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuchao/GoPixel/fileOrg"
	"time"
	"fmt"
)

var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Static("/asserts", "asserts")
	r.POST("/upload", fileOrg.Upload)
	return r
}
func renderTask(r *gin.Engine) {
	ticker := time.NewTicker(5 * time.Second)
	for t := range ticker.C {
		fmt.Println(t)
	}
}
func main() {
	r := setupRouter()
	go renderTask(r)
	r.Run(":8080")
}
