package models

type Availability struct {
	ID             string     `json:"id" db:"id"`
	UserID         string     `json:"userId" db:"user_id"`
	EventID        string     `json:"eventId" db:"event_id"`
	AvailableSlots []TimeSlot `json:"availableSlots,omitempty" db:"-"`
}
