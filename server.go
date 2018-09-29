package main

import (
	"github.com/gin-gonic/gin"
	. "tyrannosaurs/util"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		Log.Info("A group of walrus emerges from the ocean")
		Log.Error("this is error")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
