// Eventos-Api-Go/src/eventos/domain/entities/evento.go
package entities

import (
	"time"
)

type Event struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Location         string    `json:"location"`
	Date             time.Time `json:"date"` // Usamos time.Time para manejar fechas
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
		CreatedAt:        time.Now(), // Asignamos la fecha actual como valor por defecto
	}
}

// Métodos Getters y Setters
func (e *Event) GetID() int {
	return e.ID
}

func (e *Event) SetID(id int) {
	e.ID = id
}

func (e *Event) GetName() string {
	return e.Name
}

func (e *Event) SetName(name string) {
	e.Name = name
}

func (e *Event) GetLocation() string {
	return e.Location
}

func (e *Event) SetLocation(location string) {
	e.Location = location
}

func (e *Event) GetDate() time.Time {
	return e.Date
}

func (e *Event) SetDate(date time.Time) {
	e.Date = date
}

func (e *Event) GetAvailableTickets() int {
	return e.AvailableTickets
}

func (e *Event) SetAvailableTickets(availableTickets int) {
	e.AvailableTickets = availableTickets
}

func (e *Event) GetPrice() float64 {
	return e.Price
}

func (e *Event) SetPrice(price float64) {
	e.Price = price
}

func (e *Event) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *Event) SetCreatedAt(createdAt time.Time) {
	e.CreatedAt = createdAt
}
