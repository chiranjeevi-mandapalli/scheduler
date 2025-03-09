package test

import (
	"bytes"
	"encoding/json"
	"meeting-scheduler/handlers"
	"meeting-scheduler/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	eventService := &services.EventService{}
	availabilityService := &services.AvailabilityService{}

	eventHandler := handlers.NewEventHandler(eventService)
	availabilityHandler := handlers.NewAvailabilityHandler(availabilityService)

	r.POST("/events", eventHandler.CreateEvent)
	r.GET("/events/:eventId", eventHandler.GetEvent)
	r.POST("/events/:eventId/slots", eventHandler.AddEventSlots)
	r.POST("/events/:eventId/availability", availabilityHandler.SubmitAvailability)
	r.GET("/events/:eventId/recommendations", eventHandler.GetRecommendations)

	return r
}

func TestCreateEvent(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	requestBody, _ := json.Marshal(map[string]interface{}{
		"title":    "Brainstorming Meeting",
		"duration": 60,
	})
	req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestAddEventSlots(t *testing.T) {
	r := setupRouter()
	eventID := "test-event-id"
	w := httptest.NewRecorder()
	requestBody, _ := json.Marshal(map[string]interface{}{
		"slots": []map[string]interface{}{
			{"start_time": "2025-01-12T14:00:00Z", "end_time": "2025-01-12T16:00:00Z", "timezone": "EST"},
		},
	})
	req, _ := http.NewRequest("POST", "/events/"+eventID+"/slots", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestSubmitAvailability(t *testing.T) {
	r := setupRouter()
	eventID := "test-event-id"
	w := httptest.NewRecorder()
	requestBody, _ := json.Marshal(map[string]interface{}{
		"user_id": "user-123",
		"availability": []map[string]interface{}{
			{"start_time": "2025-01-12T14:00:00Z", "end_time": "2025-01-12T16:00:00Z", "timezone": "EST"},
		},
	})
	req, _ := http.NewRequest("POST", "/events/"+eventID+"/availability", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetRecommendations(t *testing.T) {
	r := setupRouter()
	eventID := "test-event-id"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/events/"+eventID+"/recommendations", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
