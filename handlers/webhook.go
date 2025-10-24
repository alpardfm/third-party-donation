package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"saweria-webhook/models"

	"github.com/gin-gonic/gin"
)

// SaweriaWebhookHandler menangani webhook dari Saweria
func SaweriaWebhookHandler(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.SaweriaDonationPayload

		// Bind JSON payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			log.Printf("Error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, models.WebhookResponse{
				Status:  "error",
				Message: "Invalid JSON payload",
			})
			return
		}

		jsonPrint, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			log.Printf("Error marshalling payload: %v", err)
		} else {
			log.Printf("Received Saweria webhook payload:\n%s", string(jsonPrint))
		}
	}
}
