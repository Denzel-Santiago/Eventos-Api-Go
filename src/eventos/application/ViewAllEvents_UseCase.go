// Eventos-Api-Go/src/eventos/application/ViewAllEvents_UseCase.go
package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type ViewAllEvents struct {
	eventRepo domain.IEvent
}

func NewViewAllEvents(eventRepo domain.IEvent) *ViewAllEvents {
	return &ViewAllEvents{eventRepo: eventRepo}
}

func (ve *ViewAllEvents) Execute() ([]entities.Event, error) {
	return ve.eventRepo.GetAll()
}