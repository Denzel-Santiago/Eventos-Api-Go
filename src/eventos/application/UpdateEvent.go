package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type UpdateEvent struct {
	db domain.IEvent
}

func NewUpdateEvent(db domain.IEvent) *UpdateEvent {
	return &UpdateEvent{db: db}
}

func (ue *UpdateEvent) Execute(id int, event entities.Event) error {
	return ue.db.Update(id, event)
}
