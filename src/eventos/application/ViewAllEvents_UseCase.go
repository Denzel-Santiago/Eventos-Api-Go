package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type ViewAllEvents struct {
	db domain.IEvent
}

func NewViewAllEvents(db domain.IEvent) *ViewAllEvents {
	return &ViewAllEvents{db: db}
}

func (ve *ViewAllEvents) Execute() ([]entities.Event, error) {
	return ve.db.GetAll()
}
