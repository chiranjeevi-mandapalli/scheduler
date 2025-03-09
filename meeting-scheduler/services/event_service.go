package services

import (
	"meeting-scheduler/models"
	"meeting-scheduler/repositories"
)

type EventService struct {
	eventRepo repositories.EventRepository
}

func NewEventService(eventRepo repositories.EventRepository) *EventService {
	return &EventService{eventRepo: eventRepo}
}

func (s *EventService) CreateEvent(event *models.Event) error {
	return s.eventRepo.Save(event)
}


func (s *EventService) GetEvent(eventID string) (*models.Event, error) {
	return s.eventRepo.FindByID(eventID)
}

func (s *EventService) AddEventSlots(eventID string, slots []models.TimeSlot) error {
	return s.eventRepo.AddSlots(eventID, slots)
}

func (s *EventService) GetRecommendations(eventID string) ([]models.TimeSlot, error) {
	return s.eventRepo.FindAvailableSlots(eventID)
}
