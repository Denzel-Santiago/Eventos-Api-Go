// Eventos-Api-Go/src/eventos/domain/entities/evento.go
package entities

import (
	"time"
)

type Event struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Location         string    `json:"location"`
	Date             time.Time `json:"date"`
	AvailableTickets int       `json:"available_tickets"`
	Price            float64   `json:"price"`
	CreatedAt        time.Time `json:"created_at"`
}

// Constructor sin ID, ya que lo genera la BD
func NewEvent(name, location string, date time.Time, availableTickets int, price float64) *Event {
	return &Event{
		Name:             name,
		Location:         location,
		Date:             date,
		AvailableTickets: availableTickets,
		Price:            price,
		CreatedAt:        time.Now(),
	}
}