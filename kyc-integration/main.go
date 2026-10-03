package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func main() {
	router := gin.Default()

	router.GET("/someGet", getting)
	router.Run()
}
