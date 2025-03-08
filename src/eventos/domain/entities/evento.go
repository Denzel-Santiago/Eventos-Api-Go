package entities

import "time"

type Evento struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Location         string    `json:"location"`
	Date             time.Time `json:"date"`
	AvailableTickets int       `json:"available_tickets"`
	Price            float64   `json:"price"`
	CreatedAt        time.Time `json:"created_at"`
}
