package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	hostname, _ := os.Hostname()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, hostname+": pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, hostname+": no route")
	})
	err := r.Run(":9001")
	if err != nil {
		log.Fatalf("r.Run err: %v", err)
	}
}
