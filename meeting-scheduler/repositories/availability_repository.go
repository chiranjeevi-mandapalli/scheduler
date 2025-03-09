package repositories

import (
	"database/sql"
	"meeting-scheduler/models"
)

type AvailabilityRepository interface {
	Save(userId,userName,userEmail, eventID string, availability []models.TimeSlot) error
}

type availabilityRepositoryImpl struct {
	db *sql.DB
}

func NewAvailabilityRepository(db *sql.DB) AvailabilityRepository {
	return &availabilityRepositoryImpl{db: db}
}

func (r *availabilityRepositoryImpl) Save(userID, userName, userEmail, eventID string, availability []models.TimeSlot) error {
	tx, err := r.db.Begin() 
	if err != nil {
		return err
	}

	userQuery := "INSERT INTO users (user_id, name, email) VALUES ($1, $2, $3) ON CONFLICT (user_id) DO NOTHING"
	_, err = tx.Exec(userQuery, userID, userName, userEmail)
	if err != nil {
		tx.Rollback()
		return err
	}
	query := "INSERT INTO availability (id, user_id, event_id, start_time, end_time, timezone) VALUES ($1, $2, $3, $4, $5, $6)"
	for _, slot := range availability {
		_, err := tx.Exec(query, slot.ID, userID, eventID, slot.StartTime, slot.EndTime, slot.Timezone)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
