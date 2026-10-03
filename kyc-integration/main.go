package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"son514/kyc-integration/models"
)

func createSession(kyc *models.KYC) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreateSessionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		session, err := kyc.CreateSession(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, session)
	}
}

func main() {
	godotenv.Load()

	kyc := models.NewKYC(os.Getenv("API_KEY"))

	router := gin.Default()
	router.POST("/kyc/sessions", createSession(kyc))
	router.Static("/web", "./web")
	router.Run()
}
