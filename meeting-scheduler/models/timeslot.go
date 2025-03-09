package models

import "time"

type TimeSlot struct {
	ID        string    `json:"id" db:"id"`
	EventID   string    `json:"eventId" db:"event_id"`
	StartTime time.Time `json:"startTime" db:"start_time"`
	EndTime   time.Time `json:"endTime" db:"end_time"`
	Timezone  string    `json:"timezone" db:"timezone"`
}
