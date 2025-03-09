package models

type Event struct {
	ID       string     `json:"id" db:"id"`
	Title    string     `json:"title" db:"title"`
	Duration int        `json:"duration" db:"duration"` // In minutes
	Slots    []TimeSlot `json:"slots,omitempty" db:"-"`
}
