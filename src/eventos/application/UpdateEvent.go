//Eventos-Api-Go/src/eventos/application/UpdateEvent.go
package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type UpdateEvent struct {
	Db domain.IEvent
}

func NewUpdateEvent(db domain.IEvent) *UpdateEvent {
	return &UpdateEvent{Db: db}
}

func (ue *UpdateEvent) Execute(id int, event entities.Event) error {
	return ue.Db.Update(id, event)
}
