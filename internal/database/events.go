package database

import (
	"context"
	"database/sql"
	"time"
)

type EventModel struct {
	DB *sql.DB
}

type Events struct {
	Id          int       `json:"id"`
	OwnerId     int       `json:"ownerId" binding:"required"`
	Name        string    `json:"name" binding:"required,min=3"`
	Description string    `json:"description" binding:"required, min=10"`
	Date        time.Time `json:"date" binding:"required, datetime=2006-01-02"`
	Location    string    `json:"location" binding:"required, min=3"`
}

func (m *EventModel) Insert(event *Events) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (owner_id, name, description, date, location) VALUES ($1, $2, $3, $4, $5)"

	return m.DB.QueryRowContext(
		ctx,
		query,
		event.OwnerId,
		event.Name,
		event.Description,
		event.Date,
		event.Date,
		event.Location).Scan(&event.Id)
}
