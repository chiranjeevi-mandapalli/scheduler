package handlers

import (
	"database/sql"
	"meeting-scheduler/repositories"
	"meeting-scheduler/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {

	eventRepo := repositories.NewEventRepository(db)
	availabilityRepo := repositories.NewAvailabilityRepository(db)

	eventService := services.NewEventService(eventRepo)
	availabilityService := services.NewAvailabilityService(availabilityRepo)

	eventHandler := NewEventHandler(eventService)
	availabilityHandler := NewAvailabilityHandler(availabilityService)

	router.GET("/health", healthHandler)

	api := router.Group("/events")
	{
		api.POST("", eventHandler.CreateEvent)
		api.GET("/:eventId", eventHandler.GetEvent)
		api.POST("/:eventId/slots", eventHandler.AddEventSlots)
		api.GET("/:eventId/recommendations", eventHandler.GetRecommendations)

		api.POST("/:eventId/availability", availabilityHandler.SubmitAvailability)
	}
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Running"})
}
