package services

import (
	"meeting-scheduler/models"
	"meeting-scheduler/repositories"
)

type AvailabilityService struct {
	availabilityRepo repositories.AvailabilityRepository
}

func NewAvailabilityService(availabilityRepo repositories.AvailabilityRepository) *AvailabilityService {
	return &AvailabilityService{availabilityRepo: availabilityRepo}
}

func (s *AvailabilityService) SubmitAvailability(userId,userName,userEmail, eventID string, availability []models.TimeSlot) error {
	return s.availabilityRepo.Save(userId,userName,userEmail, eventID, availability)
}
