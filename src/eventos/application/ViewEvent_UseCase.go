package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type ViewEvent struct {
	db domain.IEvent
}

func NewViewEvent(db domain.IEvent) *ViewEvent {
	return &ViewEvent{db: db}
}

func (vu *ViewEvent) Execute(id int) (entities.Event, error) {
	return vu.db.FindByID(id)
}
