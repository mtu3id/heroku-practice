package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := gin.New()
	r.GET("/ping", PingHandler)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Run server failed: %v", err)
	}
}

// PingHandler handle ping request.
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
