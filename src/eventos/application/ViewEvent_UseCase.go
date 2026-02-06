// Eventos-Api-Go/src/eventos/application/ViewEvent_UseCase.go
package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type ViewEvent struct {
	eventRepo domain.IEvent
}

func NewViewEvent(eventRepo domain.IEvent) *ViewEvent {
	return &ViewEvent{eventRepo: eventRepo}
}

func (vu *ViewEvent) Execute(id int) (entities.Event, error) {
	return vu.eventRepo.FindByID(id)
}