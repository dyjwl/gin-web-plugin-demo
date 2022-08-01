package main

import (
	"net/http"

	"github.com/dyjwl/gin-web-plugin-demo/pkg/log"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	log.InitLog()
	r.GET("/ping", func(c *gin.Context) {
		log.Info("get ping from client. will send pong from server")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
