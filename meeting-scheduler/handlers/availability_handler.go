package handlers

import (
	"fmt"
	"log"
	"meeting-scheduler/models"
	"meeting-scheduler/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AvailabilityHandler struct {
	availabilityService *services.AvailabilityService
}

func NewAvailabilityHandler(availabilityService *services.AvailabilityService) *AvailabilityHandler {
	return &AvailabilityHandler{availabilityService: availabilityService}
}

func (h *AvailabilityHandler) SubmitAvailability(c *gin.Context) {
	eventId := c.Param("eventId")
	userId := c.Query("userId")
	userName := c.Query("userName")
	userEmailId := c.Query("userEmailId")

	// Validate required fields
	missingFields := []string{}
	if userId == "" {
		missingFields = append(missingFields, "userId")
	}
	if userName == "" {
		missingFields = append(missingFields, "userName")
	}
	if userEmailId == "" {
		missingFields = append(missingFields, "userEmailId")
	}
	if eventId == "" {
		missingFields = append(missingFields, "eventId")
	}

	if len(missingFields) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Missing required fields: %v", missingFields),
		})
		return
	}

	// Parse request body
	var availability []models.TimeSlot
	if err := c.ShouldBindJSON(&availability); err != nil {
		log.Printf("JSON binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON format",
			"details": err.Error(),
		})
		return
	}

	// Call service to save availability
	err := h.availabilityService.SubmitAvailability(userId, userName, userEmailId, eventId, availability)
	if err != nil {
		log.Printf("Failed to submit availability: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit availability"})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{"message": "Availability submitted successfully"})
}
