package main

import (
	"net/http"
	"os"
	"son514/kyc-integration/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func checkSession(kyc *models.KYC) gin.HandlerFunc {
	return func(c *gin.Context) {
		decision, err := kyc.CheckSession(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, decision)
	}
}

func main() {
	godotenv.Load()

	kyc := models.NewKYC(os.Getenv("API_KEY"))

	router := gin.Default()
	router.POST("/kyc/sessions", createSession(kyc))
	router.POST("/kyc/sessions/:id", checkSession(kyc))
	router.Static("/web", "./web")
	router.Run()
}
