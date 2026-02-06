// Eventos-Api-Go/src/eventos/application/CreateEvents_UseCase.go
package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type CreateEventUseCase struct {
	eventRepo domain.IEvent
}

func NewCreateEventUseCase(eventRepo domain.IEvent) *CreateEventUseCase {
	return &CreateEventUseCase{
		eventRepo: eventRepo,
	}
}

func (uc *CreateEventUseCase) Run(event entities.Event) (entities.Event, error) {
	
	return uc.eventRepo.Save(event)
}