package handlers

import (
	"meeting-scheduler/models"
	"meeting-scheduler/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventService *services.EventService
}

func NewEventHandler(eventService *services.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.eventService.CreateEvent(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func (h *EventHandler) GetEvent(c *gin.Context) {
	eventID := c.Param("eventId")

	event, err := h.eventService.GetEvent(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *EventHandler) AddEventSlots(c *gin.Context) {
	eventID := c.Param("eventId")
	var slots []models.TimeSlot

	if err := c.ShouldBindJSON(&slots); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.eventService.AddEventSlots(eventID, slots)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add slots"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Slots added successfully"})
}

func (h *EventHandler) GetRecommendations(c *gin.Context) {
	eventID := c.Param("eventId")

	slots, err := h.eventService.GetRecommendations(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recommendations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recommended_slots": slots})
}
