package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("demo")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong!!!"})
	})
	r.Run(":6681")
}
