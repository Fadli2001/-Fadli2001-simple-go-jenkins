package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	log.Println("Hello World")
	route := gin.Default()
	GetPing(route)
	err := route.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Server is not running %s", err)
	}
}

func GetPing(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
}
