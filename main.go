package main

import (
	"log"
	"os"
	"saweria-webhook/handlers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get configuration from environment variables
	port := getEnv("PORT", "3000")
	secretKey := getEnv("SAWERIA_SECRET_KEY", "your_default_secret_key_here")

	// Initialize Gin router
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Routes
	router.GET("/", healthCheck)
	router.POST("/webhook/saweria", handlers.SaweriaWebhookHandler(secretKey))

	// Start server
	log.Printf("Server starting on port %s", port)
	log.Printf("Webhook URL: http://localhost:%s/webhook/saweria", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":   "Saweria Webhook API is running",
		"timestamp": getCurrentTime(),
	})
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getCurrentTime() string {
	// Return current time in RFC3339 format
	return time.Now().Format(time.RFC3339)
}
