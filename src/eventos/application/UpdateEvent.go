// Eventos-Api-Go/src/eventos/application/UpdateEvent.go
package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type UpdateEvent struct {
	eventRepo domain.IEvent
}

func NewUpdateEvent(eventRepo domain.IEvent) *UpdateEvent {
	return &UpdateEvent{eventRepo: eventRepo}
}

func (ue *UpdateEvent) Execute(id int, event entities.Event) error {
	return ue.eventRepo.Update(id, event)
}