package repositories

import (
	"database/sql"
	"meeting-scheduler/models"
)

type EventRepository interface {
	Save(event *models.Event) error
	FindByID(eventID string) (*models.Event, error)
	AddSlots(eventID string, slots []models.TimeSlot) error
	FindAvailableSlots(eventID string) ([]models.TimeSlot, error)
}

type eventRepositoryImpl struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &eventRepositoryImpl{db: db}
}

func (r *eventRepositoryImpl) Save(event *models.Event) error {
	query := "INSERT INTO events (event_id, title, duration) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, event.ID, event.Title, event.Duration)
	return err
}

func (r *eventRepositoryImpl) FindByID(eventID string) (*models.Event, error) {
	var event models.Event
	query := "SELECT event_id, title, duration FROM events WHERE event_id = $1"
	err := r.db.QueryRow(query, eventID).Scan(&event.ID, &event.Title, &event.Duration)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepositoryImpl) AddSlots(eventID string, slots []models.TimeSlot) error {
	query := "INSERT INTO event_slots (event_id, start_time, end_time, timezone) VALUES ($1, $2, $3, $4)"
	for _, slot := range slots {
		_, err := r.db.Exec(query, eventID, slot.StartTime, slot.EndTime, slot.Timezone)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *eventRepositoryImpl) FindAvailableSlots(eventID string) ([]models.TimeSlot, error) {
	var slots []models.TimeSlot
	query := "SELECT start_time, end_time, timezone FROM event_slots WHERE event_id = $1"
	rows, err := r.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var slot models.TimeSlot
		if err := rows.Scan(&slot.StartTime, &slot.EndTime, &slot.Timezone); err != nil {
			return nil, err
		}
		slots = append(slots, slot)
	}

	return slots, nil
}
