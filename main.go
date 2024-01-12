package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":   "sidiq pratomo",
			"occupy": "sealabs bootcamp",
		})
	})
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":    "Hello",
			"subtitle": "Belajar Golang",
		})
	})

	router.Run(":8888")
}
